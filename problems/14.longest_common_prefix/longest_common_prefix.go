package longest_common_prefix

import (
	"leetcode/utils"
	"strings"
)

func longestCommonPrefix(strs []string) string {

	minLength := int(^uint(0) >> 1)
	for _, str := range strs {
		minLength = utils.Min(minLength, len(str))
	}

	result := strings.Builder{}
	for i := 0; i < minLength; i++ {
		letter := strs[0][i]
		for _, str := range strs {
			if letter != str[i] {
				return result.String()
			}
		}
		result.WriteByte(letter) //  = result + string(letter)
	}

	return result.String()
}
