package process

type option struct {
	MaxProcess int
}

type Option func(*option)

func WithMaxProcess(maxProcess int) Option {
	return func(o *option) {
		o.MaxProcess = maxProcess
	}
}
