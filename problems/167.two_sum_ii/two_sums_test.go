package two_sums_2

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestTwoSum2(t *testing.T) {
	var tests = []struct {
		input    Input
		expected []int
	}{
		{Input{[]int{2, 7, 11, 15}, 9}, []int{1, 2}},
		{Input{[]int{2, 3, 4}, 6}, []int{1, 3}},
		{Input{[]int{-1, 0}, -1}, []int{1, 2}},
	}

	for _, test := range tests {
		if output := twoSum2(test.input.nums, test.input.target); utils.Equal(output, test.expected) == false {
			t.Error("TestTwoSum Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
