package valid_anagram

import "testing"

type Input struct {
	s string
	t string
}

type Test struct {
	input    Input
	expected bool
}

func TestCanConstruct(t *testing.T) {
	var tests = []Test{
		{Input{"anagram", "nagaram"}, true},
		{Input{"rat", "car"}, false},
	}

	for _, test := range tests {
		if output := isAnagram(test.input.s, test.input.t); output != test.expected {
			t.Error("TestCanConstruct Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
