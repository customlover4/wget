package utils

import (
	"errors"
	"net/url"
)

var (
	ErrWrongURLFormat = errors.New("format for url: scheme://host/path")
)

func ParseUrl(u, host, scheme string) (*url.URL, error) {
	r, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	if r.Scheme == "" {
		r.Scheme = scheme
	}
	if r.Host == "" {
		r.Host = host
	}
	if r.Path == "" {
		r.Path = "/"
	}

	return r, nil
}

func ParseEntry(u string) (*url.URL, error) {
	entry, err := url.Parse(u)
	if err != nil {
		return entry, err
	}
	if entry.Host == "" || entry.Scheme == "" {
		return entry, ErrWrongURLFormat
	}
	if entry.Path == "" {
		entry.Path = "/"
	}

	return entry, err
}
