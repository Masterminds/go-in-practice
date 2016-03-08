package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://mattfarina.com")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r.Body.Close()
	fmt.Println(string(o))

	r2, err := http.Get("http://mattfarina.com/about")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	o, err = ioutil.ReadAll(r2.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	r2.Body.Close()
	fmt.Println(string(o))
}
