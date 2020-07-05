package main

import (
	"github.com/pabloos/http/cache"
	"github.com/pabloos/http/greet"
	"log"
	"net/http"
	"os"
)

func newServer() *http.Server {
	messages := make(map[string]greet.Greet)
	c := cache.Cache{
		Messages: messages,
	}
	return &http.Server{
		Addr:      ":8080",
		Handler:   newMux(c),
		TLSConfig: tlsConfig(),
		ErrorLog:  log.New(os.Stderr, "HTTP Server says: ", log.Llongfile),
	}
}
