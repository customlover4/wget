package service

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"wgetNew/entities/link"
	"wgetNew/entities/queue"
	"wgetNew/utils"
)

const (
	WorkersLimit = 10
)

var (
	ErrWrongHeadersFromPage = errors.New("wrong response headers from page")
	Err0BytesWritten        = errors.New("0 bytes written from body")
	ErrWrongProcessor       = errors.New("we have empty file extension in text processor, need - html, css, etc")
	BreakSignal             = errors.New("break signal")
)

type Service struct {
	visited   map[link.Link]string
	host      string
	scheme    string
	dir       string
	mirroring bool
	qlinks    *queue.Queue
	qfiles    *queue.Queue
}

type Response struct {
	Data    *bytes.Buffer
	Content [2]string
}

type PageElement struct {
	link link.Link
	deep int
}

type FileElement struct {
	name  string
	links []link.Link
}

func NewService(dir string, mirroring bool) *Service {
	return &Service{
		visited:   map[link.Link]string{},
		dir:       dir,
		mirroring: mirroring,
		qlinks:    queue.NewQueue(),
		qfiles:    queue.NewQueue(),
	}
}

func (s *Service) Start(u string, deep int) error {
	entry, err := utils.ParseEntry(u)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s.host = entry.Host
	s.scheme = entry.Scheme

	s.qlinks.Add(&PageElement{
		link: link.Link{
			Formatted: entry.String(),
			FromFile:  entry.Path,
		},
		deep: deep,
	})

	fmt.Println("start downloading files...")
	s.pages()
	s.files()

	return nil
}
