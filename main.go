//go:build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s%s!", r.URL.Path[8:], r.URL.Path[8:])
}

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/double/", handler2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
