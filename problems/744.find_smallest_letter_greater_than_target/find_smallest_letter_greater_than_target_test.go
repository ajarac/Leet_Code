package find_smallest_letter_greater_than_target

import "testing"

type Input struct {
	letters []byte
	target  byte
}

type Test struct {
	input    Input
	expected byte
}

func TestNextGreatestLetter(t *testing.T) {
	var tests = []Test{
		{Input{[]byte("cfj"), byte('a')}, byte('c')},
		{Input{[]byte("cfj"), byte('c')}, byte('f')},
		{Input{[]byte("cfj"), byte('d')}, byte('f')},
		{Input{[]byte("cfj"), byte('j')}, byte('c')},
		{Input{[]byte("ab"), byte('z')}, byte('a')},
	}

	for _, test := range tests {
		if output := nextGreatestLetter(test.input.letters, test.input.target); output != test.expected {
			t.Error("TestNextGreatestLetter Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
