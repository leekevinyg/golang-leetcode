package problems

import "testing"

func TestIsPalindrome(t *testing.T) {
    tests := []struct {
        name string
        in   string
        want bool
    }{
        {"example 1 - phrase with punctuation", "A man, a plan, a canal: Panama", true},
        {"example 2 - not palindrome", "race a car", false},
        {"example 3 - single space -> empty", " ", true},

        // additional cases
        {"empty string", "", true},
        {"single char", "x", true},
        {"two different", "ab", false},
        {"two same", "aa", true},
        {"mixed case punctuation", "Madam, I'm Adam", true},
        {"digits and letters", "1a2 2a1", true},
        {"digits non-palindrome", "12321", true},
        {"digits non-palindrome2", "1231", false},
        {"leet example 0P", "0P", false},
        {"unicode palindrome", "あいいあ", true},
        {"unicode mixed", "あaあ", true},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := IsPalindrome(tc.in)
            if got != tc.want {
                t.Fatalf("IsPalindrome(%q) = %v, want %v", tc.in, got, tc.want)
            }
        })
    }
}