package merge_sorted_array

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

func TestMerge(t *testing.T) {
	var tests = []struct {
		input    Input
		expected []int
	}{
		{Input{[]int{2, 0}, 1, []int{1}, 1}, []int{1, 2}},
		{Input{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
		{Input{[]int{1}, 1, []int{}, 0}, []int{1}},
		{Input{[]int{}, 0, []int{1}, 1}, []int{1}},
	}
	for _, test := range tests {
		if output := merge(test.input.nums1, test.input.m, test.input.nums2, test.input.n); utils.Equal(output, test.expected) == false {
			t.Error("TestMerge Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
