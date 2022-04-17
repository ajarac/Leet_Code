package ransom_note

import "testing"

type Input struct {
	ransomNote string
	magazine   string
}

type Test struct {
	input    Input
	expected bool
}

func TestCanConstruct(t *testing.T) {
	var tests = []Test{
		{Input{"fihjjjjei", "hjibagacbhadfaefdjaeaebgi"}, false},
		{Input{"aa", "bb"}, false},
		{Input{"a", "b"}, false},
		{Input{"aa", "aab"}, true},
	}

	for _, test := range tests {
		if output := canConstruct(test.input.ransomNote, test.input.magazine); output != test.expected {
			t.Error("TestCanConstruct Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
