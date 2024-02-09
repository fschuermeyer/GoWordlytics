package request

import (
	"errors"
	"net/http"
)

var ERR_STATUS_NOT_OK = errors.New("status not ok")

func Do(url string, userAgent string, limit int64) (string, error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", userAgent)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", ERR_STATUS_NOT_OK
	}

	defer res.Body.Close()

	body, err := ReadLimitedBytesWithSizeCheck(res.Body, limit)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
