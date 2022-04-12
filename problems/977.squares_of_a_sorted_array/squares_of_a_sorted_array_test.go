package squares_of_a_sorted_array

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    []int
	expected []int
}

func TestSortedSquares(t *testing.T) {
	var tests = []Test{
		{[]int{-1}, []int{1}},
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, test := range tests {
		if output := sortedSquares(test.input); utils.Equal(output, test.expected) == false {
			t.Error("TestSortedSquares Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
