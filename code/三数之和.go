package main

import (
	"fmt"
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组。
// 示例 1：
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
// 示例 2：
//输入：nums = []
//输出：[]
// 示例 3：
//输入：nums = [0]
//输出：[]
// 提示：
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
// Related Topics 数组 双指针 排序

func threeSum(nums []int) [][]int {
	//思路：排序+双指针。 先将数组排序，之后以开头作为起始a值，b、c分别是a后面的头尾节点，如果满足b+c == -a 说明他们相加=0，满足条件 这里由于排好序了，只要是a>0 那b+c就不可能是负数，直接跳过就行
	//如果b+c > -a 移动c的坐标，往左小的数移动  反之 b+c < -a 移动b的左边往右大的数
	result := make([][]int, 0)
	lenght := len(nums)
	sort.Ints(nums) //排序
	//起始循环，以nums[0] 作为 a
	for a := 0; a < lenght; a++ {
		//如果a往后移动的时候，当前值和移动前的值一样的话，那就跳过，因为过滤重复的
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		//c 指向最后端
		c := lenght - 1
		//b 指向a后面一位，也就是首段，这时b、c形成一个双指针，按照规则移动指针，循环为啥是b<c呢。因为bc重合后就没有继续下去的意义了
		for b := a + 1; b < c; b++ {
			if b > a + 1 && nums[b] == nums[b-1] {
				continue
			}
			for b < c && nums[b] + nums[c] > -1*nums[a] {
				c--
			}
			if b < c && nums[b] + nums[c] == -1*nums[a] {
				result = append(result, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return result
}

func main() {
	fmt.Print(threeSum([]int{-1,0,1,2,-1,-4}))
}
