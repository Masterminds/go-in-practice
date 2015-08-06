package main

import (
	"html/template"
	"net/http"
	"time"
)

var tpl = `<!DOCTYPE HTML>
<html>
  <head>
    <meta charset="utf-8">
    <title>Date Example</title>
  </head>
  <body>
  	<p>{{.Date | dateFormat "Jan 2, 2006"}}</p>
  </body>
</html>`

var funcMap = template.FuncMap{
	"dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}

func serveTemplate(res http.ResponseWriter, req *http.Request) {
	t := template.New("date")
	t.Funcs(funcMap)
	t.Parse(tpl)
	data := struct{ Date time.Time }{
		Date: time.Now(),
	}
	t.Execute(res, data)
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":8080", nil)
}
