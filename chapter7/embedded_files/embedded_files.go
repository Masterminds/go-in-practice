package main

import (
	"github.com/GeertJohan/go.rice"
	"net/http"
)

func main() {
	box := rice.MustFindBox("../files/")
	httpbox := box.HTTPBox()
	http.ListenAndServe(":8080", http.FileServer(httpbox))
}
