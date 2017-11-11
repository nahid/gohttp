package gohttp

import "net/http"

// Option is an interface of request option
type Option interface {
	apply(*Request)
}

// OptionFunc is an implementation of option interface
type OptionFunc func(*Request)

func (fn OptionFunc) apply(r *Request) {
	fn(r)
}

// SetClient option sets client c for request
func SetClient(c *http.Client) OptionFunc {
	return func(r *Request) {
		r.Client = c
	}
}
