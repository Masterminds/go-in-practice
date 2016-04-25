package main

import (
	"log"
	"os"
)

func main() {
	logfile, _ := os.Create("./log.txt")
	defer logfile.Close()

	f := log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Llongfile
	logger := log.New(logfile, "example ", f)

	//logger := log.New(logfile, "example ", log.LstdFlags|log.Lshortfile)

	logger.Println("This is a regular message.")
	logger.Fatalln("This is a fatal error.")
	logger.Println("This is the end of the function.")
}
