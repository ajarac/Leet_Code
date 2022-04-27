package longest_common_prefix

import "testing"

type Test struct {
	input    []string
	expected string
}

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []Test{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}
	for _, test := range tests {
		if output := longestCommonPrefix(test.input); output != test.expected {
			t.Error("TestLongestCommonPrefix Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
