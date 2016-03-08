//go:generate ./queue MyInt
package main

import "fmt"

type MyInt int

func main() {
	var one, two, three MyInt = 1, 2, 3
	q := NewMyIntQueue()
	q.Insert(one)
	q.Insert(two)
	q.Insert(three)

	fmt.Printf("First value: %d\n", q.Remove())
}
