package main

/*
ex to show how to serve static files like css, js or images from a dir
*/

import (
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("assets/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))
  http.ListenAndServe(":8088", nil)
}
