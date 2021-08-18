package main

/**
给定一个m x n 二维字符网格board 和一个字符串单词word 。如果word 存在于网格中，返回 true ；否则，返回 false 。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）
示例 1：
输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
输出：true
示例 2：
输入：board = [["a","b"],["c","d"]], word = "abcd"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ju-zhen-zhong-de-lu-jing-lcof
 */

/**
思路：深度优先遍历dfs, 采用回溯法（递归）向其上下左右走，这里注意一个问题，不能重复利用字母，所以在向其他方位走的时候，需要将走过的路的值置为' '
 */

func exist(board [][]byte, word string) bool {
	var dfs func(board [][]byte,i,j,index int,word string) bool
	dfs = func(board [][]byte,i,j,index int,word string) bool {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) || board[i][j] != word[index] {
			return false
		}
		if len(word)-1 == index {
			return true
		}
		board[i][j] = ' '
		b := dfs(board, i+1, j, index+1, word) ||
			 dfs(board, i-1, j, index+1, word) ||
			 dfs(board, i, j+1, index+1, word) ||
			 dfs(board, i, j-1, index+1, word)
		board[i][j] = word[index]
		return b
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, i, j,0, word) {
				return true
			}
		}
	}
	return false
}