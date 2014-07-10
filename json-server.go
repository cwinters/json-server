package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	datadir     = ""
	addr        = ""
	port        = ""
	contentType = ""
)

func main() {
	var err error

	flag.StringVar(&datadir, "d", datadir, "root directory for files")
	flag.StringVar(&port, "p", port, "port number on which to listen")

	flag.Parse()

	if port != "" {
		addr = ":" + port
	}

	if addr == "" {
		addr = ":" + os.Getenv("PORT")
	}

	if addr == ":" {
		addr = ":7878"
	}

	if datadir == "" {
		datadir = os.Getenv("DATADIR")
	}

	if datadir == "" {
		datadir, err = os.Getwd()
		if err != nil {
			datadir = "."
		}
	}

	contentType = os.Getenv("CONTENT_TYPE")
	if contentType == "" {
		contentType = "application/json"
	}

	log.Printf("server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, &handler{datadir: datadir}))
}

type handler struct {
	datadir string
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%#v", r.URL)

	rc := h.getData(r.URL.Path)
	if rc == nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "nope\n")
		return
	}

	defer rc.Close()
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", contentType)
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
