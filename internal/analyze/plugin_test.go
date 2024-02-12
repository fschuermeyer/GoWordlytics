package analyze_test

import (
	"reflect"
	"testing"

	"github.com/fschuermeyer/GoWordlytics/internal/analyze"
)

func TestAnalyzeGetPluginsFromHints(t *testing.T) {

	tests := []struct {
		name     string
		html     string
		expected []string
	}{
		{
			name:     "Empty HTML",
			html:     "",
			expected: []string{},
		},
		{
			name:     "HTML with hint1",
			html:     "<html><head><title>This is a Site</title></head><body><-- Performance optimized by W3 Total Cache --></body></html>",
			expected: []string{"w3-total-cache"},
		},
		{
			name:     "HTML with hint2",
			html:     "<html><head><title>This is a Site</title></head><body><-- Performance optimized by W3 Total Cache --><script src=\"domain.tld/cache/autoptimize/\"</body></html>",
			expected: []string{"w3-total-cache", "autoptimize"},
		},
	}

	a := analyze.New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a.SetHTMLStore(test.html)

			result := a.GetPluginsFromHints()

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d items, but got %d", len(test.expected), len(result))
				return
			}

			if !reflect.DeepEqual(result, test.expected) && !(len(result) == len(test.expected) && len(result) == 0) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestAnalyzeGetPluginsFromLinks(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected []string
	}{
		{
			name:     "Empty HTML",
			html:     "",
			expected: []string{},
		},
		{
			name:     "HTML with no links",
			html:     "<html><head><title>This is a Site</title></head><body></body></html>",
			expected: []string{},
		},
		{
			name:     "HTML with link to plugin",
			html:     "<html><head><title>This is a Site</title></head><body><script src=\"domain.tld/wp-content/plugins/yoast-seo/\"</body></html>",
			expected: []string{"yoast-seo"},
		},
		{
			name:     "HTML with link to plugin and other link",
			html:     "<html><head><title>This is a Site</title></head><body><script src=\"domain.tld/wp-content/plugins/yoast-seo/\"<script src=\"domain.tld/cache/autoptimize/\"</body></html>",
			expected: []string{"yoast-seo"},
		},
	}

	a := analyze.New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a.SetHTMLStore(test.html)

			result := a.GetPluginsFromLinks()

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d items, but got %d", len(test.expected), len(result))
				return
			}

			if !reflect.DeepEqual(result, test.expected) && !(len(result) == len(test.expected) && len(result) == 0) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
