package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)
	// http.HandleFunc("/-/", fileServerHandler)
	http.Handle("/-/", http.StripPrefix("/-", http.FileServer(noDir{http.Dir("public")})))
	http.ListenAndServe(":8080", nil)
}

// func fileServerHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Println(r.URL.Path)
// 	h := http.FileServer(http.Dir("public"))
// 	http.StripPrefix("/-/", h).ServeHTTP(w, r)
// }

type noDir struct {
	http.Dir
}

func (d noDir) Open(name string) (http.File, error) {
	f, err := d.Dir.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, os.ErrNotExist
	}
	return f, nil
}

type indexData struct {
	Name string
	List []string
}

var t = template.Must(template.ParseFiles("index.tmpl"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html ; charset=utf-8")

	// t = t.Funcs(template.FuncMap{
	// 	"incr": func(i int) int {
	// 		return i + 1
	// 	},
	// })

	data := indexData{
		Name: "GolfGrab",
		List: []string{
			"GO",
			"C",
			"CPP"},
	}

	err := t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

}
