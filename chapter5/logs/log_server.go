package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

// Simple log server
func main() {

	listener, err := net.Listen("tcp", ":1902")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Failed to accept: %s", err)
		}
		printmsg(conn)
	}
}

func printmsg(conn net.Conn) {
	var err error
	var msg string
	for {
		msg, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading message: %s\n", err)
			}
			return
		}
		fmt.Println(msg)
	}

}
