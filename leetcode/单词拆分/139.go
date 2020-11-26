package leetcode

/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)

	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	// dp[i]: s[0:i] can be word break
	// dp[i] 为真的条件是，存在 0 <= j < i 使得dp[j] = true and s[j+1:i] in wordset
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// calc dp matrix for each i from 1 to len(s)
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && wordDictMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

// @lc code=end

// 分割两份，一份word一份递归
func wordBreak2(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)

	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	return wordBreak2Recv(s, wordDictMap)
}

func wordBreak2Recv(s string, wordMap map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	for splitIndex := 0; splitIndex < len(s); splitIndex++ {
		leftPart, rightPart := s[:splitIndex], s[splitIndex:]

		if wordBreak2Recv(leftPart, wordMap) && wordMap[rightPart] {
			return true
		}
	}

	return false
}

func wordBreak3(s string, wordDict []string) bool {
	wordDictMap := make(map[string]bool)
	wordBreakResultMap := make(map[string]bool)

	for _, word := range wordDict {
		wordDictMap[word] = true
	}

	return wordBreak2RecvMemo(s, wordDictMap, wordBreakResultMap)
}

func wordBreak2RecvMemo(s string, wordMap map[string]bool, wordBreakResultMap map[string]bool) bool {
	if len(s) == 0 {
		return true
	}

	for splitIndex := 0; splitIndex < len(s); splitIndex++ {
		leftPart, rightPart := s[:splitIndex], s[splitIndex:]

		if !wordBreakResultMap[leftPart] {
			wordBreakResultMap[leftPart] = wordBreak2Recv(leftPart, wordMap)
		}
		if wordBreakResultMap[leftPart] && wordMap[rightPart] {
			return true
		}
	}

	return false
}