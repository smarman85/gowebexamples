package main

/*
net/http package doesn't do complex routing very well 
ie) segmenting a request url into single parameters
gorilla/mux package is used to create routes with named parameters
GET/POST handlers and domain restrictions

gorilla/mux
package that adapts to go's default HTTP router
increase productivity while writing web apps
compliant to go's default request handler signature
func (w http.ResponseWriter, r *http.Request)
pacakge can be mixed and matched with other HTTP libraries
go get -u github.com/gorilla/mux

Create a new router:
the router is the main router for your web app and will be later passed as a
param to the server
it will receive all http connections and pass it on to the req handlers
you will register on it like so 
r := mux.NewRouter()

requesting a request handler
once you have a router you can register handers like usual
only diff instead of calling http.HandleFunc(...) you call 
handleFunc on your router like this
r.HandleFunc(...)

URL params
gorilla/mux biggest strenght lies in the ability to extract segments from the 
request url
/books/go-programming-blueprint/page/10
this url has 2 dynamic sections
book title slug go-programing-blueprint
page 10
to match the url mentioned above you replace the dynamic segments with placeholders
in your url patter like so

r.HandleFunc("/books/{title}/page/{page}", func(w, http.ResponseWriter, r *http.Request) {
  // get the book
  // navigate to the page
}

the last thing is to get the data from the segments. The package comes with the func
mux.Vars(r) which takes the http.Request as a parameter and returns a map of the segments
func(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  vars["title"] // book title slug
  vars["page"] // the page
}

setting the http servers router
wonder what the 'nil' in 'http.ListenAndServe(":80", nil) meant
it's the param for the main router of the  http server
default is nil which means to use the default router of the net/http package
To make use of your own router, replace the 'nil' with the var of your router r
http.ListenAndServe(":80", r)

Methods
Restrict the request handler to specific HTTP methods.

r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

Hostnames & Subdomains
Restrict the request handler to specific hostnames or subdomains.

r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com")

Schemes
Restrict the request handler to http/https.

r.HandleFunc("/secure", SecureHandler).Schemes("https")
r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

Path Prefixes & Subrouters
Restrict the request handler to specific path prefixes.

bookrouter := r.PathPrefix("/books").Subrouter()
bookrouter.HandleFunc("/", AllBooks)
bookrouter.HandleFunc("/{title}", GetBook)
*/

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
  })

  r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello\n")
  })

  http.ListenAndServe(":8088", r)
}
