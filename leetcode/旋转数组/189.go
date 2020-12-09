package leetcode

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	visted := make([]bool, len(nums))

	for i, v := range nums {
		tmp := v
		newIndex := (k + i) % len(nums)
		for !visted[newIndex] {
			nums[newIndex], tmp = tmp, nums[newIndex]
			visted[newIndex] = true
			newIndex = (k + newIndex) % len(nums)
		}
	}
}
// @lc code=end

