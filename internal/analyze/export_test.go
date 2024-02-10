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
