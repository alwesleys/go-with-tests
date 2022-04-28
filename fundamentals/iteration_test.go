package fundamentals

import (
	"testing"
)

func TestReverse(t *testing.T) {
	got := Reverse("Go Language")
	want := "egaugnaL oG"

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("Go Language")
	}
}

func Reverse(s string) string {
	var r []byte
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}

	return string(r)
}
