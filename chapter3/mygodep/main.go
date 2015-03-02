package main

import (
	"fmt"
	"github.com/Masterminds/cookoo"
)

func main() {
	fmt.Println("Hello Godep!")
	_, _, cxt := cookoo.Cookoo()

	cxt.Log("info", "OH HAI")
}
