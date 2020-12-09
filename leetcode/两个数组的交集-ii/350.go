/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	markMap := map[int]int{}
	result := []int{}

	for _, num := range nums1 {
		markMap[num] = markMap[num] + 1
	}

	for _, num := range nums2 {
		if markMap[num] > 0 {
			result = append(result, num)
			markMap[num] = markMap[num] - 1
		}
	}

	return result
}
// @lc code=end

