package exec

import (
	"io"
	"os"
)

type Option interface {
	apply(*options)
}

func Out(w io.Writer) Option {
	return optionFn(func(o *options) { o.stdout = w })
}

func Err(w io.Writer) Option {
	return optionFn(func(o *options) { o.stderr = w })
}

func In(r io.Reader) Option {
	return optionFn(func(o *options) { o.stdin = r })
}

type optionFn func(*options)

func (f optionFn) apply(o *options) { f(o) }

type options struct {
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

func newOptions(opts []Option) *options {
	newOpts := new(options)
	for _, opt := range opts {
		opt.apply(newOpts)
	}

	return newOpts
}

func (o *options) out() io.Writer {
	if o.stdout != nil {
		return io.MultiWriter(os.Stdout, o.stdout)
	}

	return os.Stdout
}

func (o *options) err() io.Writer {
	if o.stderr != nil {
		return io.MultiWriter(os.Stderr, o.stderr)
	}

	return os.Stderr
}

func (o *options) in() io.Reader {
	if o.stdin != nil {
		return o.stdin
	}

	return os.Stdin
}
