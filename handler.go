package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type handler struct {
	datadir     string
	contentType string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v", r.URL)

	rc := h.getData(r.URL.Path)
	if rc == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "nope\n")
		return
	}

	defer rc.Close()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", h.contentType)
	io.Copy(w, rc)
}

func (h *handler) getData(name string) io.ReadCloser {
	fd, err := os.Open(path.Join(h.datadir, name))
	if err != nil {
		log.Printf("error: %v", err)
		return nil
	}

	return fd
}
