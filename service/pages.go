package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"
	"wgetNew/entities/link"
	"wgetNew/external/fs"
	"wgetNew/utils"
)

type out struct {
	links []link.Link
	link  link.Link
	deep  int
	name  string
}

// worker on pages - req, parse, save
// pop i(if < WorkersLimit, pop q.Len elems
// if > WorkersLimit pop WorkersLimit elems)
// and run i workers
func (s *Service) pages() {
	wg := &sync.WaitGroup{}

	for !s.qlinks.Empty() {
		l := s.qlinks.Len()
		i := 0
		if l > WorkersLimit {
			i = WorkersLimit
		} else {
			i = l
		}

		o := make(chan out, i)
		for j := 0; j < i; j++ {
			if !s.qlinks.Empty() {
				el := (s.qlinks.Pop()).(*PageElement)
				if _, ok := s.visited[el.link]; ok {
					continue
				}
				wg.Add(1)
				go s.pageWorker(wg, el, o)
			}
		}
		wg.Wait()

		close(o)
		for tmp := range o {
			if tmp.name == "" {
				continue 
				// we dont be here if have err, that mean only one
				// what we didnt make new file (didnt parse this link, didnt generate name)
			}

			tmp.links = link.Unique(tmp.links)
			s.visited[tmp.link] = tmp.name
			for _, l := range tmp.links {
				if _, ok := s.visited[l]; ok {
					continue
				}
				s.qlinks.Add(&PageElement{l, tmp.deep - 1})
			}
			if len(tmp.links) != 0 {
				s.qfiles.Add(&FileElement{tmp.name, tmp.links})
			}
		}
	}
}

// dont block, cause we write i elems and close c in popElems()
func (s *Service) pageWorker(wg *sync.WaitGroup, el *PageElement, o chan out) {
	defer wg.Done()

	links, name, err := s.processPage(el.link.Formatted, el.deep)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	o <- out{links, el.link, el.deep, name}
}

// (html || css) - req -> parse -> save
// other - req -> save
func (s *Service) processPage(u string, deep int) ([]link.Link, string, error) {
	var err error

	resp, err := s.page(u)
	if err != nil {
		return nil, "", err
	}
	if resp.Content[0] == "" {
		return nil, "", ErrWrongHeadersFromPage
	}

	ext := utils.FileExtension(resp.Content)
	name := utils.NumeredName() + ext

	wasHtml := false
	links := make([]link.Link, 0, 10)
	switch resp.Content[1] {
	case "html":
		wasHtml = true
		if deep < 0 {
			return nil, "", nil
		}

		links, err = s.parseHTML(resp)
		if err != nil {
			return nil, "", err
		}
	case "css":
		links = s.parseCSS(resp.Data.String())
	}

	if !wasHtml && s.mirroring {
		name = "static/" + name
	}
	err = fs.Write(s.dir, name, bytes.NewReader(resp.Data.Bytes()))
	if err != nil {
		return nil, "", err
	}

	return links, name, nil
}
