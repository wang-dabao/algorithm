package main

import "fmt"

/**
有序的整形数组二分查找
 */

func erfen(nums []int,tag int) int {
	left,right := 0,len(nums)-1
	for left < right {
		index := left + (right-left)/2
		if nums[index] == tag {
			return index
		}
		if nums[index] < tag {
			left = index + 1
		}else {
			right = index - 1
		}
	}
	return -1
}

func erfen2(nums []int,tag int) int {
	var find func(nums []int,left,right int) int
	find = func(nums []int, left, right int) int {
		if left > right {
			return -1
		}
		if left+1 == right && nums[left] == tag {
			return left
		}
		index := left + (right-left)/2
		if nums[index] == tag {
			return index
		}else if nums[index] < tag {
			return find(nums,index+1,right)
		}else {
			return find(nums,left,index+1)
		}
	}
	return find(nums,0,len(nums)-1)
}

func main() {
	fmt.Println(erfen2([]int{1,3,4,6,8,9,10,11},10))
}