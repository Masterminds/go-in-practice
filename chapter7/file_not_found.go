package main

import (
	"fmt"
	fs "github.com/Masterminds/go-fileserver"
	"net/http"
)

func main() {

	fs.NotFoundHandler = func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, "The requested page could not be found.")
	}

	dir := http.Dir("./files")
	http.ListenAndServe(":8080", fs.FileServer(dir))
}
