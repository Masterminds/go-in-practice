package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("./templates/simple.html"))
}

type Page struct {
	Title, Content string
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "An Example",
		Content: "Have fun stormin’ da castle.",
	}
	var b bytes.Buffer
	err := t.Execute(&b, p)
	if err != nil {
		fmt.Fprint(w, "A error occured.")
		return
	}
	b.WriteTo(w)
}

func main() {
	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}
