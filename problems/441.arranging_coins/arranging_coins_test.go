package arranging_coins

import "testing"

type Test struct {
	input    int
	expected int
}

func TestArrangeCoins(t *testing.T) {
	var tests = []Test{
		{5, 2},
		{8, 3},
		{1, 1},
		{3, 2},
	}

	for _, test := range tests {
		if output := arrangeCoins(test.input); output != test.expected {
			t.Error("TestArrangeCoins Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
