package top_k_frequent_elements

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	nums []int
	k    int
}

type Test struct {
	input    Input
	expected []int
}

func TestTopKFrequent(t *testing.T) {
	var tests = []Test{
		{Input{[]int{1, 1, 1, 2, 2, 3, 4}, 1}, []int{1}},
		{Input{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
		{Input{[]int{1}, 1}, []int{1}},
	}
	for _, test := range tests {
		if output := topKFrequent(test.input.nums, test.input.k); !utils.Equal(output, test.expected) {
			t.Error("TestTopKFrequent Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
