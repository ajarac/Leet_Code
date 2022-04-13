package intersection_of_two_arrays_ii

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums1 []int
	nums2 []int
}

type Test struct {
	input    Input
	expected [][]int
}

func TestIntersect(t *testing.T) {
	var tests = []Test{
		{Input{[]int{1, 2, 2, 1}, []int{2, 2}}, [][]int{{2, 2}}},
		{Input{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}}, [][]int{{4, 9}, {9, 4}}},
	}

	for _, test := range tests {
		output := intersect(test.input.nums1, test.input.nums2)
		anyEqual := false
		for _, ints := range test.expected {
			if utils.Equal(output, ints) {
				anyEqual = true
			}
		}
		if anyEqual == false {
			t.Error("TestIntersect Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}
