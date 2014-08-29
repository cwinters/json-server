package main

import (
	"log"
	"net/http"
)

var (
	defaultDatadir = "."
	defaultPort    = "7878"
)

func main() {
	cfg := &config{}
	cfg.Grok(defaultPort, defaultDatadir)
	log.Printf("START [Listen: %s] [Directory: %s]", cfg.Addr, cfg.DataDir)
	log.Fatal(http.ListenAndServe(cfg.Addr, &handler{
		datadir:     cfg.DataDir,
		contentType: cfg.ContentType,
	}))
}
