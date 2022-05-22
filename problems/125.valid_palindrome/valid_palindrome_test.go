package valid_palindrome

import "testing"

type Test struct {
	input    string
	expected bool
}

func TestIsPalindrome(t *testing.T) {
	var tests = []Test{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, test := range tests {
		if output := isPalindrome(test.input); output != test.expected {
			t.Error("TestIsPalindrome Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
