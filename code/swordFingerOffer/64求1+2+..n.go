package main

/**
求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
示例 1：
输入: n = 3
输出:6
示例 2：
输入: n = 9
输出:45
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qiu-12n-lcof
 */

/**
思路：不让用这 不让用那....那就递归被
 */

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}