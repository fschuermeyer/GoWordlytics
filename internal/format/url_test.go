package format_test

import (
	"testing"

	"github.com/fschuermeyer/GoWordlytics/internal/format"
)

func TestURL(t *testing.T) {

	tests := []struct {
		name   string
		url    string
		format string
		valid  bool
	}{
		{"valid url", "https://www.example.com", "https://www.example.com", true},
		{"wrong protocol seperator", "http:///www.example.com", "http://www.example.com", true},
		{"wrong multiple protocol seperator", "http://////www.example.com", "http://www.example.com", true},
		{"valid case upper url", "HTTPS://WWW.EXAMPle.CoM", "https://www.example.com", true},
		{"invalid tld", "https://www.example", "", false},
		{"not accpted protocol", "ftp://www.example.com", "", false},
		{"invalid protocol", "example://www.example.com", "", false},
		{"valid url with whitespaces", " https://www.example.com ", "https://www.example.com", true},
		{"missing protocol", " www.example.com ", "http://www.example.com", true},
		{"remove query params", " www.example.com?example=1 ", "http://www.example.com", true},
		{"remove fragments params", " www.example.com#example ", "http://www.example.com", true},
		{"empty string", "", "", false},
		{"only protocol", "https://", "", false},
		{"multiple subdomains", "https://www.archive.blog.example.com", "https://www.archive.blog.example.com", true},
		{"url without protocol and userdata", "user:password@example.com", "http://user:password@example.com", true},
		{"url with userdata", "https://user:password@example.com", "https://user:password@example.com", true},
		{"url with userdata with uppercase letters", "https://uSEr:pasSWord@example.com", "https://uSEr:pasSWord@example.com", true},
		{"url with user only", " USER@example.com", "http://USER@example.com", true},
		{"url with path", "http://example.com/example", "http://example.com/example", true},
		{"komplex case", "  user:pasSWord@www.archive.blog.example.com/test?example=1#example  ", "http://user:pasSWord@www.archive.blog.example.com/test", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			format, valid := format.URL(tt.url)

			if valid != tt.valid {
				t.Errorf("URL Valid(%v) = %v, want %v", tt.url, valid, tt.valid)
			}

			if format != tt.format {
				t.Errorf("URL Format(%v) = %v, want %v", tt.url, format, tt.format)
			}

		})
	}
}
