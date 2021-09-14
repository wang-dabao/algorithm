package main

import "fmt"

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器。
// 示例 1：
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。
// 示例 2：
//输入：height = [1,1]
//输出：1
// 示例 3：
//输入：height = [4,3,2,1,4]
//输出：16
// 示例 4：
//输入：height = [1,2,1]
//输出：2
// 提示：
//
// n = height.length
// 2 <= n <= 3 * 104
// 0 <= height[i] <= 3 * 104
// Related Topics 贪心 数组 双指针

func maxArea(height []int) int {
	//思路，滑动窗口，前后双指针，开始分别指向起点和终点，用area记录面积。每次移动相对值小的内个指针，因为面积的大小取决去矮的柱子。直到前后指针碰面
	area := 0
	h,k := 0,0
	start,end := 0,len(height)-1
	for start != end {
		if height[end] > height[start] {
			h = height[start]
			k = end-start
			start++
		}else {
			h = height[end]
			k = end-start
			end--
		}
		if area < k * h {
			area = k * h
		}
	}
	return area  //时间复杂度 O(N) 最坏情况遍历整个数组 空间O(1)
}

func main() {
	fmt.Print(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

