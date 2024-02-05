package format

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func extractUserData(value string) (url.Userinfo, bool) {
	if !strings.Contains(value, "@") {
		return url.Userinfo{}, false
	}

	parts := strings.Split(value, "@")

	if len(parts) != 2 {
		return url.Userinfo{}, false
	}

	userdata := parts[0]

	if !strings.Contains(userdata, "://") {
		return url.Userinfo{}, false
	}

	parts = strings.Split(userdata, "://")

	if len(parts) != 2 {
		return url.Userinfo{}, false
	}

	user := strings.Split(parts[1], ":")

	if len(user) > 2 {
		return url.Userinfo{}, false
	}

	if len(user) == 2 {
		return *url.UserPassword(user[0], user[1]), true
	}

	return *url.User(user[0]), true
}

func hasProtocol(value string) bool {
	value = strings.ToLower(value)

	return strings.HasPrefix(value, "http") && strings.Contains(value, "://")
}

func URL(value string) (string, bool) {
	value = strings.TrimSpace(value)

	if !hasProtocol(value) {
		value = fmt.Sprintf("http://%s", value)
	}

	userdata, userok := extractUserData(value)

	for strings.Contains(value, ":///") {
		value = strings.ReplaceAll(value, "///", "//")
	}

	host, err := url.Parse(strings.ToLower(value))

	if err != nil {
		return "", false
	}

	if host.Scheme == "" {
		host.Scheme = "http"
	}

	if host.RawQuery != "" {
		host.RawQuery = ""
	}

	if host.Fragment != "" {
		host.Fragment = ""
	}

	_, icann := publicsuffix.PublicSuffix(host.Host)

	if icann == false {
		return "", false
	}

	if !strings.HasPrefix(host.Scheme, "http") {
		return "", false
	}

	if userok {
		host.User = &userdata
	}

	return host.String(), true
}
