package file

import "os"

type option struct {
	filePermission os.FileMode
	dirPermission  os.FileMode
}

type Option func(*option)

func WithFilePermission(filePermission os.FileMode) Option {
	return func(o *option) {
		o.filePermission = filePermission
	}
}

func WithDirPermission(dirPermission os.FileMode) Option {
	return func(o *option) {
		o.dirPermission = dirPermission
	}
}
