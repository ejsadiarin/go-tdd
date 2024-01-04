package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("e", 5)
	expected := "eeeee"

	if repeated != expected {
		t.Errorf("expected %s, but got %s", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("e", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("e", 5)
	fmt.Printf("%s", result)
	// Output: eeeee
}
