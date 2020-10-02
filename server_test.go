package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var expected Phrase

	expected.Text = "Hello Medium!"
	result := HelloWorld()

	if expected.Text != result.Text {
		t.Errorf("Phrase was incorrect, got: %s, want: %s.", result.Text, expected.Text)
	}
}
