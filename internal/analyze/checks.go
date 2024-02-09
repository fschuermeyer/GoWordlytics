package analyze

import (
	"strings"

	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

func (a *Analyze) setBody(url string) error {
	size, err := request.CalculateMiB(2)

	if err != nil {
		return err
	}

	body, err := request.Do(url, a.userAgent, size)

	if err == request.ERR_STATUS_NOT_OK {
		return err
	}

	a.data.htmlIndex = body

	return nil
}

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
