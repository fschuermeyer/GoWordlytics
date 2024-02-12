package analyze

import (
	"strings"
)

func (a *Analyze) isWordpress() bool {
	if len(a.data.htmlIndex) == 0 {
		return false
	}

	for _, indicator := range a.indicators {
		if strings.Contains(a.data.htmlIndex, indicator) {
			return true
		}
	}

	return false
}

func (a *Analyze) hasReadme() bool {
	content := a.getContent("readme.html", 1)

	if len(content) == 0 {
		return false
	}

	return strings.Contains(content, a.IndicatorsReadme.indicator)
}
