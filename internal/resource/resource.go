package resource

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/worldline-go/klient"
	"golang.org/x/net/context"
)

type Resource struct {
	client *klient.Client

	baseURL *url.URL
}

// New creates a new HTTP Resource.
//   - Accepts a baseURL string like "https://example.com".
func New(baseURL string) (*Resource, error) {
	client, err := klient.New(
		klient.WithBaseURL(baseURL),
		klient.WithLogger(slog.Default()),
	)
	if err != nil {
		return nil, err
	}

	baseURLParse, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Resource{
		client:  client,
		baseURL: baseURLParse,
	}, nil
}

// Fetch fetches the data from the given URL.
//   - Accepts 2xx status codes.
func (r *Resource) Fetch(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	var data []byte
	if err := r.client.Do(req, func(r *http.Response) error {
		if err := klient.UnexpectedResponse(r); err != nil {
			return err
		}

		var err error
		data, err = io.ReadAll(r.Body)
		if err != nil {
			// ignore EOF error
			if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return data, nil
}

// Links returns the <a> href links from the given data.
func (r *Resource) Links(data []byte) []string {
	var links []string

	Parser(data).ParseLinks(func(link string) {
		ok, err := isRelatedLink(r.baseURL, link)
		if err != nil {
			slog.Error("failed to check related link", slog.String("error", err.Error()))

			return
		}

		if ok {
			links = append(links, link)
		}
	})

	return links
}
