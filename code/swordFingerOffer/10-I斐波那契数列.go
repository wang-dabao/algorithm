package main

/**
写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：

F(0) = 0, F(1)= 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
输入：n = 2
输出：1
示例 2：
输入：n = 5
输出：5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fei-bo-na-qi-shu-lie-lcof
 */

/**
思路：动态规划入门级算法。动态规划三要素：dp数组、状态转移方程、初始值
dp数组： 一维数组，存储第N项的值。int[N] = value
状态转移方程：F(N) = F(N - 1) + F(N - 2)
初始值：F(0) = 0, F(1)= 1
 */

func fib(n int) int {
	var dp []int
	dp = append(dp,0,1)
	for i := 2; i <= n; i++ {
		//答案需要取模 1e9+7（1000000007）
		dp = append(dp,(dp[i-1] + dp[i-2]) % 1000000007)
	}
	return dp[n]
}