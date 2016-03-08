package main

import (
	"time"
)

func main() {
	//time.Sleep(5 * time.Second)

	//sleep := time.After(5 * time.Second)
	//<-sleep
	//pp := make(chan string)
	//go pingpong(pp, true)
	//go pingpong(pp, false)
	ch := make(chan bool)
	go ping(ch)
	go pong(ch)
	ch <- true
	select {}
}

func ping(ch chan bool) {
	for {
		select {
		case <-ch:
			println("ping")
			ch <- true
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func pong(ch chan bool) {
	for {
		select {
		case <-ch:
			println("pong")
			ch <- true
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func pingpong(ch chan string, start bool) {
	if start {
		ch <- "ping"
	}

	for {
		select {
		case str := <-ch:
			println(str)
			if str == "ping" {
				ch <- "pong"
				time.Sleep(200 * time.Millisecond)
				continue
			}
			ch <- "ping"
		}
	}
}
