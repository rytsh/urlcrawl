package process

import "net/url"

func URLToPath(urlStr string) (string, error) {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	path := parsedURL.Host + parsedURL.Path

	return path, nil
}
