package search_insert_position

import "testing"

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
		{input{[]int{1, 3, 5, 6}, 2}, 1},
	}
	for _, test := range tests {
		if output := searchInsert(test.input.nums, test.input.target); output != test.expected {
			t.Error("Binary Search Test Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
