package first_unique_character_in_a_string

import "testing"

type Test struct {
	input    string
	expected int
}

func TestFirstUniqChar(t *testing.T) {
	var tests = []Test{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, test := range tests {
		if output := firstUniqChar(test.input); output != test.expected {
			t.Error("TestFirstUniqChar Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
