package main

import (
	"testing"
)

func assertSum(t *testing.T, want int, got int, given []int) {
	if got != want {
		t.Errorf("got %d but want %d given %v", got, want, given)
	}
}

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assertSum(t, want, got, numbers)
	})
}
