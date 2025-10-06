package service

import (
	"fmt"
	"os"
	"sync"
	"wgetNew/entities/replacer"
)

func (s *Service) files() {
	files, data := s.popFiles()

	rpl := replacer.NewReplacer(data)

	fc := make(chan string)
	go func() {
		for _, v := range files {
			fc <- v
		}
		close(fc)
	}()

	wg := &sync.WaitGroup{}
	for i := 0; i < WorkersLimit; i++ {
		wg.Add(1)
		go s.fileWorker(wg, fc, rpl)
	}
	wg.Wait()
}

func (s *Service) popFiles() (files []string, data []string) {
	data = make([]string, 0, s.qfiles.Len()*2)
	files = make([]string, 0, s.qfiles.Len())
	for !s.qfiles.Empty() {
		el := (s.qfiles.Pop()).(*FileElement)
		fileName := el.name
		oldLink := el.links

		files = append(files, fileName)
		for _, v := range oldLink {
			old := v.FromFile
			new, ok := s.visited[v]
			if !ok || new == "" {
				new = v.Formatted
			}
			data = append(data, "\""+old+"\"", "\""+new+"\"")
		}
	}

	return
}

func (s *Service) fileWorker(wg *sync.WaitGroup, fc chan string, rpl *replacer.Replacer) {
	defer wg.Done()
	for fName := range fc {
		err := rpl.Do(s.dir, fName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
