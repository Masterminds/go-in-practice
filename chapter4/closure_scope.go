package main

import "fmt"

func main() {
	var msg string
	defer func() {
		fmt.Println(msg)
	}()
	msg = "Hello world"
}
