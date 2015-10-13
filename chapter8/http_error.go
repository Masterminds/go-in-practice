package main

import "net/http"

func diaplayError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "An Error Occurred", http.StatusForbidden)
}

func main() {
	http.HandleFunc("/", diaplayError)
	http.ListenAndServe(":8080", nil)
}
