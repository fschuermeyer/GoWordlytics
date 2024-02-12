package analyze

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/fschuermeyer/GoWordlytics/internal/format"
	"github.com/fschuermeyer/GoWordlytics/internal/report"
	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

func (a *Analyze) getPlugins() []report.PluginDetails {
	plugins := format.UniqueSlice(a.getPluginsSlugs())

	var pluginList []report.PluginDetails

	for _, plugin := range plugins {
		data, ok := a.getPlugin(plugin)

		if ok {
			pluginList = append(pluginList, data)
		}
	}

	return pluginList
}

func (a *Analyze) getPlugin(slug string) (report.PluginDetails, bool) {
	var plugin report.PluginDetails

	size, err := request.CalculateMiB(6)

	if err != nil {
		return plugin, false
	}

	jsonData, err := request.Do(fmt.Sprintf(a.apiPlugins, slug), a.userAgent, size)

	if jsonData == "" || err != nil {
		return plugin, false
	}

	err = json.Unmarshal([]byte(jsonData), &plugin)

	if err != nil {
		return plugin, false
	}

	return plugin, true
}

func (a *Analyze) getPluginsSlugs() []string {
	var plugins []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.data.htmlIndex))

	if err != nil {
		return []string{}
	}

	plugins = append(plugins, a.getPluginsFromLinks(doc)...)

	plugins = append(plugins, a.getPluginsFromHints(a.data.htmlIndex)...)

	return plugins
}

func (a *Analyze) getPluginsFromHints(html string) []string {
	var plugins []string

	for _, hint := range a.hintPlugins {
		if strings.Contains(html, hint.key) {
			plugins = append(plugins, hint.slug)
		}
	}

	return plugins
}

func (a *Analyze) getPluginsFromLinks(doc *goquery.Document) []string {
	var plugins []string

	doc.Find("script, link").Each(func(i int, s *goquery.Selection) {
		src, ok := s.Attr("src")

		if !ok {
			src, _ = s.Attr("href")
		}

		if src == "" {
			return
		}

		src = strings.TrimSpace(strings.ToLower(src))

		for _, indicator := range a.pluginIndicators {
			if strings.Contains(src, indicator) {
				key := strings.Split(src, indicator)

				if len(key) > 1 {
					key = strings.Split(key[1], "/")
				}

				if len(key) > 0 {
					plugins = append(plugins, key[0])
				}
			}
		}
	})

	return plugins
}
