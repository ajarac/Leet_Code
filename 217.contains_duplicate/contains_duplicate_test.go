package contains_duplicate

import (
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	var tests = []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, test := range tests {
		if output := containsDuplicate(test.input); output != test.expected {
			t.Error("TestContainsDuplicate Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}

}