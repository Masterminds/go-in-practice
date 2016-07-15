package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", displayTest)
	http.ListenAndServe(":8080", nil)
}

func displayTest(w http.ResponseWriter, r *http.Request) {
	t := r.Header.Get("Accept")
	var err error
	var b []byte
	var ct string
	switch t {
	case "application/vnd.mytodos.json; version=2.0":
		data := testMessageV2{"Version 2"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=2.0"
	case "application/vnd.mytodos.json; version=1.0":
		fallthrough
	default:
		data := testMessageV1{"Version 1"}
		b, err = json.Marshal(data)
		ct = "application/vnd.mytodos.json; version=1.0"
	}

	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}
	w.Header().Set("Content-Type", ct)
	fmt.Fprint(w, string(b))
}

type testMessageV1 struct {
	Message string `json:"message"`
}

type testMessageV2 struct {
	Info string `json:"info"`
}
