package main

import (
	"log/syslog"
)

func main() {
	logger, err := syslog.New(syslog.LOG_LOCAL3, "narwhal")
	if err != nil {
		panic("Cannot attach to syslog")
	}

	logger.Debug("Debug message.")
	logger.Notice("Notice message.")
	logger.Warning("Warning message.")
	logger.Alert("Alert message.")
}
