package main

import "fmt"

/**
地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

示例 1：
输入：m = 2, n = 3, k = 1
输出：3
示例 2：
输入：m = 3, n = 1, k = 0
输出：1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof
 */

/**
思路：深度优先遍历dfs 也是基本的回溯法+剪枝  其中需要写一个判断横纵坐标个位数之和是否大于K的方法
最优解当然是广度优先遍历（bfs 核心是拿队列把下一次走的所有路径存储起来，之后消费队列中的元素，自己写一写吧...）
 */

func movingCount(m int, n int, k int) int {
	muns := make([][]int,m)
	for i, _ := range muns {
		muns[i] = make([]int,n)
	}
	check := func(i,j,k int) bool {
		checkSum := 0
		for i > 0 {
			checkSum += i%10
			i = i/10
		}
		for j > 0 {
			checkSum += j%10
			j = j/10
		}
		if checkSum > k {
			return true
		}
		return false
	}
	var dfs func(i,j,sum int,muns [][]int) int
	dfs = func(i, j, sum int,muns [][]int) int {
		if i<0 || j<0 || i>=m || j>=n || check(i,j,k) || muns[i][j] == 1{
			return 0
		}
		muns[i][j] = 1
		return sum + 1 + dfs(i-1,j,sum,muns) +
			dfs(i+1,j,sum,muns) +
			dfs(i,j-1,sum,muns) +
			dfs(i,j+1,sum,muns)
	}
	return dfs(0,0,0,muns)
}

func main() {
	fmt.Print(movingCount(3,2,17))
}
