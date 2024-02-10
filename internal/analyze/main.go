package analyze

import (
	"errors"

	"github.com/fschuermeyer/GoWordlytics/internal/report"
)

var ERR_MALFORMED_URL = errors.New("malformed URL")

func NewReport(url string) (report.Report, error) {
	a := New()

	var r report.Report

	r.SetUrl(url)
	a.setUrl(url)
	a.setBody()

	r.SetHasWordPress(a.isWordpress())

	if !r.HasWordPress() {
		return r, nil
	}

	r.SetVersion(a.version())

	r.SetHasReadme(a.hasReadme())

	return r, nil
}
