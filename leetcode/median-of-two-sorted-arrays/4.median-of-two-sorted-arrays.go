package leetcode

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var (
        nums1Length = len(nums1)
        nums2Length = len(nums2)
    )

    if nums1Length > nums2Length {
        return findMedianSortedArrays(nums2, nums1)
    }

    var (
        iMin = 0
        iMax = nums1Length
        i, j int
        halfLength = (nums1Length + nums2Length + 1) / 2
        maxLeft int
        minRight int
        isOdd = (nums1Length + nums2Length) % 2 == 1
    )

    for iMin <= iMax {
        i  = (iMin + iMax) / 2
        j = halfLength - i

        if i < iMax && nums2[j - 1] > nums1[i] {
            // i is small
            iMin = i + 1
        } else if i > iMin && nums1[i - 1] > nums2[j] {
            // i is big
            iMax = i - 1
        } else {
            // i is perfect
            // i == 0
            // i == nums1Length
            if i == 0 {
                maxLeft = nums2[j - 1]
            } else if j == 0 {
                maxLeft = nums1[i - 1]
            } else {
                maxLeft = max(nums1[i - 1], nums2[j - 1])
            }

            if isOdd {
                return float64(maxLeft)
            }

            if i == nums1Length {
                minRight = nums2[j]
            } else if j == nums2Length {
                minRight = nums1[i]
            } else {
                minRight = min(nums1[i], nums2[j])
            }

            return float64(minRight + maxLeft) / 2
        }
    }
    return 0.0
}

func min(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a int, b int) int {
    if a > b {
        return a
    }
    return b
}
// @lc code=end

