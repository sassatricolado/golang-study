package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	verifySum := func(t *testing.T, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("expected '%d' but got '%d'", expected, result)
		}
	}

	t.Run("sum empty slice", func(t *testing.T) {
		result := SumAllOfRest([]int{}, []int{0, 9})
		expected := []int{0, 9}
		verifySum(t, result, expected)
	})

	t.Run("sum all slices", func(t *testing.T) {
		result := SumAllOfRest([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}
		verifySum(t, result, expected)
	})
}
