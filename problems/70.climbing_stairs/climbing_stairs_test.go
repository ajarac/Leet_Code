package climbing_stairs

import "testing"

type Test struct {
	input    int
	expected int
}

func TestClimbStairs(t *testing.T) {
	var tests = []Test{
		{5, 8},
		{2, 2},
		{3, 3},
		{4, 5},
	}

	for _, test := range tests {
		if output := climbStairs(test.input); output != test.expected {
			t.Error("TestClimbStairs Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
