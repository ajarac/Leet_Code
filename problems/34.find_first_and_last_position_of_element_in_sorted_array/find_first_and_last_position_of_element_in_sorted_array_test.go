package find_first_and_last_position_of_element_in_sorted_array

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

type Test struct {
	input    Input
	expected []int
}

func TestSearchRange(t *testing.T) {
	var tests = []Test{
		{Input{[]int{1, 3}, 1}, []int{0, 0}},
		{Input{[]int{2, 2}, 2}, []int{0, 1}},
		{Input{[]int{5, 7, 7, 8, 8, 10}, 8}, []int{3, 4}},
		{Input{[]int{5, 7, 7, 8, 8, 10}, 6}, []int{-1, -1}},
		{Input{[]int{}, 0}, []int{-1, -1}},
	}

	for _, test := range tests {
		if output := searchRange(test.input.nums, test.input.target); utils.Equal(test.expected, output) == false {
			t.Error("TestSearchRange Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
