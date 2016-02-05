package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl = `package {{.Package}}

type {{.MyType}}Queue struct {
	q []{{.MyType}}
}

func New{{.MyType}}Queue() *{{.MyType}}Queue {
	return &{{.MyType}}Queue{
		q: []{{.MyType}}{},
	}
}

func (o *{{.MyType}}Queue) Insert(v {{.MyType}}) {
	o.q = append(o.q, v)
}

func (o *{{.MyType}}Queue) Remove() {{.MyType}} {
	if len(o.q) == 0 {
		panic("Oops.")
	}
	first := o.q[0]
	o.q = o.q[1:]
	return first
}
`

func main() {
	tt := template.Must(template.New("queue").Parse(tpl))
	for i := 1; i < len(os.Args); i++ {
		dest := strings.ToLower(os.Args[i]) + "_queue.go"
		file, err := os.Create(dest)
		if err != nil {
			fmt.Printf("Could not create %s: %s (skip)\n", dest, err)
			continue
		}

		vals := map[string]string{
			"MyType":  os.Args[i],
			"Package": os.Getenv("GOPACKAGE"),
		}
		tt.Execute(file, vals)

		file.Close()
	}
}
