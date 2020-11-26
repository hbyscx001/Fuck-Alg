package leetcode

import "unicode/utf8"

/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start

func partition(s string) [][]string {
	res := &[][]string{}

	partitionRecursive(s, []string{}, res)

	return *res
}

func partitionRecursive(s string, path []string, res *[][]string) {
	if len(s) == 0 {
		tmp := make([]string, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
		return
	}

	for i := 1; i <= len(s); i++ {
		if isPalindrome(s[:i]) {
			path = append(path, s[:i])
			partitionRecursive(s[i:], path, res)
			path = path[:len(path) - 1]
		}
	}
}

func isPalindrome(s string) bool {
	runeList := []rune{}

	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		runeList = append(runeList, r)
		s = s[size:]
	}

	start := 0
	end := len(runeList) - 1

	for start <= end {
		if runeList[start] != runeList[end] {
			return false
		}
		start++
		end--
	}

	return true
}
// @lc code=end

