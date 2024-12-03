package resource

import (
	"net/url"
	"strings"
)

// isRelatedLink is a parser for the given data.
func isRelatedLink(baseURL *url.URL, link string) (bool, error) {
	parseLink, err := url.Parse(link)
	if err != nil {
		return false, err
	}

	if parseLink.Host != "" && parseLink.Host != baseURL.Host {
		return false, nil
	}

	if parseLink.Scheme != "" && parseLink.Scheme != baseURL.Scheme {
		return false, nil
	}

	if parseLink.Path == "" && parseLink.Host == "" {
		return false, nil
	}

	// check if the path is a subpath of the base URL
	if baseURL.Path != "" && !strings.HasPrefix(parseLink.Path, baseURL.Path) {
		return false, nil
	}

	return true, nil
}
