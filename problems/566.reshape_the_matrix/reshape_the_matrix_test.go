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
		if output := matrixReshape(test.input.mat, test.input.r, test.input.c); deepEqual(output, test.expected) == false {
			t.Error("TestMatrixReshape Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}

func deepEqual(mat1 [][]int, mat2 [][]int) bool {
	if len(mat1) != len(mat2) {
		return false
	}
	for i := range mat1 {
		if utils.Equal(mat1[i], mat2[i]) == false {
			return false
		}
	}
	return true
}
