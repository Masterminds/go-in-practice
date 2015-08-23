package canary

import (
	"io"
	"testing"
)

type MyWriter struct{}

func (m *MyWriter) Write([]byte) error {
	return nil
}

func main() {
	m := map[string]interface{}{
		"w": &MyWriter{},
	}
}

func doSomething(m map[string]interface{}) {
	w := m["w"].(io.Writer)
}

func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}
