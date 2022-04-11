package maximum_subarray

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		input    []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for _, test := range tests {
		if output := maxSubArray(test.input); output != test.expected {
			t.Error("TestMaxSubArray Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}
