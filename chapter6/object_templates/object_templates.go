package main

import (
	"bytes"
	"html/template"
	"net/http"
)

var t *template.Template
var qc template.HTML

func init() {
	t = template.Must(template.ParseFiles("index.html", "quote.html"))
}

type Page struct {
	Title   string
	Content template.HTML
}

type Quote struct {
	Quote, Person string
}

func main() {
	q := &Quote{
		Quote: `You keep using that word. I do not think
				it means what you think it means.`,
		Person: "Inigo Montoya",
	}
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "quote.html", q)
	qc = template.HTML(b.String())

	http.HandleFunc("/", diaplayPage)
	http.ListenAndServe(":8080", nil)
}

func diaplayPage(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title:   "A Quote",
		Content: qc,
	}
	t.ExecuteTemplate(w, "index.html", p)
}
