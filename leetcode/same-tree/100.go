package leetcode

/*
 * @lc app=leetcode id=100 lang=golang
 *
 * [100] Same Tree
 */

// @lc code=start

func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil || q == nil {
		if p == nil && q == nil {
			return true
		}
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
// @lc code=end

