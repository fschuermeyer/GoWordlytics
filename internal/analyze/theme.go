package analyze

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fschuermeyer/GoWordlytics/internal/format"
	"github.com/fschuermeyer/GoWordlytics/internal/report"
	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

func (a *Analyze) getThemes() []report.Theme {
	var themes []report.Theme

	for _, theme := range a.getThemeNames() {
		themeData := a.getThemeData(theme)

		if len(themeData.Version) > 0 || len(themeData.Name) > 0 {
			themes = append(themes, themeData)
		}
	}

	return themes
}

func (a *Analyze) getThemeData(name string) report.Theme {
	url := fmt.Sprintf("%swp-content/themes/%s/style.css", a.data.url, name)

	limit, err := request.CalculateMiB(5)

	if err != nil {
		return report.Theme{}
	}

	response, err := request.Do(url, a.userAgent, limit)

	if err != nil {
		return report.Theme{}
	}

	themeMap := a.parseCSSThemeString(response)

	return a.getThemeDetials(themeMap, name)
}

func (a *Analyze) getThemeDetials(themeMap map[string]string, name string) report.Theme {
	themeDetails := report.Theme{}

	fields := map[string]func(string){
		"name":         func(value string) { themeDetails.Name = value },
		"uri":          func(value string) { themeDetails.URI = value },
		"version":      func(value string) { themeDetails.Version = value },
		"author":       func(value string) { themeDetails.Author = value },
		"author uri":   func(value string) { themeDetails.AuthorURI = value },
		"licence":      func(value string) { themeDetails.Licence = value },
		"licence uri":  func(value string) { themeDetails.LicenceURI = value },
		"description":  func(value string) { themeDetails.Description = value },
		"min version":  func(value string) { themeDetails.MinVersion = value },
		"max version":  func(value string) { themeDetails.MaxVersion = value },
		"required php": func(value string) { themeDetails.RequiredPHP = value },
		"template":     func(value string) { themeDetails.Template = value },
		"status":       func(value string) { themeDetails.Status = value },
		"tags":         func(value string) { themeDetails.Tags = strings.Split(value, ",") },
		"text domain":  func(value string) { themeDetails.TextDomain = value },
		"domain path":  func(value string) { themeDetails.DomainPath = value },
	}

	for key, setValue := range fields {
		if value, ok := themeMap[key]; ok {
			setValue(value)
		}
	}

	if len(themeDetails.Name) == 0 {
		themeDetails.Name = name
	}

	return themeDetails
}

func (a *Analyze) getThemeNames() []string {
	var themes []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.data.htmlIndex))

	if err != nil {
		return themes
	}

	doc.Find("link[rel=stylesheet]").Each(func(index int, s *goquery.Selection) {
		href := s.AttrOr("href", "")

		if !(strings.Contains(href, "/wp-content/themes/")) {
			return
		}

		parts := strings.Split(href, "/themes/")

		if len(parts) != 2 {
			return
		}

		themes = append(themes, strings.Split(parts[1], "/")[0])

	})

	return format.UniqueSlice(themes)
}

func (a *Analyze) parseCSSThemeString(cssString string) map[string]string {
	themeInfo := make(map[string]string)

	lines := strings.Split(cssString, "\n")
	collectStrings := false

	for _, line := range lines {
		breakAfter := false
		line := strings.TrimSpace(line)

		if strings.HasPrefix(strings.TrimSpace(line), "/*") {
			collectStrings = true
		}

		if strings.Contains(line, "*/") {
			line = strings.Split(line, "*/")[0]
			breakAfter = true
		}

		if !collectStrings {
			continue
		}

		comment := strings.TrimSpace(strings.TrimSuffix(strings.TrimPrefix(line, "/*"), "*/"))

		pairs := strings.Split(comment, ";")

		for _, pair := range pairs {
			parts := strings.SplitN(strings.TrimSpace(pair), ":", 2)

			if len(parts) != 2 {
				continue
			}

			key := strings.ToLower(strings.TrimSpace(strings.Trim(strings.TrimSpace(parts[0]), "*")))
			value := strings.TrimSpace(parts[1])
			themeInfo[key] = value
		}

		if collectStrings && len(line) == 0 || breakAfter {
			break
		}
	}

	return themeInfo
}
