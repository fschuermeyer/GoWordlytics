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

	a.setBody(r.GetUrl())

	r.SetHasWordPress(a.isWordpress())

	if !r.HasWordPress() {
		return r, nil
	}

	return r, nil
}
