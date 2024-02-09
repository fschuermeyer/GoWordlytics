package report

type Header struct {
	niceName    string
	displayName string
}

type Theme struct {
	Name        string
	URI         string
	Version     string
	Author      string
	AuthorURI   string
	Licence     string
	LicenceURI  string
	Description string
	MinVersion  string
	MaxVersion  string
	RequiredPHP string
	Template    string
	Status      string
	Tags        []string
	TextDomain  string
	DomainPath  string
}

type PluginDetails struct {
	Name                   string         `json:"name"`
	Slug                   string         `json:"slug"`
	Version                string         `json:"version"`
	Author                 string         `json:"author"`
	AuthorProfile          string         `json:"author_profile"`
	Requires               string         `json:"requires"`
	Tested                 string         `json:"tested"`
	RequiresPHP            string         `json:"requires_php"`
	RequiresPlugins        []string       `json:"requires_plugins"`
	Compatibility          []string       `json:"compatibility"`
	Rating                 int            `json:"rating"`
	Ratings                map[string]int `json:"ratings"`
	NumRatings             int            `json:"num_ratings"`
	SupportThreads         int            `json:"support_threads"`
	SupportThreadsResolved int            `json:"support_threads_resolved"`
	Downloaded             int            `json:"downloaded"`
	LastUpdated            string         `json:"last_updated"`
	Added                  string         `json:"added"`
	Homepage               string         `json:"homepage"`
	Description            string         `json:"description"`
	FAQ                    string         `json:"faq"`
	Changelog              string         `json:"changelog"`
	Screenshots            map[string]struct {
		Src     string `json:"src"`
		Caption string `json:"caption"`
	} `json:"screenshots"`
	Tags         map[string]string `json:"tags"`
	Versions     map[string]string `json:"versions"`
	DownloadLink string            `json:"download_link"`
	DonateLink   string            `json:"donate_link"`
	Contributors map[string]string `json:"contributors"`
}
