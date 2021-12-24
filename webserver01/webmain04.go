package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	h := logger(http.HandlerFunc(indexHandler))
	err := http.ListenAndServe(":8080", h)
	log.Println(err)
}

type middleware func(http.Handler) http.Handler

func logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("requestURI: %s , path: %s , Query: %s\n", r.RequestURI, r.URL.Path, r.URL.Query())
		t := time.Now()
		h.ServeHTTP(w, r)
		diff := time.Now().Sub(t)
		log.Printf("path: %s, time: %d\n", r.URL.Path, diff/time.Microsecond)
	})
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index page"))
}
