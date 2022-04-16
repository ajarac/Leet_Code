package kth_missing_positive_number

import "testing"

type Input struct {
	arr []int
	k   int
}

type Test struct {
	input    Input
	expected int
}

func TestFindKthPositive(t *testing.T) {
	var tests = []Test{
		{Input{[]int{2, 3, 4, 7, 11}, 5}, 9},
		{Input{[]int{1, 2, 3, 4}, 2}, 6},
	}

	for _, test := range tests {
		if output := findKthPositive(test.input.arr, test.input.k); output != test.expected {
			t.Error("TestFindKthPositive Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}
