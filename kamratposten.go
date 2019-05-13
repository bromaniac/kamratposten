package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

func init() {
	if _, err := os.Stat("items"); os.IsNotExist(err) {
		os.Mkdir("items", 0755)
	}
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		os.Mkdir("images", 0755)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := "kamratposten"
	tmpl := template.Must(template.ParseFiles("views/index.html"))

	tmpl.Execute(w, data)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("views/submitPage.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	i := item{
		ID:      1,
		Title:   r.FormValue("Title"),
		URL:     r.FormValue("URL"),
		Image:   "foo",
		Text:    r.FormValue("Text"),
		Created: time.Now(),
		Parent:  0,
		Kid:     0,
		By:      "first",
		Kind:    "post",
	}

	writeItem(i, 1)

	tmpl.Execute(w, struct{ Success bool }{true})
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
	r.HandleFunc("/post", postHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	srv.ListenAndServe()
}
