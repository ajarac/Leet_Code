package maximum_score_words_formed_by_letters

import "testing"

type Input struct {
	words   []string
	letters []byte
	score   []int
}

type Test struct {
	input    Input
	expected int
}

func TestMaxScoreWords(t *testing.T) {
	var tests = []Test{
		{Input{[]string{"dog", "cat", "dad", "good"}, []byte{'a', 'a', 'c', 'd', 'd', 'd', 'g', 'o', 'o'}, []int{1, 0, 9, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}, 23},
		{Input{[]string{"xxxz", "ax", "bx", "cx"}, []byte{'z', 'a', 'b', 'c', 'x', 'x', 'x'}, []int{4, 4, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 0, 10}}, 27},
		{Input{[]string{"leetcode"}, []byte{'l', 'e', 't', 'c', 'o', 'd'}, []int{0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}}, 0},
	}
	for _, test := range tests {
		if output := maxScoreWords(test.input.words, test.input.letters, test.input.score); output != test.expected {
			t.Errorf("TestMaxScoreWords Failed: %v inputted, %d expected, received: %d", test.input, test.expected, output)
		}
	}
}
