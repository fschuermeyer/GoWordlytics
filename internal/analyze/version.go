package analyze

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (a *Analyze) version(url string) string {
	var version string

	if len(a.data.htmlIndex) != 0 {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.data.htmlIndex))

		if err != nil {
			version = a.versionByMetaTag(doc)
		}

		if err != nil && len(version) == 0 {
			version = a.versionByEnquedScripts(doc)
		}
	}

	return "1.0"
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
	return "1.0"
}

func (a *Analyze) versionByRssFeed() string {
	return "1.0"
}

func (a *Analyze) versionByLoginPage() string {
	return "1.0"
}
