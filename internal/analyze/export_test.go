package analyze

import "github.com/PuerkitoBio/goquery"

func (a *Analyze) SetHTMLStore(html string) {
	a.data.htmlIndex = html
}

func (a *Analyze) VersionByMetaTag(doc *goquery.Document) string {
	return a.versionByMetaTag(doc)
}

func (a *Analyze) VersionByEnquedScripts(doc *goquery.Document) string {
	return a.versionByEnquedScripts(doc)
}

func (a *Analyze) VersionByLoginPage(resp string) string {
	return a.versionByLoginPage(resp)
}

func (a *Analyze) VersionByRssFeed(resp string) string {
	return a.versionByRssFeed(resp)
}

func (a *Analyze) GetPluginsFromLinks() []string {
	return a.getPluginsFromLinks()
}

func (a *Analyze) GetPluginsFromHints() []string {
	return a.getPluginsFromHints()
}

func (a *Analyze) ParseCSSThemeString(cssString string) map[string]string {
	return a.parseCSSThemeString(cssString)
}
