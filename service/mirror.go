package service

import (
	"fmt"
	"net/http"
	"os"
	"wgetNew/external/fs"
)

func (s *Service) Mirror() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.hand)
	mux.Handle(
		"/static/", http.StripPrefix(
			"/static/", http.FileServer(http.Dir(s.dir+"/"+"static")),
		),
	)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("\nstart mirroring")

	if err := server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func (s *Service) hand(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path
	if name == "/" {
		name = "0.html"
	} else if len(name) != 0 && name[0] == '/' {
		name = name[1:]
	}

	path := ""
	if s.dir != "." {
		path = s.dir + "/" + name
	} else {
		path = name
	}
	f, err := fs.ReadString(path)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Write([]byte(f))
}
