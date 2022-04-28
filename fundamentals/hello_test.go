package fundamentals

import (
	"testing"
)

//* TEST
func TestHello(t *testing.T) {
	// returned value from the code
	got := Hello()
	// expected value
	want := "Hello world"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

//* code to test
func Hello() string {
	// domain
	return "Hello world"
}
