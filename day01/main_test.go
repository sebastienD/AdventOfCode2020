package main

import (
	"testing"
)

func TestShouldComputeValueForFirstPart(t *testing.T) {
	var testcase = []struct {
		data     []int
		expected int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 514579},
	}
	for _, tc := range testcase {
		actual := firstPart(tc.data)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}

func TestShouldComputeValueForSecondPart(t *testing.T) {
	var testcase = []struct {
		data     []int
		expected int
	}{
		{[]int{1721, 979, 366, 299, 675, 1456}, 241861950},
	}
	for _, tc := range testcase {
		actual := secondPart(tc.data)
		if actual != tc.expected {
			t.Errorf("Expected %v, actual %v\n", tc.expected, actual)
		}
	}
}
