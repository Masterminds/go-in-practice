package main

import (
	"net/http"
)

func main() {
	dir := http.Dir("./files")
	http.ListenAndServe(":8080", http.FileServer(dir))
}
