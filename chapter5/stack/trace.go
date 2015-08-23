package main

import (
	"fmt"
	"runtime"
)

func main() {
	foo()
}

func foo() {
	bar()
}

func bar() {
	buf := make([]byte, 1024)
	runtime.Stack(buf, false)
	fmt.Printf("Trace:\n %s\n", buf)

	caller, file, line, _ := runtime.Caller(0)
	fmt.Printf("Ptr: %d File: %s, Line: %d", caller, file, line)
}
