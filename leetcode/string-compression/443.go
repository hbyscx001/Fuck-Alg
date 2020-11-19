package leetcode

import "fmt"

/*
 * @lc app=leetcode id=443 lang=golang
 *
 * [443] String Compression
 */

// @lc code=start

type BytePair struct {
	c byte
	count int
}

func compress(chars []byte) int {

	// Case1: empty chars
	if len(chars) <=1 {
		return len(chars)
	}

	// Case2: normal
	sentryIdx := 0
	sentryChar := chars[0]
	writeIdx := 0

	for i := 0; i <= len(chars); i++ {
		if i == len(chars) || chars[i] != sentryChar {
			chars[writeIdx] = sentryChar
			writeIdx++

			if i > sentryIdx {
				stringCount := fmt.Sprintf("%d", i - sentryIdx)
				for j := 0; j < len(stringCount); j++ {
					chars[writeIdx] = stringCount[j]
					writeIdx++
				}
			}
			sentryChar = chars[i]
			sentryIdx = i
		}
	}

	return writeIdx
}
// @lc code=end

