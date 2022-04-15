package search_a_2D_matrix

import "testing"

type Input struct {
	matrix [][]int
	target int
}

type Test struct {
	input    Input
	expected bool
}

func TestSearchMatrix(t *testing.T) {
	var tests = []Test{
		{Input{[][]int{{1, 3}}, 1}, true},
		{Input{[][]int{{1}}, 1}, true},
		{Input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3}, true},
		{Input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13}, false},
	}

	for _, test := range tests {
		if output := searchMatrix(test.input.matrix, test.input.target); output != test.expected {
			t.Error("TestSearchMatrix Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
