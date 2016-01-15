package main

import (
	"log"
	"runtime"
	"time"
)

func monitorRuntime() {
	log.Println("Number of CPUs:", runtime.NumCPU())
	m := &runtime.MemStats{}
	for {
		r := runtime.NumGoroutine()
		log.Println("Number of goroutines", r)

		runtime.ReadMemStats(m)
		log.Println("Allocated memory", m.Alloc)
		time.Sleep(10 * time.Second)
	}
}

func main() {
	go monitorRuntime()

	i := 0
	for i < 40 {
		go func() {
			time.Sleep(15 * time.Second)
		}()
		i = i + 1
		time.Sleep(1 * time.Second)
	}
}
