package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about", aboutHanderler)
	err := http.ListenAndServe(":8080", mux)
	log.Println(err)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("index page"))
}

func aboutHanderler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about page"))

}
