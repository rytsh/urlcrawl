package file

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	DefaultFilePermission os.FileMode = 0o644
	DefaultDirPermission  os.FileMode = 0o755

	DefaultExtension = ".data"
)

type File struct {
	destionation string

	filePermission os.FileMode
	dirPermission  os.FileMode
}

func New(destionation string, opts ...Option) (*File, error) {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}

	if o.filePermission == 0 {
		o.filePermission = DefaultFilePermission
	}

	if o.dirPermission == 0 {
		o.dirPermission = DefaultDirPermission
	}

	// create destination directory
	if err := os.MkdirAll(destionation, o.dirPermission); err != nil {
		return nil, fmt.Errorf("failed to create destination directory: %w", err)
	}

	return &File{
		destionation:   destionation,
		filePermission: o.filePermission,
		dirPermission:  o.dirPermission,
	}, nil
}

func (f *File) Set(path string, data []byte) error {
	filePath := filepath.Join(f.destionation, path) + DefaultExtension

	if err := os.MkdirAll(filepath.Dir(filePath), f.dirPermission); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	if err := os.WriteFile(filePath, data, f.filePermission); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func (f *File) Get(path string) ([]byte, error) {
	filePath := filepath.Join(f.destionation, path) + DefaultExtension

	return os.ReadFile(filePath)
}

func (f *File) Has(path string) bool {
	filePath := filepath.Join(f.destionation, path) + DefaultExtension

	if _, err := os.Stat(filePath); err == nil {
		return true
	}

	return false
}
