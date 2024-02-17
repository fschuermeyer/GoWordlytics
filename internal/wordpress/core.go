package wordpress

type API struct {
	url        string
	api        string
	apiVersion string
	userAgent  string
}

func New(url string, userAgent string, apiVersion string) *API {
	return &API{
		url:        url,
		api:        url + "wp-json/",
		apiVersion: apiVersion,
		userAgent:  userAgent,
	}
}
