package two_sums

import (
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
		if output := twoSum(test.input.nums, test.input.target); Equal(output, test.expected) == false {
			t.Error("TestTwoSum Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
