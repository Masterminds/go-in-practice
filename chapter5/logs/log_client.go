package main

import (
	"log"
	"net"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:1902")
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}
	defer conn.Close()

	f := log.Ldate | log.Lshortfile
	logger := log.New(conn, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Panicln("This is a panic.")
}
