package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(mux))
	log.Println(err)
}

func mux(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/":
		indexHandler(w, r)
	case "/about":
		aboutHanderler(w, r)
	default:
		http.NotFound(w, r)

	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("index page"))
}

func aboutHanderler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about page"))
}
