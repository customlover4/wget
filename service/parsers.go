package service

import (
	"bytes"
	"net/url"
	"strings"
	"wgetNew/entities/link"
	"wgetNew/utils"

	"golang.org/x/net/html"
)

func (s *Service) parseCSS(data string) []link.Link {
	r := make([]link.Link, 0)
	sub := "url(\""
	for idx := strings.Index(data, sub); idx != -1; {
		url := ""
		i := idx + len(sub)
		for ; data[i] != '"'; i++ {
			url += string(data[i])
		}
		u, err := utils.ParseUrl(url, s.host, s.scheme)
		if err != nil {
			continue
		}
		r = append(r, link.Link{
			Formatted: u.String(),
			FromFile:  url,
		})
		data = data[i:]
		idx = strings.Index(data, sub)
	}

	return r
}

func (s *Service) parseHTML(resp *Response) ([]link.Link, error) {
	res := make([]link.Link, 0)

	z, err := html.Parse(bytes.NewReader(resp.Data.Bytes()))
	if err != nil {
		return res, err
	}

	b := new(bytes.Buffer)
	err = html.Render(b, z)
	if err != nil {
		return res, err
	}

	resp.Data = b

	tz := html.NewTokenizer(bytes.NewReader(resp.Data.Bytes()))
	for {
		tt := tz.Next()
		if tt == html.ErrorToken {
			break
		}
		t := tz.Token()

		attrs := t.Attr
		for _, attr := range attrs {
			if attr.Key == "style" {
				res = append(res, s.parseCSS(attr.Val)...)
				continue
			}
			if childU, ok := s.validate(attr.Key, attr.Val); ok {
				res = append(res, link.Link{
					Formatted: childU.String(),
					FromFile:  attr.Val,
				})
			}
		}
	}

	return res, nil
}

func (s *Service) validate(key, val string) (*url.URL, bool) {
	if !utils.IsLink(key) {
		return nil, false
	}
	if strings.Contains(val, "#") {
		return nil, false
	}

	childU, err := utils.ParseUrl(val, s.host, s.scheme)
	if err != nil {
		return nil, false
	}
	if childU.Host != s.host {
		return nil, false
	}

	return childU, true
}
