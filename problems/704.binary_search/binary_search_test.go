package binary_search

import (
	"testing"
)

type input struct {
	nums   []int
	target int
}

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		input    input
		expected int
	}{
		{input{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{input{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{input{[]int{-1, 0, 3, 5, 9, 12}, 13}, -1},
	}
	for _, test := range tests {
		if output := search(test.input.nums, test.input.target); output != test.expected {
			t.Error("TestBinarySearch Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
