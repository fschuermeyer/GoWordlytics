package analyze_test

import (
	"reflect"
	"testing"

	"github.com/fschuermeyer/GoWordlytics/internal/analyze"
)

func TestParseCSSThemeString(t *testing.T) {
	tests := []struct {
		name     string
		css      string
		expected map[string]string
	}{
		{
			name: "parses theme css correctly",
			css: `/*
				Theme Name: My Theme
				Theme URI: http://example.com/my-theme
				Description: A description of my theme
				Link: http://example.com/my-theme
				Version: 1.0
				Author: John Doe
				Author URI: http://example.com/
			*/

				body {
					background-color: #fff;
					color: #333;
				}`,
			expected: map[string]string{
				"theme name":  "My Theme",
				"theme uri":   "http://example.com/my-theme",
				"description": "A description of my theme",
				"version":     "1.0",
				"author":      "John Doe",
				"author uri":  "http://example.com/",
				"link":        "http://example.com/my-theme",
			},
		},
		{
			name:     "returns empty map for empty css string",
			css:      "* { color: #fff; }",
			expected: map[string]string{},
		},
		{
			name: "returns empty map for empty comment",
			css: `/*
			*/
			body {
				background-color: #fff;
				color: #333;
			}
		`,
			expected: map[string]string{},
		},
		{
			name: "parses complex data correctly",
			css: `/*
			Theme Name: My Theme
			Theme URI: http://example.com/my-theme
			Tags: Basis, Framework, HTML5, CSS, JS, Woocommerce, Germanized, AAM, ACF, all the stuff
			Description: A description of my theme
			Link: http://example.com/my-theme
			Text Domain:  something
			Version: 1.0 */`,
			expected: map[string]string{
				"theme name":  "My Theme",
				"theme uri":   "http://example.com/my-theme",
				"description": "A description of my theme",
				"version":     "1.0",
				"link":        "http://example.com/my-theme",
				"tags":        "Basis, Framework, HTML5, CSS, JS, Woocommerce, Germanized, AAM, ACF, all the stuff",
				"text domain": "something",
			},
		},
		{
			name: "parses complex data correctly with license at the end",
			css: `
			/*
				Theme Name: WordPress.org Parent Theme, 2021 edition
				Theme URI: https://github.com/WordPress/wporg-parent-2021
				Author: WordPress.org
				Author URI: https://wordpress.org/
				Description: The WordPress.org Parent Theme is a foundation for themes used on sites in the WordPress.org ecosystem.
				Version: 1.0.0
				License: GNU General Public License v2 or later
				Text Domain: wporg

				WordPress.org Parent Theme, 2021 edition is distributed under the terms of the GNU GPL.
				This theme is based on version 1.1 of the Blockbase theme.

				This program is free software: you can redistribute it and/or modify
				it under the terms of the GNU General Public License as published by
				the Free Software Foundation, either version 2 of the License, or
				(at your option) any later version.

				This program is distributed in the hope that it will be useful,
				but WITHOUT ANY WARRANTY; without even the implied warranty of
				MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
				GNU General Public License for more details.
			*/
			`,
			expected: map[string]string{
				"theme name":  "WordPress.org Parent Theme, 2021 edition",
				"theme uri":   "https://github.com/WordPress/wporg-parent-2021",
				"author":      "WordPress.org",
				"author uri":  "https://wordpress.org/",
				"description": "The WordPress.org Parent Theme is a foundation for themes used on sites in the WordPress.org ecosystem.",
				"version":     "1.0.0",
				"license":     "GNU General Public License v2 or later",
				"text domain": "wporg",
			},
		},
		{
			name: "extra special case",
			css: `/*!
			Theme Name: Example
			Author: Example
			Author URI: https://example.de
			Version: 1.0.0
			Tested up to: 6.1
			Requires PHP: 7.4
			Text Domain: example
			*/@font-face{font-family:"icomoon";src:url(assets/fonts/icomoon.eot?h9gxjf);src:url(assets/fonts/icomoon.eot?h9gxjf#iefix) format("embedded-opentype"),url(assets/fonts/icomoon.ttf?h9gxjf) format("truetype"),url(assets/fonts/icomoon.woff?h9gxjf) format("woff"),url(assets/fonts/icomoon.svg?h9gxjf#icomoon) format("svg")`,
			expected: map[string]string{
				"theme name":   "Example",
				"author":       "Example",
				"author uri":   "https://example.de",
				"version":      "1.0.0",
				"tested up to": "6.1",
				"requires php": "7.4",
				"text domain":  "example",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := analyze.New().ParseCSSThemeString(tt.css)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseThemeCss() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
