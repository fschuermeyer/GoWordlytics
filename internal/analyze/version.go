package analyze

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (a *Analyze) version(url string) string {
	var version string

	if len(a.data.htmlIndex) > 0 {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.data.htmlIndex))

		if err == nil {
			version = a.versionByMetaTag(doc)
		}

		if err == nil && len(version) == 0 {
			version = a.versionByEnquedScripts(doc)
		}
	}

	if len(version) == 0 {
		return "0.0.0"
	}

	return version
}

func (a *Analyze) versionByMetaTag(doc *goquery.Document) string {
	s := doc.Find("meta[name=generator]").First()

	value := strings.TrimSpace(s.AttrOr("content", ""))

	for _, indicator := range a.vIndicatorsMetaTag {
		if strings.HasPrefix(value, indicator.indicator) {
			return strings.ReplaceAll(value, indicator.indicator, "")
		}
	}

	return ""
}

func (a *Analyze) versionByEnquedScripts(doc *goquery.Document) string {
	for _, indicator := range a.vIndicatorsEnquedScripts {
		var sources []*goquery.Selection
		var attr string

		if strings.Contains(indicator.indicator, ".css") {
			doc.Find("link").Each(func(i int, s *goquery.Selection) {
				sources = append(sources, s)
			})
			attr = "href"
		} else {
			doc.Find("script").Each(func(i int, s *goquery.Selection) {
				sources = append(sources, s)
			})
			attr = "src"
		}

		for _, s := range sources {
			attrContent := strings.TrimSpace(s.AttrOr(attr, ""))

			if len(attrContent) == 0 || !strings.Contains(attrContent, indicator.indicator) {
				continue
			}

			value := strings.Split(attrContent, indicator.indicator)[1]

			if len(value) > 1 && len(value) < 8 {
				return value
			}
		}
	}

	return ""
}

func (a *Analyze) versionByRssFeed() string {
	return "1.0"
}

func (a *Analyze) versionByLoginPage() string {
	return "1.0"
}
