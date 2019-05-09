package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type item struct {
	ID      int64
	Title   string
	URL     string
	Image   string
	Text    string
	Created time.Time
	Parent  int64
	Kid     int64
	By      string // user
	Kind    string // post || comment
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := "kamratposten"
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	srv.ListenAndServe()
}
