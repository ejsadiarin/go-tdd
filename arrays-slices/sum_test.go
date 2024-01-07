package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	slicesNumbers := []int{1, 1, 2, 3}

	got := Sum(slicesNumbers)
	want := 7

	if got != want {
		t.Errorf("want %d, but got %d, given %v", want, got, slicesNumbers)
	}
}

func TestSumAll(t *testing.T) {
	t.Run("two slices as args", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("two tail slices as args", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
