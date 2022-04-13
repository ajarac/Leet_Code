package reverse_integer

import "testing"

type Test struct {
	input    int
	expected int
}

func TestReverseInteger(t *testing.T) {
	var tests = []Test{
		{1534236469, 0},
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for _, test := range tests {
		if output := reverse(test.input); output != test.expected {
			t.Error("TestReverseInteger Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
