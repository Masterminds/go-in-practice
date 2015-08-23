package main

import "fmt"

func main() {

	err := &ParseError{
		Message: "Unexpected char ';'",
		Line:    5,
		Char:    38,
	}

	fmt.Println(err.Error())
}

type ParseError struct {
	Message    string
	Line, Char int
}

func (p *ParseError) Error() string {
	format := "%s on Line %d, Char %d"
	return fmt.Sprintf(format, p.Message, p.Line, p.Char)
}
