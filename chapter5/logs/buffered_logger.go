package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	logger := New("localhost:1902", 30*time.Second)

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 20; i++ {

		go func(logger *TcpLogger, id int) {
			for j := 0; j < 100; j++ {
				logger.Printf("Client %d message %d.", id, j)
				time.Sleep(1 * time.Second)
			}
			wg.Done()
		}(logger, i)
	}

	wg.Wait()
}

type TcpLogger struct {
	*log.Logger
	Addr    string
	Timeout time.Duration
	Queue   *logQueue
}

func New(addr string, timeout time.Duration) *TcpLogger {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		panic("Failed to connect to localhost:1902")
	}

	buff := make(chan []byte, 10)
	queue := &logQueue{conn, buff}

	f := log.Ldate | log.Lshortfile
	logger := log.New(queue, "example ", f)

	l := &TcpLogger{
		logger,
		addr,
		timeout,
		queue,
	}

	go DequeueLogs(l)

	return l
}

func (t *TcpLogger) Reconnect() error {
	limit := 10
	for i := 0; i < limit; i++ {
		conn, err := net.DialTimeout("tcp", t.Addr, t.Timeout)
		if err == nil {
			t.Queue.Destination = conn
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	msg := "Failed to reconnect after %d tries."
	return fmt.Errorf(msg, limit)
}

type logQueue struct {
	Destination io.Writer
	Messages    chan []byte
}

func (l *logQueue) Write(data []byte) (int, error) {

	// Write can never modify data, so we copy it.
	msg := make([]byte, len(data))
	copy(msg, data)

	l.Messages <- msg

	return len(data), nil
}

func DequeueLogs(logger *TcpLogger) {
	for msg := range logger.Queue.Messages {
		if _, e := logger.Queue.Destination.Write(msg); e != nil {
			fmt.Println("Attempting reconnect.")
			if e := logger.Reconnect(); e != nil {
				fmt.Printf("Failed reconnect. Dropping message. Msg: %s", msg)
			} else {
				fmt.Println("Reconnected. Resending messages.")
				logger.Queue.Destination.Write(msg)
			}
		}
	}
}
