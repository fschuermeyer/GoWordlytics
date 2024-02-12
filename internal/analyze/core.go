package analyze

import (
	"fmt"

	"github.com/fschuermeyer/GoWordlytics/internal/format"
	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

func (a *Analyze) setUrl(url string) {
	url, ok := format.URL(url)

	if ok {
		a.data.url = url
	}
}

func (a *Analyze) setBody() error {
	size, err := request.CalculateMiB(2)

	if err != nil {
		return err
	}

	body, err := request.Do(a.data.url, a.userAgent, size)

	if err == request.ERR_STATUS_NOT_OK {
		return err
	}

	a.data.htmlIndex = body

	return nil
}

func (a *Analyze) getContent(path string, miblimit int64) string {
	limit, err := request.CalculateMiB(miblimit)

	if err != nil {
		return ""
	}

	resp, err := request.Do(fmt.Sprintf("%s%s", a.data.url, path), a.userAgent, limit)

	if err != nil {
		return ""
	}

	return resp
}
