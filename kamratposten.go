package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

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
