package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=140 lang=golang
 *
 * [140] 单词拆分 II
 */

// @lc code=start
func wordBreak22(s string, wordDict []string) []string {
	wordMap := map[string]bool{}
	for _, word := range wordDict {
		wordMap[word] = true
	}

	if len(s) == 0 {
		return []string{""}
	}

	// dp[i] 表示s[:i]的单词分割列表
	dp := make([][]string, len(s) + 1)
	dp[0] = []string{""}

	for i := 1; i <= len(s); i++ {
		for j :=0; j < i; j++ {
			splitWordCombine := dp[j]
			suffix := s[j:i]
			if splitWordCombine != nil && wordMap[suffix] {
				tmp := make([]string, len(splitWordCombine))
				for is := range tmp {
					if prefixStrings := splitWordCombine[is]; len(prefixStrings) > 0 {
						tmp[is] = prefixStrings + " " + suffix
					} else {
						tmp[is] = suffix
					}
				}
				dp[i] = append(dp[i], tmp...)
			}
		}
	}

	return dp[len(s)]
}

func wordBreak(s string, wordDict []string) []string {
	wordSet := map[string]bool{}
	for _, word := range wordDict {
		wordSet[word] = true
	}

	stringLength := len(s)
	dp := make([][][]string, stringLength + 1)
	
	var backtrack func(int) [][]string
	// backtrace是一个闭包函数，他用来遍历s[:i]的切片得到单词拆分的结果
	// 单词拆分的结果是一个二维切片，第一维是不同的分割结果
	backtrack = func(end int) [][]string {
		var wordList [][]string
		if wordList = dp[end]; len(wordList) == 0 {
			for start := end - 1; start >= 0; start-- {
				word := s[start:end]
				if _, has := wordSet[word]; has {
					if start == 0 {
						wordList = append(wordList, []string{word})
					} else {
						prevWordSubList := backtrack(start)
						for _, subList := range prevWordSubList {
							wordList = append(wordList, append(subList, word))
						}
					}
				}
			}
			dp[end] = wordList
		}
		return wordList
	}

	results := []string{}
	for _, splitWordList := range backtrack(stringLength) {
		results = append(results, strings.Join(splitWordList, " "))
	}

	return results
}


// @lc code=end

