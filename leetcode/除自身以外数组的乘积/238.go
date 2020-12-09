/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	leftProduct := make([]int, len(nums))

	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	fmt.Println(leftProduct)

	rightProductTmp := 1
	for j := len(nums) - 2; j>= 0; j-- {
		rightProductTmp = rightProductTmp * nums[j+1]
		leftProduct[j] = rightProductTmp * leftProduct[j]
	}

	fmt.Println(leftProduct)

	return leftProduct
}
// @lc code=end

