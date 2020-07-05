package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pabloos/http/cache"
	"github.com/pabloos/http/greet"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func POST(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func Debug(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(dump))
	}
}

func Delay(delay time.Duration, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer h.ServeHTTP(w, r)

		time.Sleep(delay)
	}
}

func Cached(h http.HandlerFunc) http.HandlerFunc {
	messages := make(map[string]greet.Greet)
	c := cache.Cache{
		Messages: messages,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error in body")
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		defer h.ServeHTTP(w, r)

		var t greet.Greet
		var gr greet.Greet
		var found bool

		json.NewDecoder(r.Body).Decode(&t)
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		gr, found = c.Get(t)
		if found {
			fmt.Fprintf(w, "%s, from %s is in the cache\n", gr.Name, gr.Location)
			fmt.Fprintf(w, "Actual Cache %s\n", c.GetMessages())
		} else {
			fmt.Fprintf(w, "%s, from %s not found in cache\n", t.Name, t.Location)
			fmt.Fprintf(w, "Actual Cache %s\n", c.GetMessages())
			c.Set(t)
		}
	}
}
