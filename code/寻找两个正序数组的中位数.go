package main

import "fmt"

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 示例 1：
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
// 示例 3：
//输入：nums1 = [0,0], nums2 = [0,0]
//输出：0.00000
// 示例 4：
//输入：nums1 = [], nums2 = [1]
//输出：1.00000
// 示例 5：
//输入：nums1 = [2], nums2 = []
//输出：2.00000
// 提示：
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？
// Related Topics 数组 二分查找 分治


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//两个从小到大的数组长度是已知的 len = len(nums1) + len(nums2)，所以基本合并数组的中位数的位置会得出 k = len / 2 ，之后就变成了寻找第k个最小的数
	lenght := len(nums1) + len(nums2)
	k , yu := lenght / 2 , lenght % 2
	end := 0
	start := 0
	//前一个数
	var r1 int
	//当前的数
	var r2 int
	//分为找两个指针，指向两个数组，比较两指针所指向的值的大小，小的内个向后移动，直到将k找完
	for i := 0; i <= k; i++ {
		r1 = r2
		if start < len(nums1) && (end >= len(nums2) || nums1[start] < nums2[end]) {
			r2 = nums1[start]
			start++
		}else {
			r2 = nums2[end]
			end++
		}
	}
	if yu == 0 {
		return float64(r1 + r2) / 2.0
	}
	return float64(r2)
}

func main()  {
	fmt.Printf("re: %+v" ,findMedianSortedArrays([]int{1,2},[]int{3,4}))
}

