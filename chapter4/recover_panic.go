package main

import (
	"errors"
	"fmt"
)

func main() {
	msg := "Everything's fine"
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Trapped panic: %s (%T)\n", err, err)
		}
		fmt.Println(msg)
	}()

	yikes()
}

func yikes() {
	panic(errors.New("Something bad happened."))
}
