package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		testCase string
		input    []int
		target   int
		expected int
	}{
		{"Test1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.testCase, func(t *testing.T) {
			search := binarySearch(tc.input, tc.target)
			if search != tc.expected {
				t.Errorf("want %d, got %d", tc.expected, search)
			}
			t.Logf("%d", search)
		})
	}
}
