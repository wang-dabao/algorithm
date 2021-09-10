package main

/**
输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。
示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[2,7] 或者 [7,2]
示例 2：
输入：nums = [10,26,30,31,47,60], target = 40
输出：[10,30] 或者 [30,10]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-wei-sde-liang-ge-shu-zi-lcof
 */

/**
思路：和"长度最小的子数组"是一样的，双指针滑动窗口，头尾指针，之后相加，因为是递增的，如果大于目标值，尾结点向前移动，如果小于头节点向后，知道两个相遇
 */

func twoSum(nums []int, target int) []int {
	head,tail := 0,len(nums)-1
	for head < tail {
		if nums[head] + nums[tail] == target {
			return []int{nums[head],nums[tail]}
		}
		if nums[head] + nums[tail] > target {
			tail--
		}
		if nums[head] + nums[tail] < target {
			head++
		}
	}
	return nil
}