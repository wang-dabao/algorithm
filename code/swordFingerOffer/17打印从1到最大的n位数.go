package main

/**
输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999
示例 1:
输入: n = 1
输出: [1,2,3,4,5,6,7,8,9]
 */


func printNumbers(n int) []int {
	//定义数组的边界
	max := 1
	for ; n>0; n-- {
		max *= 10
	}
	var res []int
	for i := 1; i < max; i++ {
		res = append(res,i)
	}
	return res
}