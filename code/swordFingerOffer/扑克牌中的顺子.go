package main

/**
从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。
示例1:
输入: [1,2,3,4,5]
输出: True
示例2:
输入: [0,0,1,2,5]
输出: True

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bu-ke-pai-zhong-de-shun-zi-lcof
 */

/**
思路：遍历整个数组，用map记录遍过的数据，防止有重复的数据，记录遍历得到的最大值和最小值。两者之差小于数组长度，说明是顺子
 */

func isStraight(nums []int) bool {
	hmap := make(map[int]bool, len(nums))
	max,min := 0,14
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		if hmap[nums[i]] {
			return false
		}
		hmap[nums[i]] = true
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max - min < len(nums)
}