package min_cost_climbing_stairs

import "testing"

type Test struct {
	input    []int
	expected int
}

func TestMinCostClimbingStairs(t *testing.T) {
	var tests = []Test{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}

	for _, test := range tests {
		if output := minCostClimbingStairs(test.input); output != test.expected {
			t.Error("TestMinCostClimbingStairs Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
