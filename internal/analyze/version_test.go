package analyze_test

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/fschuermeyer/GoWordlytics/internal/analyze"
)

func TestVersionByMetaTag(t *testing.T) {

	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "no meta tag",
			html: `<html><head></head><body></body></html>`,
			want: "",
		},
		{
			name: "empty meta tag",
			html: `<html><head><meta name="generator" content=""></head><body></body></html>`,
			want: "",
		},
		{
			name: "meta tag with no content",
			html: `<html><head><meta name="generator"></head><body></body></html>`,
			want: "",
		},
		{
			name: "meta tag with version",
			html: `<html><head><meta name="generator" content="WordPress 5.5.1"></head><body></body></html>`,
			want: "5.5.1",
		},
		{
			name: "meta tag with no wordpress version",
			html: `<html><head><meta name="generator" content="Elements 5.5.1"></head><body></body></html>`,
			want: "",
		},
		{
			name: "multiple meta tags",
			html: `<html><head><meta name="generator" content="WordPress 5.5.1"><meta name="generator" content="WordPress 6.5.1"></head><body></body></html>`,
			want: "5.5.1",
		},
		{
			name: "no html",
			html: ``,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyze.New()

			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			got := a.VersionByMetaTag(doc)

			if got != tt.want {
				t.Errorf("VersionByMetaTag() got = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestVersionByEnquedScripts(t *testing.T) {
	tests := []struct {
		name string
		html string
		want string
	}{
		{
			name: "no scripts",
			html: `<html><head></head><body></body></html>`,
			want: "",
		},
		{
			name: "no matching scripts",
			html: `<html><head><script src="script1.js"></script><script src="script2.js"></script></head><body></body></html>`,
			want: "",
		},
		{
			name: "matching script with version",
			html: `<html><head><script src="script1.js"></script><script src="script2.js"></script><script src="/wp-includes/js/wp-embed.min.js?ver=1.2.3"></script></head><body></body></html>`,
			want: "1.2.3",
		},
		{
			name: "matching script with invalid version",
			html: `<html><head><script src="script1.js"></script><script src="script2.js"></script><script src="script3.js?v=invalid"></script></head><body></body></html>`,
			want: "",
		},
		{
			name: "multiple matching scripts",
			html: `<html><head><script src="script1.js?v=1.2.3"></script><script src="/wp-includes/js/wp-embed.min.js?ver=5.4.2"></script><script src="/wp-includes/css/dist/block-library/style.min.css?ver=7.8.9"></script></head><body></body></html>`,
			want: "5.4.2",
		},
		{
			name: "no html",
			html: ``,
			want: "",
		},
		{
			name: "no head",
			html: `<html><body></body></html>`,
			want: "",
		},
		{
			name: "css link & script",
			html: `<html><head><link rel="stylesheet" href="/wp-includes/css/dist/block-library/style.min.css?ver=7.8.9"><script src="/wp-includes/js/wp-embed.min.js?ver=5.4.2"></script></head><body></body></html>`,
			want: "7.8.9",
		},
		{
			name: "css link",
			html: `<html><head><link rel="stylesheet" href="/wp-includes/css/dist/block-library/style.min.css?ver=7.8.9"></head><body></body></html>`,
			want: "7.8.9",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyze.New()

			doc, _ := goquery.NewDocumentFromReader(strings.NewReader(tt.html))
			got := a.VersionByEnquedScripts(doc)

			if got != tt.want {
				t.Errorf("VersionByEnquedScripts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestVersionByLoginPage(t *testing.T) {
	tests := []struct {
		name string
		resp string
		want string
	}{
		{
			name: "no link",
			resp: `<html><head></head><body></body></html>`,
			want: "",
		},
		{
			name: "no matching link",
			resp: `<html><head><link rel="stylesheet" href="style.css"></head><body></body></html>`,
			want: "",
		},
		{
			name: "matching link with version",
			resp: `<html><head><link rel='stylesheet' href='https://example.com/wp-admin/load-styles.php?c=1&amp;dir=ltr&amp;load%5B%5D=dashicons,buttons,forms,l10n,login&amp;ver=4.9.8' type='text/css' media='all' /></head><body></body></html>`,
			want: "4.9.8",
		},
		{
			name: "matching link with version",
			resp: `<html><head><link rel="stylesheet" href="https://example.com/wp-admin/load-styles.php?c=1&amp;dir=ltr&amp;load%5B%5D=dashicons,buttons,forms,l10n,login&amp;ver=1.2.0"></head><body></body></html>`,
			want: "1.2.0",
		},
		{
			name: "matchin with multiple links",
			resp: `<html>
			<head>
				<link rel="stylesheet" href="https://example.com/wp-admin/load-styles.php?c=1&amp;dir=ltr&amp;load%5B%5D=dashicons,buttons,forms,l10n,login&amp;ver=1.2.0">
				<link rel='stylesheet' id='forms-css' href='https://www.example.com/wp-admin/css/forms.min.css?ver=6.2.1' type='text/css' media='all' />
				<link rel='stylesheet' id='l10n-css' href='https://www.example.com/wp-admin/css/l10n.min.css?ver=6.2.2' type='text/css' media='all' />
				<link rel='stylesheet' id='login-css' href='https://www.example.com/wp-admin/css/login.min.css?ver=6.2.3' type='text/css' media='all' />
			</html>`,
			want: "1.2.0",
		},
		{
			name: "no html",
			resp: ``,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := analyze.New()

			got := a.VersionByLoginPage(tt.resp)

			if got != tt.want {
				t.Errorf("VersionByLoginPage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
