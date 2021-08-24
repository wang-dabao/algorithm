package main

/**
请实现一个函数用来匹配包含'. '和'*'的正则表达式。模式中的字符'.'表示任意一个字符，而'*'表示它前面的字符可以出现任意次（含0次）。
在本题中，匹配是指字符串的所有字符匹配整个模式。例如，字符串"aaa"与模式"a.a"和"ab*ac*a"匹配，但与"aa.a"和"ab*a"均不匹配。

示例 1:
输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:
输入:
s = "aa"
p = "a*"
输出: true
示例3:
输入:
s = "ab"
p = ".*"
输出: true
解释:".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:
输入:
s = "aab"
p = "c*a*b"
输出: true
解释:因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:
输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

解释:因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zheng-ze-biao-da-shi-pi-pei-lcof
 */

/**
思路： 动态规划经典进阶题，正则表达式，我们先来找出动态规划的三要素：dp数组、状态转移方程、初始值
dp数组：dp[i][j] i为s字符串的长度 j为字符串p的长度  dp[i][j]的值是记录能否匹配
状态转移方程：
1.p[j]==s[i] 当p的最后一个字符和s的最后一个字符相等 dp[i][j] = dp[i-1][j-1]
2.p[j]=='.' 当p的最后一个字符是'.' 他可以匹配任意  dp[i][j] = dp[i-1][j-1]
3.p[j]=='*' p的最后一个字符时'*' 这个比较麻烦..因为'*'的组合 可以前面是字母也可以是'.' 前面如果是字母的时候
 3.1 p[j-1] == '.' || p[j-1] = s[i]
   3.1.1 dp[i][j] = dp[i-1][j]
   3.1.2 dp[i][j] = dp[i][j-1]
   3.1.3 dp[i][j] = dp[i][j-2]
 3.2 除了3.1的情况，那就只剩前面是字母但是p[j-1] != s[i]  dp[i][j] = dp[i][j-2]
 */

func isMatch(s string, p string) bool {
	m,n := len(s),len(p)
	dp := make([][]bool, m+1)
	for o := 0; o < len(dp); o++ {
		dp[o] = make([]bool,n+1)
	}
	dp[0][0] = true
	//初始化首行，这么写比较好理解一些
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	//i.j代表字符串的长度,
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			//p[j]==s[i]
			if p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			//p[j]=='.'
			if p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			//p[j]=='*'
			if p[j-1] == '*' {
				if p[j-1-1] == '.' || p[j-1-1] == s[i-1] {
					dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
				}else if p[j-1-1] != s[i-1] {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[m][n]
}
