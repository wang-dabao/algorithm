package main

/**
一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。
示例 1:
输入: [0,1,3]
输出: 2
示例2:
输入: [0,1,2,3,4,5,6,7,9]
输出: 8
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/que-shi-de-shu-zi-lcof
 */

/**
思路，在排序数组中查找数字的变形题，就是数字的值和数组的下标不一致就是有问题的,当然可以暴力破解，也可以二分法
 */

func missingNumber(nums []int) int {
	left,right := 0,len(nums)
	for left < right {
		index := (left+right) / 2
		if nums[index] == index {
			left = index + 1
		}else {
			right = index
		}
	}
	return left
}