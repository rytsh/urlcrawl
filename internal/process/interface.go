package process

import "context"

//go:generate mockgen -source=${GOFILE} -destination=interface_test.go -package=${GOPACKAGE} Storage Resource

type Storage interface {
	Set(path string, data []byte) error
	Get(path string) ([]byte, error)
	Has(path string) bool
}

type Resource interface {
	Fetch(ctx context.Context, url string) ([]byte, error)
	Links(data []byte) []string
}
