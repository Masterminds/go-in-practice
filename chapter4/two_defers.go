package main

func main() {
	defer func() { println("a") }()
	defer func() { println("b") }()
}
