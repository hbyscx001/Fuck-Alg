/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}

	x := len(matrix)-1
	y := 0

	for x >= 0 && y < len(matrix[0]) {
		currentValue := matrix[x][y]
		if target == currentValue {
			return true
		} else if target < currentValue {
			x--
		} else if target > currentValue {
			y++
		} else {
			return false
		}
	}

	return false
}
// @lc code=end

