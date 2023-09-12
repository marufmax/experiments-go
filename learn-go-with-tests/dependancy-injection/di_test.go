package dependancy_injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Maruf")

	got := buffer.String()
	want := "Hello, Maruf"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
