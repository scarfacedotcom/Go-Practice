package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, 1, 0},
	}
	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Addition failed for %d + %d: expected %d, got %d", tc.a, tc.b, tc.expected, result)
		}
	}
}
