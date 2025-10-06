package web

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

var (
	ErrStatusCode = errors.New("status code is not 200 (not OK)")
)

func Req(url string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Content-Type", "charset=utf-8")

	client := &http.Client{
		Timeout: 30 * time.Minute,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second, // Таймаут рукопожатия TLS
			ResponseHeaderTimeout: 10 * time.Second, // Таймаут получения заголовков
			ExpectContinueTimeout: 5 * time.Second,
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		b := new(bytes.Buffer)
		io.Copy(b, resp.Body)
		return nil, fmt.Errorf("%w: %d (%s)", ErrStatusCode, resp.StatusCode, url)
	}

	return resp, nil
}
