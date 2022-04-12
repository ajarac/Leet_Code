package rotate_array

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums  []int
	steps int
}

type Test struct {
	input    Input
	expected []int
}

func TestRotateArray(t *testing.T) {
	var tests = []Test{
		{Input{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{Input{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
	}
	for _, test := range tests {
		if output := rotate(test.input.nums, test.input.steps); utils.Equal(output, test.expected) == false {
			t.Error("TestRotateArray Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
