package main

import "fmt"

func main() {
	defer goodbye()

	fmt.Println("Hello world.")
}

func goodbye() {
	fmt.Println("Goodbye")
}
