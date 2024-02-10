package report

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/fschuermeyer/GoWordlytics/internal/format"
)

type Report struct {
	url            string
	hasWordPress   bool
	hasReadme      bool
	hasMuPlugins   bool
	version        string
	versionStatus  string
	versionCurrent string
	themes         []Theme
	pluginDetails  map[string]PluginDetails
	status         string
}

func (r *Report) SetUrl(url string) {
	url, ok := format.URL(url)

	if ok {
		r.url = url
	}
}

func (r *Report) SetVersion(version string) {
	r.version = strings.TrimSpace(version)
}

func (r *Report) SetVersionUpdate(status, current string) {
	r.versionStatus = status
	r.versionCurrent = current
}

func (r *Report) GetUrl() string {
	return r.url
}

func (r *Report) SetHasWordPress(hasWordPress bool) {
	r.hasWordPress = hasWordPress
}

func (r *Report) SetHasReadme(hasReadme bool) {
	r.hasReadme = hasReadme
}

func (r *Report) HasWordPress() bool {
	return r.hasWordPress
}

func (r *Report) Output() {
	color.White("\nGoWordlytics Report")
	color.Blue("Report for %s", r.GetUrl())
	fmt.Println("------------------------")

	c := color.New(color.FgGreen)

	c.Print("Has WordPress: ")
	fmt.Println(r.hasWordPress)

	c.Print("Has readme.html: ")
	fmt.Println(r.hasReadme)

	if len(r.version) > 0 {
		c.Print("Version: ")
		fmt.Printf("%s ", r.version)

		if len(r.versionStatus) > 0 && r.versionStatus != "error" && r.versionStatus != "latest" {
			text := color.RedString("(%s to %s)", r.versionStatus, r.versionCurrent)

			fmt.Print(text)
		}

		fmt.Println()
	}

	fmt.Print("------------------------\n\n")
}
