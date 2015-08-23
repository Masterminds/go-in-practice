package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
)

func main() {
	listen()
}

func listen() {
	listener, err := net.Listen("tcp", ":1026")
	if err != nil {
		fmt.Println("Failed to open port on 1026")
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection")
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Fatal error: %s", err)
		}
		conn.Close()
	}()
	reader := bufio.NewReader(conn)

	data, err := reader.ReadBytes('\n')
	if err != nil {
		fmt.Println("Failed to read from socket.")
	}

	response(data, conn)
}

func response(data []byte, conn net.Conn) {
	conn.Write(data)
	panic(errors.New("Pretend I'm a real error"))
}
