package main

import (
	"log"
	"strings"
	"testing"
	"testing/quick"
)

// Pad whitespace-pads a string to a given length.
//
// If the string is longer than that, it truncates.
func Pad(s string, max uint) string {
	log.Printf("Testing Len: %d, Str: %s\n", max, s)
	l := uint(len(s))
	if l > max {
		return s[:max]
	}
	s += strings.Repeat(" ", int(max-l))
	return s
}

func TestPad(t *testing.T) {
	if r := Pad("test", 6); len(r) != 6 {
		t.Errorf("Expected 6, got %d", len(r))
	}
}

func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}

	if err := quick.Check(fn, &quick.Config{MaxCount: 200}); err != nil {
		t.Error(err)
	}
}
