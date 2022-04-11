package guess_number_higher_or_lower

import "testing"

func TestGuessNumberHigherOrLower(t *testing.T) {
	var tests = []struct {
		input    int
		expected int
	}{
		{10, 6},
		{1, 1},
		{2, 1},
	}
	for _, test := range tests {
		badVersion := &GuessGame{pick: test.expected}
		if output := guessNumberHigherOrLower(test.input, badVersion); output != test.expected {
			t.Error("TestGuessNumberHigherOrLower Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
