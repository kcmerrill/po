package attribute

import (
	"log"
	"testing"
)

func TestNumber(t *testing.T) {
	n := New("number")
	if n.Get() != 0 {
		log.Fatalf("Expected default number value to be zero")
	}

	n.Set("1000")

	if n.Get() != 1000 {
		log.Fatalf("Expected set(1000) to update the value")
	}
}
