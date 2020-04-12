package main

import (
  "fmt"
  "net/http"
)

// func (w http.ResponseWriter, r *http.Request)
// http.ResponseWriter where you write your html response to
// http.Request contains all info about this HTTP request (URLS, Headers)


func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
  })

  // Listen for connections 
  http.ListenAndServe(":8088", nil)
}
