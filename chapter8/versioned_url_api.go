package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type testMessage struct {
	Message string `json:"message"`
}

func diaplayTest(w http.ResponseWriter, r *http.Request) {
	data := testMessage{"A test message."}
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(b))
}

func main() {
	http.HandleFunc("/api/v1/test", diaplayTest)
	http.ListenAndServe(":8080", nil)
}
