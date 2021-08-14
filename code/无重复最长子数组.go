package main

/**
给定一个数组arr，返回arr的最长无重复元素子数组的长度，无重复指的是所有数字都不相同。
子数组是连续的，比如[1,3,5,7,9]的子数组有[1,3]，[3,5,7]等等，但是[1,3,7]不是子数组
 示例1
输入：
[2,3,4,5]
复制
返回值：
4
复制
说明：
[2,3,4,5]是最长子数组
*/

func maxLength( arr []int ) int {
	// 思路 ：跟无重复最长子串一样的   map + 双指针 map存储走过的值，
	cMap := make(map[int]int,len(arr))
	//定义两个指针，初始指向头
	start,end := 0,0
	maxLen := 0
	for end < len(arr) {
		if _,ok := cMap[arr[end]]; !ok {
			cMap[arr[end]] = 1
			end++
			if maxLen < end-start {
				maxLen = end-start
			}
		}else {
			if maxLen < end-start {
				maxLen = end-start
			}
			delete(cMap,arr[start])
			start++
		}
	}
	return maxLen
}