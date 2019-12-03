package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type item struct {
	ID      int64
	Title   string
	URL     string
	Image   string
	Text    string
	Created time.Time
	Parent  int64
	Child   int64
	By      string // user
	Kind    string // post || comment
}

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
		ID:      0,
		Title:   r.FormValue("Title"),
		URL:     r.FormValue("URL"),
		Image:   "foo",
		Text:    r.FormValue("Text"),
		Created: time.Now(),
		Parent:  0,
		Child:   0,
		By:      "anon",
		Kind:    "post",
	}

	n := getNext()
	writeItem(i, n)

	tmpl.Execute(w, struct{ Success bool }{true})
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Get the next ID for post or thread
func getNext() int64 {
	var max int64 = 0
	files, err := ioutil.ReadDir("items/")
	check(err)
	for _, f := range files {
		n, err := strconv.ParseInt(f.Name(), 10, 64)
		check(err)
		if n > max {
			max = n
		}
	}
	return max + 1
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/post", postHandler)

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server is listening at %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
