package pascal_triangle

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    int
	expected [][]int
}

func TestPascalTriangle(t *testing.T) {
	var tests = []Test{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}
	for _, test := range tests {
		if output := generate(test.input); utils.DeepEqual(output, test.expected) == false {
			t.Error("TestPascalTriangle Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
