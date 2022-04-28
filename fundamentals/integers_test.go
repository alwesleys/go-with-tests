package fundamentals

import (
	"fmt"
	"testing"
)

func TestModulo(t *testing.T) {
	got := Modulo(5, 3)
	want := 2

	if got != want {
		t.Errorf("got '%d' but want '%d'", got, want)
	}
}

func ExampleModulo() {
	r := Modulo(35, 6)
	fmt.Println(r)
	// Output: 5
}

func Modulo(a, b int) int {
	return a % b
}
