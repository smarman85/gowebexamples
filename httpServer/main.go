package main

/* 
A basic http servrer has a few key jobs
process dynamic requests
serve static assets (css, js, images)
Accept connections. listen on a port

net/http package contains all utilities to accept/handle requests 
we can register a new handler with 'http.HandleFunc'
1st param is the path to match and a func to execute as a second
ex)
when you hit http://example.com/ you will be greated with a message

http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Welcome to my website!")
})

Serving static assets with
'http.FileServer and point it to a url path.
fs := http.FileServer(http.Dir("static/"))

once in place just need tp point a url path at it. just like
with dynamic requests.
!! To serve files correctly, we need to strip away a part of the url path
usually this is the name of the directory our files live in
http.Handle("/static/", http.StripPrefix("/static/", fs))

listen on a port to accept connections
http.ListenAndServe(":80", nil)

*/

import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my website!")
  })

  fs := http.FileServer(http.Dir("static/"))
  http.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":8088", nil)
}
