package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Walking a simple integer")
	var one MyInt = 1
	walk(one, 0)

	fmt.Println("Walking a simple struct")
	two := struct{ Name string }{"foo"}
	walk(two, 0)

	p := &Person{
		Name:    &Name{"Count", "Tyrone", "Rugen"},
		Address: &Address{"Humperdink Castle", "Florian"},
	}
	fmt.Println("Walking a struct with struct fields")
	walk(p, 0)
}

type MyInt int

type Person struct {
	Name    *Name
	Address *Address
}

type Name struct {
	Title, First, Last string
}

type Address struct {
	Street, Region string
}

func walk(u interface{}, depth int) {
	val := reflect.Indirect(reflect.ValueOf(u))
	t := val.Type()
	tabs := strings.Repeat("\t", depth+1)
	fmt.Printf("%sValue is type %q (%s)\n", tabs, t, val.Kind())
	if val.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldVal := reflect.Indirect(val.Field(i))

			tabs := strings.Repeat("\t", depth+2)
			fmt.Printf("%sField %q is type %q (%s)\n",
				tabs, field.Name, field.Type, fieldVal.Kind())

			if fieldVal.Kind() == reflect.Struct {
				walk(fieldVal.Interface(), depth+1)
			}
		}
	}
}
