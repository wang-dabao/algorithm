package main

import "fmt"

/**
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]
示例 2：
输入：matrix =[[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof
 */

func spiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}
	x,y := len(matrix[0]),len(matrix)
	i,j := 0,0
	for x > 0 && y > 0 {
		index := 0
		for o := 0; o < x; o++ {
			res = append(res,matrix[i][j+o])
			index = j+o
		}
		j = index ; i++; y--
		if y == 0 {
			continue
		}
		for o := 0; o < y; o++ {
			res = append(res,matrix[i+o][j])
			index = i+o
		}
		i = index; j--; x--
		if x == 0 {
			continue
		}
		for o := 0; o < x; o++ {
			res = append(res,matrix[i][j-o])
			index = j-o
		}
		j = index ; i--; y--
		if y == 0 {
			continue
		}
		for o := 0; o < y; o++ {
			res = append(res,matrix[i-o][j])
			index = i-o
		}
		i = index ; j++; x--
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
}