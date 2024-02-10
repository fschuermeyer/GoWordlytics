package analyze

import (
	"errors"

	"github.com/fschuermeyer/GoWordlytics/internal/report"
	"github.com/fschuermeyer/GoWordlytics/internal/wordpress"
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

	version := a.version()

	r.SetVersion(version)

	r.SetHasReadme(a.hasReadme())

	resp := wordpress.GetLatestVersion(a.userAgent, a.apiVersion, version)

	if resp.Response != "error" {
		r.SetVersionUpdate(resp.Response, resp.Current)
	}

	return r, nil
}
