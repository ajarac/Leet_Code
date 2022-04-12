package two_sums

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	var tests = []struct {
		input    Input
		expected []int
	}{
		{Input{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{Input{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{Input{[]int{3, 3}, 6}, []int{0, 1}},
	}

	for _, test := range tests {
		if output := twoSum(test.input.nums, test.input.target); utils.Equal(output, test.expected) == false {
			t.Error("TestTwoSum Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
