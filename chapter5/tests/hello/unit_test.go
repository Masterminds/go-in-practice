package hello

import "testing"

func TestHello(t *testing.T) {
	if Hello() != "hello" {
		t.Errorf("Expected 'hello', but got '%s'", Hello())
	}
}
