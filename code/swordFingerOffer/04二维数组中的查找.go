package main

import "fmt"

/**
在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。
示例:
现有矩阵 matrix 如下：
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
给定 target=5，返回true。

给定target=20，返回false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof
 */

/**
思路： 暴力破解，双重for循环直接一个一个比 （面试这么写估计死定了.....）
还有一种方法 线性查找，我们已知这个数组都是升序的，也就是说我们从又上角的值开始比较，如果又上角大于目标值，就向左移动，小于目标值就向下移动
 */

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := len(matrix[0])-1, 0
	for i >= 0 && j <= len(matrix)-1 {
		if matrix[j][i] == target {
			return true
		}
		if matrix[j][i] > target {
			i--
			continue
		}
		if  matrix[j][i] < target {
			j++
			continue
		}
	}
	return false
}
func main() {
	fmt.Print(findNumberIn2DArray([][]int{{1, 1}},0))
}