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
		data := a.getPluginInfo(a.getPlugin(plugin))

		pluginList = append(pluginList, data)
	}

	return pluginList
}

func (a *Analyze) getPlugin(slug string) report.PluginDetails {
	var plugin report.PluginDetails

	plugin.Slug = slug

	size, err := request.CalculateMiB(6)

	if err != nil {
		return plugin
	}

	jsonData, err := request.Do(fmt.Sprintf(a.apiPlugins, slug), a.userAgent, size)

	if jsonData == "" || err != nil {
		return plugin
	}

	err = json.Unmarshal([]byte(jsonData), &plugin)

	if err != nil {
		return plugin
	}

	return plugin
}

func (a *Analyze) getPluginInfo(plugin report.PluginDetails) report.PluginDetails {
	for _, hint := range a.hintPlugins {
		if hint.slug == plugin.Slug && plugin.Name == "" {
			plugin.Name = hint.name + " (+)"
		}

		if hint.slug == plugin.Slug && plugin.Homepage == "" {
			plugin.Homepage = hint.uri
		}
	}

	return plugin
}

func (a *Analyze) getPluginsSlugs() []string {
	var plugins []string

	plugins = append(plugins, a.getPluginsFromLinks()...)

	plugins = append(plugins, a.getPluginsFromHints()...)

	return plugins
}

func (a *Analyze) getPluginsFromHints() []string {
	var plugins []string

	for _, hint := range a.hintPlugins {
		if strings.Contains(a.data.htmlIndex, hint.key) {
			plugins = append(plugins, hint.slug)
		}
	}

	return plugins
}

func (a *Analyze) getPluginsFromLinks() []string {
	var plugins []string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(a.data.htmlIndex))

	if err != nil {
		return []string{}
	}

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
					plugins = append(plugins, strings.ToLower(key[0]))
				}
			}
		}
	})

	return plugins
}
