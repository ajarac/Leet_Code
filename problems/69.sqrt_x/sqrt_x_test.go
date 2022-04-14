package sqrt_x_

import "testing"

type Test struct {
	input    int
	expected int
}

func TestSqrtX(t *testing.T) {
	var tests = []Test{{4, 2}, {8, 2}}

	for _, test := range tests {
		if output := mySqrt(test.input); output != test.expected {
			t.Error("TestSqrtX Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
