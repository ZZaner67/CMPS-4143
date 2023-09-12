package mascot_test

import (
	"testing"

	mascot "example.com/go-demo-1/Mascot"
)

func TestMascot(t *testing.T) {
	if mascot.Bestmascot() != "Go Gopher" {
		t.Fatal("Wrong mascot :(")
	}
}
