package main

import (
	"testing"
)

func TestShouldComputeValueForFirstPart(t *testing.T) {
	var testcase = []struct {
		data     []int
		loop     int
		expected int
	}{
		{[]int{0, 3, 6}, 10, 0},
		{[]int{1, 3, 2}, 2020, 1},
		{[]int{2, 1, 3}, 2020, 10},
		{[]int{1, 2, 3}, 2020, 27},
		{[]int{2, 3, 1}, 2020, 78},
		{[]int{3, 2, 1}, 2020, 438},
		{[]int{3, 1, 2}, 2020, 1836},
	}
	for _, tc := range testcase {
		actual := firstPart(tc.data, tc.loop)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}

func TestShouldComputeValueForSecondPart(t *testing.T) {
	var testcase = []struct {
		data     []int
		loop     int
		expected int
	}{
		{[]int{0, 3, 6}, 30000000, 175594},
		{[]int{1, 3, 2}, 30000000, 2578},
		{[]int{2, 1, 3}, 30000000, 3544142},
		{[]int{1, 2, 3}, 30000000, 261214},
		{[]int{2, 3, 1}, 30000000, 6895259},
		{[]int{3, 2, 1}, 30000000, 18},
		{[]int{3, 1, 2}, 30000000, 362},
	}
	for _, tc := range testcase {
		actual := firstPart(tc.data, tc.loop)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}
