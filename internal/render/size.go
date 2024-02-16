package render

import (
	"golang.org/x/term"
)

func getWidth(fallback int) int {
	if !term.IsTerminal(0) {
		return fallback
	}

	width, _, err := term.GetSize(0)

	if err != nil {
		return fallback
	}

	return width
}
