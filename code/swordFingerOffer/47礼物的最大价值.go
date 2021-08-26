package main

/**
在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
示例 1:
输入:
[
 [1,3,1],
 [1,5,1],
 [4,2,1]
]
输出: 12
解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/li-wu-de-zui-da-jie-zhi-lcof
 */

/**
思路：和内个机器人路径啥的都差不多，一个思路，动态规划，因为当前值，取决于上一步的值，而且是有起点和终点的
dp数组 dp[i][j] i,j都是长度 值存储最多价值
状态转移方程 dp[i][j] = max(dp[i-1][j],dp[i][j-1]) + grid[i-1][j-1]
初始值 dp[0][0] = 0 和边界
 */
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m,n := len(grid),len(grid[0])
	dp := make([][]int,m+1)
	for o := 0; o <= m; o++ {
		dp[o] = make([]int,n+1)
	}
	dp[0][0] = 0
	for i:=0; i<=m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
				continue
			}
			if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j] + grid[i-1][j-1]
			}else {
				dp[i][j] = dp[i][j-1] + grid[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}