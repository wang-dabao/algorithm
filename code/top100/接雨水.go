package main

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
示例 2：
输入：height = [4,2,0,3,2,5]
输出：9
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/trapping-rain-water
 */

/**
思路：双指针；头指针和尾指针。因为每个位置能接雨水的大小，取决于他的左右两边的最大高度中最小的内一个，**能接雨水= 两边最小的最大高度-本身的高度**
头指针和尾指针开始指向两端，分别向中间走，leftmax,和rightmax记录两端的最大值 当左边最大值比右边最大值小的时候，头指针向后移动，反之尾指针向前移动

其他还有动态规划 和 单调栈的方法 空间复杂度都是o(n) 动态规划思想跟双指针差球不多
动态规划思路：两个dp数组维护leftmax和rightmax，下标位置一致 正向遍历数组 leftmax中取值是当前值和上一个值的最大值 反向遍历数组rightmax也是如此。 最后遍历原数组，取左右的最小值相加
 */

func trap(height []int) int {
	left,right,leftmax,rightmax := 0,len(height)-1,0,0
	res := 0
	for left < right {
		if height[left] > leftmax {
			leftmax = height[left]
		}
		if height[right] > rightmax {
			rightmax = height[right]
		}
		if leftmax < rightmax {
			res = res + (leftmax - height[left])
			left++
			continue
		}else {
			res = res + (rightmax - height[right])
			right--
			continue
		}
	}
	return res
}
