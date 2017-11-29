package attribute

import (
	"log"
	"testing"
)

func TestText(t *testing.T) {
	te := New("text")
	if te.Get() != "" {
		log.Fatalf("Expected default text value to be an empty string")
	}

	te.Set("text goes here")

	if te.Get() != "text goes here" {
		log.Fatalf("Expected set(text goes here) to update the value")
	}
}
