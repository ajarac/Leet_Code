package first_bad_version

import "testing"

func TestFirstBadVersion(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{5, 4},
		{1, 1},
		{2, 1},
	}
	for _, test := range tests {
		badVersion := &BadVersion{version: test.expected}
		if output := firstBadVersion(test.input, badVersion); output != test.expected {
			t.Error("TestFirstBadVersion Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
