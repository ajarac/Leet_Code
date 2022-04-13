package move_zeroes

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    []int
	expected []int
}

func TestMoveZeroes(t *testing.T) {
	var tests = []Test{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{2, 1}, []int{2, 1}},
	}

	for _, test := range tests {
		moveZeroes(test.input)
		if utils.Equal(test.input, test.expected) == false {
			t.Error("TestMoveZeroes Failed", test.input, test.expected)
		}
	}
}
