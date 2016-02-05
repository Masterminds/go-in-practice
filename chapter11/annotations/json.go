package main

import (
	"encoding/json"
	"fmt"
)

type Name struct {
	First string `json:"firstName" xml:"FirstName"`
	Last  string `json:"lastName,omitempty"`
	Other string `not,even.a=tag`
}

func main() {
	n := &Name{"Inigo", "Montoya", ""}
	data, _ := json.Marshal(n)
	fmt.Printf("%s\n", data)
}
