package main

/**
实现pow(x,n)，即计算 x 的 n 次幂函数（即，xn）。不得使用库函数，同时不需要考虑大数问题。
示例 1：
输入：x = 2.00000, n = 10
输出：1024.00000
示例 2：
输入：x = 2.10000, n = 3
输出：9.26100
示例 3：
输入：x = 2.00000, n = -2
输出：0.25000
解释：2-2 = 1/22 = 1/4 = 0.25
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-zhi-de-zheng-shu-ci-fang-lcof
 */

/**
思路：怎么说呢，很抽象...我慢慢说。
首先我们分析 x^n 其实他可以转化成 x^n/2 * x^n/2 = x^(2*n/2) || x^(2*n/2 + 1)
也就是当n能整除2的时候，x^n 等于 (x^2)^n/2 ，那其实n如果是奇数的话，是不能整除2的，n是奇数就转化成了(x^2)^n/2 * x^1
那，我们将能计算的部分计算一下，比如 x^2 那这个方程就变成了 x^n = y^m || y^m * x^1 其中 y = x^2 m = n/2 就这样一点一点降解 当m=0的时候，y^m就等于1了..
这样如果我们最开始定义一个res=1 最开始x^n = x^n * res  每当n为奇数的时候，res = res * x^1 这样将剥离下的x收集起来，
*/

func myPow(x float64, n int) float64 {
	e := n
	if n < 0 {
		e = e * -1
		x = 1 / x
	}
	if n == 0 {
		return float64(1)
	}
	res := float64(1)
	for e > 0 {
		if e%2 == 1 {
			res *= x
		}
		x *= x
		e = e/2
	}
	return res
}