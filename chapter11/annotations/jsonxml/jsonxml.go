package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type Person struct {
	FirstName string `json:"first" xml:"firstName,attr"`
	LastName  string `json:"last" xml:"lastName"`
}

func main() {
	p := &Person{FirstName: "Inigo", LastName: "Montoya"}
	j, _ := json.MarshalIndent(p, "", "  ")
	os.Stdout.Write(j)
	println()

	x, _ := xml.MarshalIndent(p, "", "  ")
	os.Stdout.Write(x)
}
