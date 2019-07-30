package main

import "testing"

func TestSumAll(t *testing.T) {
	t.Run("With Array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		sum := Sum(numbers)
		want := 15

		if sum != want {
			t.Errorf("got %d want %d", sum, want)
		}
	})

}
