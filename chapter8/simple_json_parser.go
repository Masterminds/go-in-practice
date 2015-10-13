package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

var JSON = `{
  "name": "Miracle Max"
}`

func main() {
	var p Person
	err := json.Unmarshal([]byte(JSON), &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
