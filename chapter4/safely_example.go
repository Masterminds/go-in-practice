package main

import (
	"./safely"
	"errors"
	"time"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("Oops!"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(200)
}
