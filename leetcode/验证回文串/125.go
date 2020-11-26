package leetcode

/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func isLetterOrNumber(c byte) bool {
	crune, _ := utf8.DecodeRune([]byte{c})
	return unicode.IsLetter(crune) || unicode.IsNumber(crune)
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if !isLetterOrNumber(s[start]) {
			start++
			continue
		}
		
		if !isLetterOrNumber(s[end]) {
			end--
			continue
		}

		if strings.EqualFold(s[start:start+1], s[end:end+1]) {
			start++
			end--
		} else {
			fmt.Println(s[start: end])
			return false
		}
	}

	return true
}
// @lc code=end

