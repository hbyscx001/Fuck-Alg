/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	mark := map[int]bool{}

	for _, i := range nums {
		if _, ok := mark[i]; ok {
			return true
		}
		mark[i] = true
	}

	return false
}
// @lc code=end

