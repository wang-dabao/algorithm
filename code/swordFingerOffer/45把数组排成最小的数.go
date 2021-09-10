package main

import (
	"fmt"
	"sort"
	"strings"
)

/**
输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

示例 1:
输入: [10,2]
输出: "102"
示例2:
输入: [3,30,34,5,9]
输出: "3033459"
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof
 */

/**
思路：将数组排序，但是排序规则自定义....例如 "30" + "3" < "3" + "30" 所以"30" 在 "3"前面
写快排.基准数
 */

func minNumber(nums []int) string {
	compare := func (a,b int)bool{
		str1 := fmt.Sprintf("%d%d",a,b)
		str2 := fmt.Sprintf("%d%d",b,a)
		if str1<str2 {
		return true
		}
		return false
	}
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i],nums[j])
	})
	var res strings.Builder
	for i:=0;i<len(nums);i++{
		res.WriteString(fmt.Sprintf("%d",nums[i]))
	}
	return res.String()
}