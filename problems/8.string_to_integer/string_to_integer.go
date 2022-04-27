package string_to_integer

import (
	"math"
)

func myAtoi(s string) int {
	isNegative, str := sanitize(s)

	return toNumber(str, isNegative)
}

func sanitize(s string) (negative bool, result []byte) {
	first := true
	for i := 0; i < len(s); i++ {
		cur := s[i]
		if first {
			if cur == '-' {
				negative = true
			} else if cur == '+' {
				negative = false
			} else if cur == ' ' {
				continue
			} else if cur >= '0' && cur <= '9' {
				result = append(result, cur)
			} else {
				return false, []byte{'0'}
			}
			first = false
			continue
		}

		if cur >= '0' && cur <= '9' {
			result = append(result, cur)
		} else {
			break
		}
	}
	return
}

func toNumber(bs []byte, isNegative bool) (num int) {
	maxInt32 := math.MaxInt32
	if isNegative {
		maxInt32++
	}
	for i := 0; i < len(bs); i++ {
		if num = 10*num + int(bs[i]-'0'); num > maxInt32 {
			if isNegative {
				return math.MinInt32
			}
			return math.MaxInt32
		}
	}
	if isNegative {
		num = -num
	}
	return
}
