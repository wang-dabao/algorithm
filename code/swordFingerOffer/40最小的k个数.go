package main

import "sort"

/**
输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
示例 1：
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
 */

/**
思路：这他妈面试能考就怪了... 先排序后输出被...
 */

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr) //排序
	return arr[:k+1]
}