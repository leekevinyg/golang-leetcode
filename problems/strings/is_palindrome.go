package problems

import (
	"strings"
	"unicode"
)

func IsPalindrome(s string) bool {
	stack := make([]rune, 0)
	for _, c := range strings.ToLower(s) {
		if isAlphaNumeric := unicode.IsLetter(c) || unicode.IsDigit(c); isAlphaNumeric {
			stack = append(stack, c)
		}
	}
	originalString := string(stack)

	var b strings.Builder
	b.Grow(len(stack) * 3)

	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		b.WriteRune(c)
	}

	return b.String() == originalString
}
