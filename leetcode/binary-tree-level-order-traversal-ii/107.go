/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ans [][]int
	dfs(root, 0, &ans)

	// Reverse the traverse results
	for i, n := 0, len(ans); i < n / 2; i++ {
		ans[i], ans[n - 1 - i] = ans[n - 1 - i], ans[i]
	}
	
	return ans
}

func dfs(node *TreeNode, level int, ans *[][]int) {
	if len(*ans) <= level {
		*ans = append(*ans, []int{})
	}

	(*ans)[level] = append((*ans)[level], node.Val)

	if node.Left != nil {
		dfs(node.Left, level+1, ans)
	}

	if node.Right != nil {
		dfs(node.Right, level+1, ans)
	}
}
// @lc code=end

