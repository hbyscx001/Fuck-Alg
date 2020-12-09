package leetcode

/*
 * @lc app=leetcode.cn id=215 lang=golang
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	return quickSearch(nums, 0, len(nums) - 1, len(nums) - k)
}


func quickSearch(nums[]int, left int, right int, position int) int {
	mid := patition(nums, left, right)

	if mid == position {
		return nums[mid]
	} else if mid < position {
		return quickSearch(nums, mid+1, right, position)
	} else {
		return quickSearch(nums, left, mid-1, position)
	}

}


func patition(nums[]int, left int, right int) int {
	if len(nums) == 0 {
		return -1
	}

	if left < 0 || left > right {
		return -1
	}

	pivot := left

	for left < right {
		if nums[left] < nums[pivot] {
			left++
			continue
		}

		if nums[right] >= nums[pivot] {
			right--
			continue
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	if nums[left] > nums[pivot] {
		nums[left - 1], nums[pivot] = nums[pivot], nums[left - 1]
		return left - 1
	} else {
		nums[left], nums[pivot] = nums[pivot], nums[left]
		return left
	}
} 
// @lc code=end

