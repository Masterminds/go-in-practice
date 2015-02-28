package main

import (
	"fmt"
)

func Names() (first string, second string) {
	first = "Foo"
	second = "Bar"
	return
}

func main() {
	n1, n2 := Names()
	fmt.Println(n1, n2)
}
