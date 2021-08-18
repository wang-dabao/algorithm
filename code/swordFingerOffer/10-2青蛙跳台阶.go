package main
/**
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个n级的台阶总共有多少种跳法。

答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
示例 1：
输入：n = 2
输出：2
示例 2：
输入：n = 7
输出：21
示例 3：
输入：n = 0
输出：1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/qing-wa-tiao-tai-jie-wen-ti-lcof
 */

/**
思路：这是一个菲波那切数列的变种，也是入门级的动态规划，废话不多说，找出动态规划的三要素，这题就白捡....
dp数组： 一维数组 值是多少种跳法
状态转移方程：f(N) = f(N-1) + f(N-2)
初始值：f(0) = 1 f(1) = 1 f(2) = 2
 */

func numWays(n int) int {
	var dp []int
	dp = append(dp,1,1,2)
	for i := 3; i <= n; i++ {
		dp = append(dp,(dp[i-1] + dp[i-2]) % 1000000007)
	}
	return dp[n]
}