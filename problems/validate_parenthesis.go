package problems

func ValidateParenthesis(s string) bool {
	var stack []rune
	parenthesisMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		// Check if character is a closing bracket and
		// stack is not empty and contains the corresponding bracket type
		if opening, exists := parenthesisMap[char]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != opening {
				return false
			}
			// Pop the opening bracket
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}

	return len(stack) == 0
}
