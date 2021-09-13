package main

import "sort"

/**
47. 全排列 II
给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
示例 1：
输入：nums = [1,1,2]
输出：
[[1,1,2],
 [1,2,1],
 [2,1,1]]
示例 2：
输入：nums = [1,2,3]
输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 */

/**
思路：此题和46题是一样的，全排列——使用回溯法。并记录之前已经放进去的值
 */

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	var find func(num int,res []int,cache []bool)
	find = func(num int, res []int, cache []bool) {
		//这里的num代表是，组合排列数组确定下来几位 比如123  确定下来的是1 num=1 后面两位是随机的
		if num == len(nums) {
			result = append(result,append([]int{},res...))
			return
		}
		for i:=0; i<len(nums); i++ {
			if cache[i] || i > 0 &&  nums[i] == nums[i-1] && !cache[i-1]{
				continue
			}
			res = append(res,nums[i])
			cache[i] = true
			find(num+1,res,cache)
			cache[i] = false
			res = res[:len(res)-1]
		}
	}
	find(0,[]int{},make([]bool,len(nums)))
	return result
}
