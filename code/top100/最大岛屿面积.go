package main

import "fmt"

//给定一个包含了一些 0 和 1 的非空二维数组 grid 。
// 一个 岛屿 是由一些相邻的 1 (代表土地) 构成的组合，这里的「相邻」要求两个 1 必须在水平或者竖直方向上相邻。你可以假设 grid 的四个边缘都被
//0（代表水）包围着。
// 找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为 0 。)
// 示例 1:
// [[0,0,1,0,0,0,0,1,0,0,0,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,1,1,0,1,0,0,0,0,0,0,0,0],
// [0,1,0,0,1,1,0,0,1,0,1,0,0],
// [0,1,0,0,1,1,0,0,1,1,1,0,0],
// [0,0,0,0,0,0,0,0,0,0,1,0,0],
// [0,0,0,0,0,0,0,1,1,1,0,0,0],
// [0,0,0,0,0,0,0,1,1,0,0,0,0]]
// 对于上面这个给定矩阵应返回 6。注意答案不应该是 11 ，因为岛屿只能包含水平或垂直的四个方向的 1 。
// 示例 2:
// [[0,0,0,0,0,0,0,0]]


func maxAreaOfIsland(grid [][]int) int {
	//思路：深度优先遍历 dfs 采用递归+回溯法 分别向上下左右走，并记录当前的最大面积
	maxAre := 0
	//从数组的任意一个点出发
	for i:= 0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			sum := search(i, j, grid)
			if maxAre < sum {
				maxAre = sum
			}
		}
	}
	return maxAre
}

func search (x,y int,grid [][]int) int {
	//如果上一步是在边界上，那向上下左右走，会过界
	if x < 0 || y < 0 || x == len(grid) || y == len(grid[0]) || grid[x][y] != 1 {
		return 0
	}
	grid[x][y] = 0
	return 1 + search(x-1,y,grid) + search(x+1,y,grid) + search(x,y-1,grid) + search(x,y+1,grid)
}

func main() {
	i := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}
	fmt.Println(maxAreaOfIsland(i))
}
