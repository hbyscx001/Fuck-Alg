package leetcode

/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
	im := m - 1
	in := n - 1

	for i := m + n - 1; i > im; i-- {
		if in < 0 {
			nums1[i] = nums1[im]
			im--
		} else if im < 0 {
			nums1[i] = nums2[in]
			in--
		} else if nums1[im] > nums2[in] {
			nums1[i] = nums1[im]
			im--
		} else {
			nums1[i] = nums2[in]
			in--
		}
	}
}
// @lc code=end

