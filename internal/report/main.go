package report

import "github.com/fschuermeyer/GoWordlytics/internal/format"

type Report struct {
	url           string
	hasWordPress  bool
	hasReadme     bool
	hasMuPlugins  bool
	version       string
	versionStatus string
	themes        []Theme
	pluginDetails map[string]PluginDetails
	status        string
}

func (r *Report) SetUrl(url string) {
	url, ok := format.URL(url)

	if ok {
		r.url = url
	}
}

func (r *Report) SetVersion(version string) {
	r.version = version
}

func (r *Report) GetUrl() string {
	return r.url
}

func (r *Report) SetHasWordPress(hasWordPress bool) {
	r.hasWordPress = hasWordPress
}

func (r *Report) HasWordPress() bool {
	return r.hasWordPress
}
