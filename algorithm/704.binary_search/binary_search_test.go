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
		{input{[]int{1, 3, 5, 6}, 5}, 2},
		{input{[]int{1, 3, 5, 6}, 2}, 1},
		{input{[]int{1, 3, 5, 6}, 7}, 4},
	}
	for _, test := range tests {
		if output := search(test.input.nums, test.input.target); output != test.expected {
			t.Error("TestBinarySearch Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
