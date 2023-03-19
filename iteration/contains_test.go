package iteration

import (
	"fmt"
	"testing"
)

func assertCorrectContains(t *testing.T, want bool, got bool) {
	if got != want {
		t.Errorf("got '%t' but want '%t'", got, want)
	}
}

func TestContains(t *testing.T) {
	t.Run("return true when text includes character", func(t *testing.T) {
		got := Contains("test", "e")
		want := true
		assertCorrectContains(t, want, got)
	})

	t.Run("return false when text doen not include character", func(t *testing.T) {
		got := Contains("test", "x")
		want := false
		assertCorrectContains(t, want, got)
	})
}

func ExampleContains() {
	got := Contains("test", "t")
	fmt.Println(got)
	// Output: true
}

func BenchmarkContains(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Contains("test", "t")
	}
}
