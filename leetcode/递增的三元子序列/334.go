package leetcode

/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
import "math"

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	small := math.MaxInt32
	mid := math.MaxInt32

	for _, v := range nums {
		if v <= small {
			small = v
		} else if v <= mid {
			mid = v
		} else if v > mid {
			return true
		}
	}

	return false
}
// @lc code=end

