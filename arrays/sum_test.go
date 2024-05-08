package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum all collections", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{1, 10})
		want := []int{3, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAll([]int{}, []int{})
		want := []int{0, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 1, 1})
		want := []int{5, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1})
		want := []int{0, 1}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
