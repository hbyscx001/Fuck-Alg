package leetcode

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start

type Node struct {
	max int
	min int
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	totalSum := nums[0]
	dp := make([]Node, len(nums))
	dp[0] = Node{
		max: nums[0],
		min: nums[0],
	}

	for i := 1; i < len(nums); i++ {
		curMax := Max(nums[i], dp[i-1].max * nums[i], dp[i-1].min * nums[i])
		curMin := Min(nums[i], dp[i-1].max * nums[i], dp[i-1].min * nums[i])

		if curMax > totalSum {
			totalSum = curMax
		}

		dp[i] = Node{
			max: curMax,
			min: curMin,
		}
	}

	return totalSum
}

func Max(args ...int) int {
	if len(args) == 0 {
		return 0
	}

	max := args[0]

	for _, v := range args {
		if v > max {
			max = v
		}
	}

	return max
}

func Min(args ...int) int {
	if len(args) == 0 {
		return 0
	}

	min := args[0]

	for _, v := range args {
		if v < min {
			min = v
		}
	}

	return min
}
// @lc code=end

