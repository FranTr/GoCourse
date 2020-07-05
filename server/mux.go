package main

import (
	"github.com/pabloos/http/cache"
	"net/http"
)

func newMux(c cache.Cache) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Debug(index))
	mux.HandleFunc("/greet", Cached(c, POST(greetHandler)))
	return mux
}
