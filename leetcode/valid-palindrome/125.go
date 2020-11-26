package leetcode

import (
	"unicode"
	"unicode/utf8"
)

/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start

func isLetterOrNumber(c byte) bool {
	crune, _ := utf8.DecodeRune([]byte{c})
	return unicode.IsLetter(crune) || unicode.IsNumber(crune)
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s)

	for start < end {
		if !isLetterOrNumber(s[start]) {
			start++
			continue
		}
		
		if !isLetterOrNumber(s[end]) {
			end--
			continue
		}

		if s[start] == s[end] {
			start++
			end--
		} else {
			return false
		}
	}

	return true
}
// @lc code=end

