package main

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	priority := syslog.LOG_LOCAL3 | syslog.LOG_NOTICE
	flags := log.Ldate | log.Lshortfile
	logger, err := syslog.NewLogger(priority, flags)
	if err != nil {
		fmt.Printf("Can't attach to syslog: %s", err)
		return
	}

	logger.Println("This is a test log message.")
}
