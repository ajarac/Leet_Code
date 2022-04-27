package container_with_most_water

import "testing"

type Test struct {
	input    []int
	expected int
}

func TestMaxArea(t *testing.T) {
	var tests = []Test{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}

	for _, test := range tests {
		if output := maxArea(test.input); output != test.expected {
			t.Error("TestMaxArea Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
