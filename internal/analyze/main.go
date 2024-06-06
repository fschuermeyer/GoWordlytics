package analyze

import (
	"errors"

	"github.com/fschuermeyer/GoWordlytics/internal/report"
	"github.com/fschuermeyer/GoWordlytics/internal/simultan"
	"github.com/fschuermeyer/GoWordlytics/internal/wordpress"
)

var ERR_MALFORMED_URL = errors.New("malformed URL")

func NewReport(url string) (report.Report, error) {
	a := New()

	var r report.Report

	ok := r.SetUrl(url)

	if !ok {
		return r, ERR_MALFORMED_URL
	}

	a.setUrl(url)

	if err := a.setBody(); err != nil {
		return report.Report{}, err
	}

	r.SetHasWordPress(a.isWordpress())

	if !r.HasWordPress() {
		return r, nil
	}

	version := a.version()
	r.SetVersion(version)

	api := wordpress.New(a.data.url, a.userAgent, a.apiVersion)

	simultan.Run([]func(){
		func() {
			resp := api.GetLatestVersion(version)

			if resp.Response != "error" {
				r.SetVersionUpdate(resp.Response, resp.Current)
			}
		},
		func() {
			r.SetHasReadme(a.hasReadme())
		},
		func() {
			r.SetPlugins(a.getPlugins())
		},
		func() {
			r.SetThemes(a.getThemes())
		},
		func() {
			r.SetUsers(api.GetUsers())
		},
	})

	return r, nil
}
