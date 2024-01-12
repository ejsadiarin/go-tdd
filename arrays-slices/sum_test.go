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
	got := SumAll([]int{1, 2}, []int{0, 9})
	t.Run("two slices as args", func(t *testing.T) {
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checksum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checksum(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checksum(t, got, want)
	})
}
