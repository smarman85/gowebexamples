package main

/*
GO's html/template package provides a templating language for HTML templates
mostly used in web applications to display data
one great benefit of go's templating lang is the automatic escaping of data
no need to worry about XSS as go parses HTML and escapes all inputs before 
displaying it to the browser

First Template
Writing a template in go is simple
this ex shows a TODO list written as unordered list (ul) in HTML
when rendering templates the data passed can be any kind of Go's data structures
string, numer or nested data struct as in the following. To access the data in a template
the top mose var is accessed by '{{.}}'. The dot inside the curly braces is called a 
pipeline and the root element of the data
data := TodoPageData{
    PageTitle: "My TODO list",
    Todos: []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
    },
}
<h1>{{.PageTitle}}</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>

Controll Structures:
overview of the most commonly used ones:
for a detailed list visit text/template
Controll structure           Definition
{{/* a comment *\/}}         Defins a comment (escape added here to not break currnent comment
{{.}}                        Reders the root element
{{.Title}}                   Renders the "Title"-field in a nested element
{{if .Done}} {{else}}        Defines an if-Statement
{{end}}                      
{{range .Todo}} {{.}}        Loops over all "Todos" and renders each using {{.}}
{{end}}
{{block "content" .}}        Defines a block with the name "content"
{{end}}

Parsing Templates from files
A template can either be parsed from a string or a file on disk.
As it is usually the case, that templates are pares from disk, this ex shows how to 
do do. This ex there is a template fire in the same dir as the Go program called layout.html

tmpl, err := template.ParseFiles("layout.html")
// or
tmpl := template.Must(template.ParseFiles("layout.html"))

Execute a template in a request handler
onse a template is parsed from disk it's ready to be used in a the request handler
the 'Execute' func accepts an 'io.Writer' for writing out the template and an 'interface{}' to
pass data into the template. When the function is called on an http.ResponseWriter the
content-type is header is automatically set in the HTTP response to 
Content-Type: text/html; charset=utf-8
func(w http.ResponseWriter, r *http.Request) {
  tmpl.Execute(w, "data goes here")
}
*/

import (
  "html/template"
  "net/http"
)

type Todo struct {
  Title string
  Done bool
}

type TodoPageData struct {
  PageTitle string
  Todos []Todo
}

func main() {
  tmpl := template.Must(template.ParseFiles("layout.html"))
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    data := TodoPageData{
      PageTitle: "My TODO list",
      Todos: []Todo{
        {Title: "Task 1", Done: false},
        {Title: "Task 2", Done: true},
        {Title: "Task 3", Done: true},
      },
    }
    tmpl.Execute(w, data)
  })
  http.ListenAndServe(":8088", nil)
}








































