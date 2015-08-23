package main

import "errors"

func main() {
	errA := errors.New("Error")
	//errB := errors.New("Error")

	if errA != errA {
		panic("Expected same!")
	}
	println("Done")
}
