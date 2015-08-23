package main

import (
	"errors"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func handler(res http.ResponseWriter, req *http.Request) {
	panic(errors.New("Fake panic!"))
}
