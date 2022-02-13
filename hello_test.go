package main

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestHelloWorld(t *testing.T) {
	result := emoji.Sprint("Hello :world_map:!")
	msg := HelloMessage()
	if result != msg {
		t.Fatalf(`HelloWorld() = %v, not match for %v`, msg, result)
	}
}
