package leetcode

import "math"

/*
 * @lc app=leetcode id=887 lang=golang
 *
 * [887] Super Egg Drop
 */

// @lc code=start


func min(a, b int) int {
if a > b {
return b
}
return a
}

func max(a, b int) int {
if a > b {
return a
}
return b
}

func superEggDrop(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	matrix := make([][]int, K + 1)
	for i := 0; i <= K; i++ {
		matrix[i] = make([]int, N + 1)

		if i > 0  {
		matrix[i][1] = 1
		}
	}

	// calc matrix of 1 egg
	for i := 1; i <= N; i++ {
		matrix[1][i] = i
	}

	// 归纳问题的两个维度为，鸡蛋的个数与可以选择的楼层的个数:dp(k, n)
	// 对于一次测试来说，如果测试楼层是在x层，则有两种情况
	// case1：碎了，说明答案是dp(k-1, n-1) + 1
	// case2: 没碎，说明答案是dp(k, n - x) + 1
	// 这两种情况都有可能，我们取较大的值作为这次实验的上限，也就是说根据这次实验的结果，我们至多需要max(dp(k - 1, n - 1) + 1, dp(k, n - x) + 1)次操作
	// 当然不同的测试楼层会对结果有影响，因此，我们将遍历所有的x结果，取最小的作为最优路径
	// 先从底楼层往高楼层进行结果遍历

	for egg := 2; egg <= K; egg++ {
		for layer := 2; layer <= N; layer++ {
			for testLayer := 1; testLayer <= layer; testLayer++ {
				if matrix[egg][layer] == 0 {
					matrix[egg][layer] = layer
				}
				matrix[egg][layer] = min(matrix[egg][layer], max(matrix[egg][layer-testLayer], matrix[egg-1][testLayer-1])+1)
			}
		}
	}

	return matrix[K][N]
}


func superEggDrop2(K int, N int) int {
	if K == 0 || N == 0 {
		return 0
	}
	if K == 1 {
		return N
	}
	if N == 1 {
		return 1
	}

	dp := make([][]int, K+1)
	for i := 0; i < K+1; i++ {
		dp[i] = make([]int, N+1)
	}
	for i := 0; i <= N; i++ {
		dp[1][i] = i
	}
	for i := 1; i <= K; i++ {
		dp[i][1] = 1
	}
	for i := 2; i <= K; i++ {
		for j := 2; j <= N; j++ {
			for x := 1; x <= j; x++ {
				if dp[i][j] == 0 {
					dp[i][j] = math.MaxInt32
				}
				dp[i][j] = min(dp[i][j], max(dp[i][j-x], dp[i-1][x-1])+1)
			}
		}
	}
	return dp[K][N]
}
// @lc code=end

