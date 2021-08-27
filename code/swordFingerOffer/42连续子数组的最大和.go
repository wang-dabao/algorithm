package main

import "fmt"

/**
输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
要求时间复杂度为O(n)。

示例1:
输入: nums = [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释:连续子数组[4,-1,2,1] 的和最大，为6。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof
 */

/**
思路：动态规划.. 但是稍微变形， 用一个变量存储一个最大值..这样再数组中dp[i]存储着包含nums[i]的子数组最大值
dp数组 存储每段的最大值 dp[i] = max i是数组的长度 前提是dp[i] 一定是包含nums[i]的最大值
状态转移方程： 如果dp[i-1]<=0  dp[i] = nums[i]
 */

func maxSubArray(nums []int) int {
	var dp []int
	dp = append(dp,nums[0])
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+dp[i-1] > nums[i] {
			dp = append(dp,nums[i]+dp[i-1])
		}else {
			dp = append(dp,nums[i])
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}