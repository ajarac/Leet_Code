package decode_string

import "testing"

type Test struct {
	input    string
	expected string
}

func TestDecodeString(t *testing.T) {
	var tests = []Test{
		{"3[a2[c]]", "accaccacc"},
		{"3[a]2[bc]", "aaabcbc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, test := range tests {
		if output := decodeString(test.input); output != test.expected {
			t.Errorf("TestDecodeString Failed: %s inputted, %s expected, output %s", test.input, test.expected, output)
		}
	}
}
