package gohttp

import (
	"net/http"
	"time"
)

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

// SetTransport option sets Transport t for request
func SetTransport(t http.Transport) OptionFunc {
	return func(r *Request) {
		r.Transport = t
	}
}

// SetCookieJar option sets cookie c for request
func SetCookieJar(c http.CookieJar) OptionFunc {
	return func(r *Request) {
		r.Cookie = c
	}
}

// SetTimeout option sets timeout t for request
func SetTimeout(t time.Duration) OptionFunc {
	return func(r *Request) {
		r.Timeout = t
	}
}
