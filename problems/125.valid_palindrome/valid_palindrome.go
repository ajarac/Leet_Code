package valid_palindrome

import "unicode"

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isAlphaNumeric(s[left]) {
			left++
		}
		for left < right && !isAlphaNumeric(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func toLower(b byte) byte {
	return byte(unicode.ToUpper(rune(b)))
}

func isAlphaNumeric(b byte) bool {
	return isBetween(b, byte(48), byte(57)) || isBetween(b, byte(65), byte(90)) || isBetween(b, byte(97), byte(122))
}

func isBetween(d byte, a byte, b byte) bool {
	return a <= d && d <= b
}
