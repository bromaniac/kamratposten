package main

import (
	"html/template"
	"log"
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

var users = make(map[string]string)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := "kamratposten"
	tpl := index()

	t, err := template.New("index").Parse(tpl)
	check(err)

	err = t.Execute(w, data)
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	users["test_user"] = "$2a$14$ajq8Q7fbtFRQvXpdCq7Jcuy.Rx1h/L4J60Otx.gyNLbAYctGMJ9tK" // secret

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
