package main

import (
	"fmt"
	"strconv"
)

/**
给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，……，25 翻译成 “z”。一个数字可能有多个翻译。
请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。
示例 1:
输入: 12258
输出: 5
解释: 12258有5种不同的翻译，分别是"bccfi", "bwfi", "bczi", "mcfi"和"mzi"
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof
 */

/**
思路：动态规划....三要素
dp数组：dp[i] i是数字的位数，
状态转移方程： 我们分析，数字每加一位，就会多很多翻译的可能，假设加了一位，如果他能跟前一个数组成一个<26的数，那他就可能被翻译成一个字母，如果>25 那只能是新加的这一个自己单独翻译
1.<26: dp[i] = dp[i-2] + dp[i-1]
2.>=26: dp[i] = dp[i-1]
初始值 dp[0] = 0 dp[1] = 1
 */

func translateNum(num int) int {
	var dp []int
	dp = append(dp,1,1)
	nums := strconv.FormatInt(int64(num), 10)
	for i := 2; i <= len(nums); i++ {
		i1, _ := strconv.ParseInt(string(nums[i-1]), 10, 64)
		i2, _ := strconv.ParseInt(string(nums[i-1-1]), 10, 64)
		su := i1 + i2*10
		if su < 26 && i2 > 0{
			dp = append(dp,dp[i-2] + dp[i-1])
		}else {
			dp = append(dp,dp[i-1])
		}
	}
	return dp[len(nums)]
}

func main() {
	fmt.Println(translateNum(25))
}