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
	if r.URL.Path == "/favicon.ico" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
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
	givenPath := path.Join(h.datadir, name)
	filePaths := []string{
		givenPath,
		path.Join(h.datadir, name+".json"),
		path.Join(path.Dir(givenPath), "_default.json"),
	}
	for _, filePath := range filePaths {
		fd, _ := os.Open(filePath)
		if fd != nil {
			log.Printf("%s => %s", name, filePath)
			return fd
		}
	}
	return nil
}
