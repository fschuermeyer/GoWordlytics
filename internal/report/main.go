package report

import (
	"fmt"
	"html"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/fschuermeyer/GoWordlytics/internal/format"
	"github.com/fschuermeyer/GoWordlytics/internal/render"
	"github.com/fschuermeyer/GoWordlytics/internal/wordpress"
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
	pluginDetails  []PluginDetails
	users          []wordpress.User
	status         string
}

func (r *Report) SetUrl(url string) bool {
	url, ok := format.URL(url)

	if ok {
		r.url = url
	}

	return ok
}

func (r *Report) SetVersion(version string) {
	r.version = strings.TrimSpace(version)
}

func (r *Report) SetPlugins(plugins []PluginDetails) {
	r.pluginDetails = plugins
}

func (r *Report) SetThemes(themes []Theme) {
	r.themes = themes
}

func (r *Report) SetUsers(users []wordpress.User) {
	r.users = users
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

// Render renders the report by calling the RenderOverview, RenderPlugins, and RenderThemes methods.
// It applies the provided headline style to the rendered sections.
func (r *Report) Render() {
	var headline = lipgloss.NewStyle().Foreground(lipgloss.Color("#E7D617")).Bold(true)

	r.renderOverview(headline)

	if len(r.pluginDetails) > 0 {
		r.renderPlugins(headline)
	}

	if len(r.themes) > 0 {
		r.renderThemes(headline)
	}

	if len(r.users) > 0 {
		r.renderUsers(headline)
	}
}

// RenderOverview renders the overview of the GoWordlytics report.
// It takes a headline lipgloss.Style as a parameter and prints the report to the console.
// The report includes the URL and whether it is a WordPress site.
// If the site is a WordPress site, it also includes whether it has a readme file.
// If the report includes version information, it also includes the version number, version status, and current version.
// The report is displayed in a table format.
func (r *Report) renderOverview(headline lipgloss.Style) {
	fmt.Printf("%s\n", headline.Render("GoWordlytics Report"))

	headers := []string{"URL", "WordPress?"}
	values := []string{r.GetUrl(), fmt.Sprintf("%t", r.HasWordPress())}

	if r.HasWordPress() {
		headers = append(headers, "Readme?")
		values = append(values, fmt.Sprintf("%t", r.hasReadme))
	}

	if len(r.version) > 0 {
		headers = append(headers, "Version")
		values = append(values, r.version)

		if len(r.versionStatus) > 0 && r.versionStatus != "error" && r.versionStatus != "latest" {
			headers = append(headers, "Version Status")

			if r.versionStatus == "upgrade" {
				r.versionStatus = lipgloss.NewStyle().Foreground(lipgloss.Color("#F16208")).Render(r.versionStatus)
			}

			values = append(values, r.versionStatus)

			headers = append(headers, "Current")
			values = append(values, r.versionCurrent)
		}
	}

	render.Table(headers, [][]string{values})
}

// RenderPlugins renders the plugins report with the given headline style.
// It prints the headline and then displays a table with the plugin details,
// including the name, slug, version, number of downloads, and homepage link.
func (r *Report) renderPlugins(headline lipgloss.Style) {
	fmt.Printf("%s\n", headline.Render("Plugins Report"))

	rows := [][]string{}

	for _, plugin := range r.pluginDetails {
		rows = append(rows, []string{plugin.Slug, html.UnescapeString(plugin.Name), plugin.Version, format.InsertThousandSeparator(plugin.Downloaded, '.'), plugin.Homepage})
	}

	render.Table([]string{"Slug", "Name", "Version", "Downloads", "Link"}, rows)
}

// RenderThemes renders a report of the themes.
// It takes a headline style as input and prints the themes' information in a table format.
// The themes' information includes the name, description, text domain, author, author URI, and version.
func (r *Report) renderThemes(headline lipgloss.Style) {

	rows := [][]string{}

	for _, theme := range r.themes {
		rows = append(rows, []string{theme.Name, theme.Description, theme.TextDomain, theme.Author, theme.AuthorURI, theme.Version})
	}

	fmt.Printf("%s\n", headline.Render("Themes Report"))
	render.Table([]string{"Name", "Description", "TextDomain", "Author", "AuthorURI", "Version"}, rows)
}

func (r *Report) renderUsers(headline lipgloss.Style) {
	rows := [][]string{}

	for _, user := range r.users {
		rows = append(rows, []string{fmt.Sprintf("%d", user.ID), user.Name, html.UnescapeString(user.Description), user.URL, user.Link, user.Slug})
	}

	fmt.Printf("%s\n", headline.Render("Users Report"))

	render.Table([]string{"ID", "Name", "Description", "URL", "Link", "Slug"}, rows)
}
