package reverse_string

import (
	"leetcode/utils"
	"testing"
)

type Test struct {
	input    []byte
	expected []byte
}

func TestReverseString(t *testing.T) {
	var tests = []Test{
		{[]byte("hello"), []byte("olleh")},
		{[]byte("Hannah"), []byte("hannaH")},
	}

	for _, test := range tests {
		reverseString(test.input)
		if utils.Equal[byte](test.input, test.expected) == false {
			t.Error("ReverseString Failed: {} inputted, {} expected", test.input, test.expected)
		}

	}
}
