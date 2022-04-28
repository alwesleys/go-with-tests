package main

import (
	"fmt"
	"testing"
)

//* TEST
func TestHello(t *testing.T) {
	// returned value from the code
	got := Hello()
	// expected value
	want := "Hellow world"

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

//* code to test
func Hello() string {
	// domain
	return "Hello world"
}

func main() {
	// fmt.println = side effect
	fmt.Println(Hello())
}
