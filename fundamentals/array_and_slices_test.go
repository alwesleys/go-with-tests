package fundamentals

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{5, 7, 2, 4, 2}
		got := Average(numbers)
		want := 4

		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
		}
	})

	t.Run("Collection of 3 numbers", func(t *testing.T) {
		numbers := []int{5, 7, 3}
		got := Average(numbers)
		want := 5

		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
		}
	})
}

func Average(numbers []int) int {
	var s int
	for _, n := range numbers {
		s += n
	}

	return s / len(numbers)
}

func TestAverageAll(t *testing.T) {
	got := AverageAll([]int{2, 6}, []int{0, 998})
	want := []int{4, 499}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func AverageAll(numbers ...[]int) []int {

	var allAve []int

	for _, n := range numbers {
		allAve = append(allAve, Average(n))
	}

	return allAve
}
