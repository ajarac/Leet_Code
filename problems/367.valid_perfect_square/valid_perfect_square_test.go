package valid_perfect_square

import "testing"

type Test struct {
	input    int
	expected bool
}

func TestValidPerfectSquare(t *testing.T) {
	var tests = []Test{
		{16, true},
		{14, false},
		{1, true},
		{4, true},
	}

	for _, test := range tests {
		if output := isPerfectSquare(test.input); output != test.expected {
			t.Error("TestValidPerfectSquare Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
