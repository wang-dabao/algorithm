package main

import "fmt"

/**
给你一个非负整数数组nums ，你最初位于数组的第一个位置。
数组中的每个元素代表你在该位置可以跳跃的最大长度。
你的目标是使用最少的跳跃次数到达数组的最后一个位置。
假设你总是可以到达数组的最后一个位置。

示例 1:
输入: nums = [2,3,1,1,4]
输出: 2
解释: 跳到最后一个位置的最小跳跃数是 2。
从下标为 0 跳到下标为 1 的位置，跳1步，然后跳3步到达数组的最后一个位置。
示例 2:
输入: nums = [2,3,0,1,4]
输出: 2
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jump-game-ii
 */

/**
思路：典型的贪心算法 ：局部的最优解，每一步都得出最优的结果， 遍历整个数组，每一次都得出最优的结果，
 */

func jump(nums []int) int {
	res,index,end := 0,0,0
	for i := 0; i < len(nums)-1; i++ {
		if end < i+nums[i] {
			end = i+nums[i]
		}
		if i == index {
			index = end
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(jump([]int{1,2}))
}