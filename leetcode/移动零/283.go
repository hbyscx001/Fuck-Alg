/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	for startZeroIndex := 0; startZeroIndex < len(nums) - 1; startZeroIndex++ {
		if nums[startZeroIndex] != 0 {
			continue
		}

		for i := startZeroIndex + 1; i < len(nums); i++ {
			if nums[i] == 0 {
				continue
			}

			nums[i], nums[startZeroIndex] = nums[startZeroIndex], nums[i]
			break
		} 
	}
}
// @lc code=end

