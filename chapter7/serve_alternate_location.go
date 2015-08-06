package main

import (
	"flag"
	"html/template"
	"net/http"
)

var t *template.Template
var l = flag.String("location", "http://localhost:8080", "A location.")

var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>A Demo</title>
    <link rel="stylesheet" href="{{.Location}}/styles.css">
  </head>
  <body>
  	<p>A demo.</p>
  </body>
</html>`

func init() {
	t = template.Must(template.New("date").Parse(tpl))
}

func servePage(res http.ResponseWriter, req *http.Request) {
	data := struct{ Location *string }{
		Location: l,
	}
	t.Execute(res, data)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", servePage)
	http.ListenAndServe(":8080", nil)
}
