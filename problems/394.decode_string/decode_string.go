package decode_string

import (
	"strconv"
	"strings"
	"unicode"
)

func decodeString(s string) string {
	result, _ := parse(s, 1, 0)
	return result
}

func repeat(n int, str string) string {
	result := strings.Builder{}

	for i := 0; i < n; i++ {
		result.WriteString(str)
	}

	return result.String()
}

func parse(str string, toRepeat int, nested int) (string, string) {
	num := 0
	temp := strings.Builder{}
	i := 0
	result := strings.Builder{}
	for i < len(str) {
		if unicode.IsNumber(rune(str[i])) {
			n, _ := strconv.Atoi(string(str[i]))
			num = num*10 + n
		} else if str[i] == '[' {
			toTemp, s := parse(str[i+1:], num, nested+1)
			temp.WriteString(toTemp)
			str = s
			i = -1
			num = 0
		} else if str[i] == ']' {
			result.WriteString(repeat(toRepeat, temp.String()))
			temp.Reset()
			if nested > 0 {
				str = str[i+1:]
				break
			}
			num = 0
		} else {
			temp.WriteByte(str[i])
		}
		i++
	}
	result.WriteString(temp.String())
	return result.String(), str
}
