package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := checkDep("fortune")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Time to get your fortunte")
}

func checkDep(name string) error {
	if _, err := exec.LookPath(name); err != nil {
		es := "Could not find '%s' in PATH: %s"
		return fmt.Errorf(es, name, err)
	}

	return nil
}
