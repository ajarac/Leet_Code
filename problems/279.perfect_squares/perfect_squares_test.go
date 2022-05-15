package perfect_squares

import "testing"

type Test struct {
	input    int
	expected int
}

func TestNumSquares(t *testing.T) {
	var tests = []Test{
		{12, 3},
		{13, 2},
	}

	for _, test := range tests {
		if output := numSquares(test.input); output != test.expected {
			t.Error("TestNumSquares Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
