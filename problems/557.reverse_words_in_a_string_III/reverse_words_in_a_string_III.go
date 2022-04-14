package reverse_words_in_a_string_III

import "strings"

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	result := make([]string, len(split))
	for i, str := range split {
		result[i] = reverseWord(str)
	}
	return strings.Join(result, " ")
}

func reverseWord(str string) string {
	s := []byte(str)
	size := len(s)
	for i := 0; i < size/2; i++ {
		s[i], s[size-i-1] = s[size-i-1], s[i]
	}
	return string(s)
}
