package reshape_the_matrix

import (
	"leetcode/utils"
	"testing"
)

type Input struct {
	mat [][]int
	r   int
	c   int
}

type Test struct {
	input    Input
	expected [][]int
}

func TestMatrixReshape(t *testing.T) {
	var tests = []Test{
		{Input{[][]int{{1, 2}, {3, 4}}, 1, 4}, [][]int{{1, 2, 3, 4}}},
		{Input{[][]int{{1, 2}, {3, 4}}, 2, 4}, [][]int{{1, 2}, {3, 4}}},
	}

	for _, test := range tests {
		if output := matrixReshape(test.input.mat, test.input.r, test.input.c); utils.DeepEqual(output, test.expected) == false {
			t.Error("TestMatrixReshape Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
