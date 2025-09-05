package problems

import (
	"testing"
)

func TestValidateParenthesis(t *testing.T) {
	t.Run("valid parentheses", func(t *testing.T) {
		input := "()"
		expected := true
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})

	t.Run("valid nested parentheses", func(t *testing.T) {
		input := "((()))"
		expected := true
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})

	t.Run("valid mixed parentheses", func(t *testing.T) {
		input := "{[()]}"
		expected := true
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})

	t.Run("invalid parentheses - wrong order", func(t *testing.T) {
		input := "([)]"
		expected := false
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})

	t.Run("invalid parentheses - unbalanced", func(t *testing.T) {
		input := "((())"
		expected := false
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		input := ""
		expected := true
		result := ValidateParenthesis(input)
		if result != expected {
			t.Errorf("ValidateParenthesis(%q) = %v; want %v", input, result, expected)
		}
	})
}
