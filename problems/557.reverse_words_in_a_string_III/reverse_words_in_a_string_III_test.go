package reverse_words_in_a_string_III

import "testing"

type Test struct {
	input    string
	expected string
}

func TestReverseWords(t *testing.T) {
	var tests = []Test{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
		{"God Ding", "doG gniD"},
	}
	for _, test := range tests {
		if output := reverseWords(test.input); output != test.expected {
			t.Error("TestReverseWords Failed: {} inputted, {} expected, output {}", test.input, test.expected, output)
		}
	}
}
