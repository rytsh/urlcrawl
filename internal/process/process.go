package process

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

var DefaultMaxProcess = 10

type Process struct {
	storage  Storage
	resource Resource

	maxProcess int
}

func New(storage Storage, resource Resource, opts ...Option) *Process {
	o := option{}
	for _, opt := range opts {
		opt(&o)
	}

	if o.MaxProcess <= 0 {
		o.MaxProcess = DefaultMaxProcess
	}

	return &Process{
		storage:    storage,
		resource:   resource,
		maxProcess: o.MaxProcess,
	}
}

// Process processes the given URL recursively.
func (p *Process) Process(ctx context.Context, urlStr string) error {
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(p.maxProcess)

	g.Go(func() error {
		return p.process(ctx, g, urlStr)
	})

	if err := g.Wait(); err != nil {
		return err //nolint:wrapcheck // no need
	}

	return nil
}

func (p *Process) process(ctx context.Context, g *errgroup.Group, urlStr string) error {
	var data []byte
	var err error

	urlPath, err := URLToPath(urlStr)
	if err != nil {
		return fmt.Errorf("failed to convert URL to path: %w", err)
	}

	if p.storage.Has(urlPath) {
		data, err = p.storage.Get(urlPath)
		if err != nil {
			return fmt.Errorf("failed to get data: %w", err)
		}
	} else {
		data, err = p.resource.Fetch(ctx, urlStr)
		if err != nil {
			return fmt.Errorf("failed to fetch data: %w", err)
		}

		if err := p.storage.Set(urlPath, data); err != nil {
			return fmt.Errorf("failed to set data: %w", err)
		}
	}

	links := p.resource.Links(data)

	for _, link := range links {
		g.Go(func() error {
			if err := p.process(ctx, g, link); err != nil {
				return fmt.Errorf("failed to process link: %w", err)
			}

			return nil
		})
	}

	return nil
}
