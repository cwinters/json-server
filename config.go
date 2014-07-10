package main

import (
	"flag"
	"os"
)

type config struct {
	Port        string
	Addr        string
	DataDir     string
	ContentType string
}

func (c *config) Grok(defaultPort, defaultDatadir string) {
	flag.StringVar(&c.DataDir, "d", defaultDatadir, "root directory for files")
	flag.StringVar(&c.Port, "p", defaultPort, "port number on which to listen")

	flag.Parse()

	c.setAddr(defaultPort)
	c.setDataDir(defaultDatadir)
	c.setContentType()
}

func (c *config) setAddr(defaultPort string) {
	if c.Port != "" {
		c.Addr = ":" + c.Port
	}

	if c.Addr == "" {
		c.Addr = ":" + os.Getenv("PORT")
	}

	if c.Addr == ":" {
		c.Addr = ":" + defaultPort
	}
}

func (c *config) setDataDir(defaultDatadir string) {
	var err error

	if c.DataDir == "" {
		c.DataDir = os.Getenv("DATADIR")
	}

	if c.DataDir == "" {
		c.DataDir, err = os.Getwd()
		if err != nil {
			c.DataDir = defaultDatadir
		}
	}
}

func (c *config) setContentType() {
	c.ContentType = os.Getenv("CONTENT_TYPE")
	if c.ContentType == "" {
		c.ContentType = "application/json"
	}
}
