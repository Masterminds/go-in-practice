package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	ct := "application/vnd.mytodos.json; version=2.0"
	req, _ := http.NewRequest("GET", "http://localhost:8080/test", nil)
	req.Header.Set("Accept", ct)
	res, _ := http.DefaultClient.Do(req)
	if res.Header.Get("Content-Type") != ct {
		fmt.Println("Unexpected content type returned")
		return
	}
	b, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", b)
}
