package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	//var slice = make([]int, 1, 1)
	checkSums := func(t *testing.T, sum, want []int) {
		t.Helper()
		if !reflect.DeepEqual(sum, want) {
			t.Errorf("got %v want %v", sum, want)
		}
	}

	t.Run("With variable slices", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		numb2 := []int{3, 4}
		numb3 := []int{5, 6}
		sum := SumAll(numbers, numb2, numb3)
		want := []int{15, 7, 11}

		checkSums(t, sum, want)
	})

	t.Run("with single slice", func(t *testing.T) {
		sum := SumAll([]int{1, 1})
		want := []int{2}

		checkSums(t, sum, want)

	})

	t.Run("with empty slice", func(t *testing.T) {
		sum := SumAll([]int{})
		want := []int{0}

		checkSums(t, sum, want)

	})
}
