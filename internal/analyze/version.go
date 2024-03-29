package analyze

import (
	"encoding/xml"
	"html"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (a *Analyze) version() string {
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
		resp := a.getContent("wp-login.php", 4)

		if len(resp) != 0 {
			version = a.versionByLoginPage(resp)
		}
	}

	if len(version) == 0 {
		resp := a.getContent("/feed", 4)

		if len(resp) != 0 {
			version = a.versionByRssFeed(resp)
		}
	}

	return version
}

func (a *Analyze) versionByMetaTag(doc *goquery.Document) string {
	s := doc.Find("meta[name=generator]").First()

	value := s.AttrOr("content", "")

	if strings.HasPrefix(value, a.vIndicatorsMetaTag.indicator) {
		version := strings.TrimSpace(strings.ReplaceAll(value, a.vIndicatorsMetaTag.indicator, ""))

		return strings.TrimSpace(version)
	}

	return ""
}

func (a *Analyze) versionByEnquedScripts(doc *goquery.Document) string {
	version := ""

	doc.Find("link,script").EachWithBreak(func(i int, s *goquery.Selection) bool {
		attrs, ok := s.Attr("href")

		if !ok {
			attrs, ok = s.Attr("src")
		}

		if !ok {
			return false
		}

		for _, indicator := range a.vIndicatorsEnquedScripts {
			if strings.Contains(attrs, indicator.indicator) {
				value := strings.Split(attrs, indicator.indicator)[1]

				if len(value) > 1 && len(value) < 8 {
					version = value
					return false
				}
			}
		}

		return true
	})

	return version
}

func (a *Analyze) versionByLoginPage(resp string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp))

	if err != nil {
		return ""
	}

	version := ""

	doc.Find("link").EachWithBreak(func(i int, s *goquery.Selection) bool {
		for _, indicator := range a.vIndicatorsLoginPage {
			attr := html.UnescapeString(strings.TrimSpace(s.AttrOr("href", "")))

			split := indicator.split

			if len(split) == 0 {
				split = indicator.indicator
			}

			if strings.Contains(attr, indicator.indicator) {
				value := strings.Split(attr, split)[1]

				if len(value) > 1 && len(value) < 8 {
					version = value
					return false
				}
			}
		}

		return true
	})

	return version
}

type RssFeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Generator string `xml:"generator"`
}

func (a *Analyze) versionByRssFeed(resp string) string {
	var rssFeed RssFeed

	err := xml.Unmarshal([]byte(resp), &rssFeed)

	if err != nil {
		return ""
	}

	if len(rssFeed.Channel.Generator) > 0 && strings.Contains(rssFeed.Channel.Generator, a.vIndicatorsRssFeed.indicator) {
		return strings.TrimSpace(strings.ReplaceAll(rssFeed.Channel.Generator, a.vIndicatorsRssFeed.indicator, ""))
	}

	return ""
}
