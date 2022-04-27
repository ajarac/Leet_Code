package string_to_integer

import "testing"

type Test struct {
	input    string
	expected int
}

func TestMyAtoi(t *testing.T) {
	var tests = []Test{
		{"+-12", 0},
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"3.14159", 3},
		{"+1", 1},
	}

	for _, test := range tests {
		if output := myAtoi(test.input); output != test.expected {
			t.Error("TestMyAtoi Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
