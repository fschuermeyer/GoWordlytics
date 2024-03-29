package wordpress

import (
	"encoding/json"
	"fmt"

	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

type WordPressUpdate struct {
	Offers []WordPressOffer `json:"offers"`
}

type WordPressOffer struct {
	Response     string `json:"response"`
	Download     string `json:"download"`
	Current      string `json:"current"`
	Version      string `json:"version"`
	PHPVersion   string `json:"php_version"`
	MySQLVersion string `json:"mysql_version"`
}

func (api *API) GetLatestVersion(version string) WordPressOffer {
	var url string

	if len(version) > 0 {
		url = fmt.Sprintf("%s?version=%s", api.apiVersion, version)
	}

	return api.requestLatestVersion(api.userAgent, url)
}

func (api *API) requestLatestVersion(ua, url string) WordPressOffer {
	offer := WordPressOffer{
		Response: "error",
	}

	limit, err := request.CalculateMiB(10)

	if err != nil {
		return offer
	}

	response, err := request.Do(url, ua, limit)

	if err != nil {
		return offer
	}

	var update WordPressUpdate

	err = json.Unmarshal([]byte(response), &update)

	if err != nil || len(update.Offers) == 0 {
		return WordPressOffer{
			Response: "error",
		}
	}

	return update.Offers[0]
}
