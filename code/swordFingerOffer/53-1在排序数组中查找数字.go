package main

/**
统计一个数字在排序数组中出现的次数。
示例 1:
输入: nums = [5,7,7,8,8,10], target = 8
输出: 2
示例2:
输入: nums = [5,7,7,8,8,10], target = 6
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof
 */


/**
思路：因为是已经排好序的数组，我们可以用二分法查找。
 */

func search(nums []int, target int) int {
	searchIndex := func(nums []int, target int) int {
		left,right := 0,len(nums)
		for left < right {
			index := (right+left) / 2
			if nums[index] >= target {
				right = index
			}else {
				left = index + 1
			}
		}
		return left
	}
	if len(nums) == 0 {
		return 0
	}
	lindex := searchIndex(nums, target)
	rindex := searchIndex(nums, target+1) - 1
	return rindex-lindex+1
}