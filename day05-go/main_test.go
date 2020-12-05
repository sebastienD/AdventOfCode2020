package main

import (
	"testing"
)

func TestShouldComputeRow(t *testing.T) {
	var testcase = []struct {
		data     string
		expected int
	}{
		{"FBFBBFF", 44},
		{"BFFFBBF", 70},
		{"FFFBBBF", 14},
		{"BBFFBBF", 102},
	}
	for _, tc := range testcase {
		actual, _ := compute(tc.data, 'F', 0, 127)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}

func TestShouldComputeCol(t *testing.T) {
	var testcase = []struct {
		data     string
		expected int
	}{
		{"RLR", 5},
		{"RRR", 7},
		{"RLL", 4},
	}
	for _, tc := range testcase {
		_, actual := compute(tc.data, 'L', 0, 7)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}
