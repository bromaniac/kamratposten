package main

import ( 
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "log"
    "os"
)

func HomeHandler() {
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", HomeHandler)

    srv := &hhtp.Server{
        Addr:   "0.0.0.0:8080",
        WriteTimeout: time.Seconds * 15,
        ReadTimeout:  time.Seconds * 15,
        IdleTimeout:  time.Seconds * 60,
        Handler: r,
    }
    os.Exit(0)
}
