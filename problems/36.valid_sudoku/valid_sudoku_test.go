package valid_sudoku

import "testing"

type Test struct {
	input    [][]byte
	expected bool
}

func TestValidSudoku(t *testing.T) {
	var tests = []Test{
		{board3, false},
		{board1, true},
		{board2, false},
	}

	for _, test := range tests {
		if output := isValidSudoku(test.input); output != test.expected {
			t.Error("TestValidSudoku Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
