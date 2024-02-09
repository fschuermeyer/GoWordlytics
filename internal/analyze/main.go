package analyze

import (
	"errors"

	"github.com/fschuermeyer/GoWordlytics/internal/format"
	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

var ERR_MALFORMED_URL = errors.New("malformed URL")
var config Configuration

func init() {
	config = Configuration{
		userAgent: "GoWordlytics",
	}
}

func NewReport(url string) (Report, error) {
	var currentReport Report
	url, ok := format.URL(url)

	if !ok {
		return Report{}, ERR_MALFORMED_URL
	}

	currentReport.Url = url

	size, err := request.CalculateMiB(2)

	if err != nil {
		return Report{}, err
	}

	body, err := request.Do(url, config.userAgent, size)

	if err != nil {
		return Report{}, err
	}

	if len(body) == 0 {
		currentReport.IsWordPress = false
		currentReport.Status = "No WordPress Website"

		return currentReport, nil
	}

	return Report{}, nil
}
