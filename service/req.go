package service

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"wgetNew/entities/safe"
	"wgetNew/external/web"
	"wgetNew/utils"
)

func (s *Service) page(u string) (*Response, error) {
	resp, err := web.Req(u)
	if err != nil {
		return nil, err

	}
	defer resp.Body.Close()

	r := &Response{}

	tmp := strings.Split(
		strings.Split(
			resp.Header.Get("Content-Type"), ";",
		)[0],
		"/",
	)
	if len(tmp) < 2 {
		r.Content = [2]string{}
	}
	r.Content = [2]string{tmp[0], tmp[1]}

	// buffer with mutex on every operation
	b := safe.NewBuffer()

	contentLen, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	fmt.Printf("downloading: %s\n", resp.Request.URL.String())
	if err == nil && contentLen/1024/1024 > 10 {
		go utils.DownloadProgress(b, contentLen)
	}

	written, err := io.Copy(b, resp.Body)
	if written == 0 {
		return nil, Err0BytesWritten
	}

	r.Data = b.Buf()
	b = nil

	return r, nil
}
