package zigzag_conversion

import "testing"

type Input struct {
	str  string
	rows int
}

func TestZigZagConversion(t *testing.T) {
	var tests = []struct {
		input    Input
		expected string
	}{
		{Input{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{Input{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{Input{"A", 1}, "A"},
	}

	for _, test := range tests {
		if output := convert(test.input.str, test.input.rows); output != test.expected {
			t.Error("TestZigZagConversion Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)

		}
	}
}
