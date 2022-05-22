package subsets

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    []int
	expected [][]int
}

func TestSubsets(t *testing.T) {
	var tests = []Test{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}

	for _, test := range tests {
		if output := subsets(test.input); utils.DeepEqual(output, test.expected) {
			t.Error("TestSubsets Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
