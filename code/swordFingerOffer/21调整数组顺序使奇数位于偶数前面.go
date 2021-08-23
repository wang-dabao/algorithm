package main

import "fmt"

/**
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。
示例：
输入：nums =[1,2,3,4]
输出：[1,3,2,4]
注：[3,1,2,4] 也是正确的答案之一。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof
 */

/**
思路：暴力破解，就是遍历整个数组，奇数的放一起，偶数的放一起，之后再合并两个数组
当然面试不能这么写，基本GG了就.. 还有思路就是用双指针..头尾指针的办法 一个指向头，一个指向尾，之后如果头指针对应的值是奇数，就向后移动，尾指针指的是偶数就向前移动，不然就交换
 */

func exchange(nums []int) []int {
	start,end := 0,len(nums)-1
	for start < end {
		if nums[end] % 2 == 0 {
			end--
			continue
		}
		if nums[start] % 2 == 1 {
			start++
			continue
		}
		if nums[start] % 2 == 0 && nums[end] % 2 == 1 {
			tmp := nums[end]
			nums[end] = nums[start]
			nums[start] = tmp
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{2,16,3,5,13,1,16,1,12,18,11,8,11,11,5,1}))
}