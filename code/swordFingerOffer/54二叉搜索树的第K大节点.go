package main

import algorithm "algorithm/const"

/**
给定一棵二叉搜索树，请找出其中第k大的节点。
示例 1:
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 4
示例 2:
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 4
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof
 */

/**
思路：二叉搜索树的特性是  中序遍历之后，能得到一个有序递增数组...那么问题就简单那了
中序遍历，写个非递归的吧
 */

func kthLargest(root *algorithm.TreeNode, k int) int {
	var nums [] int
	var stack []*algorithm.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums,node.Val)
		root = node.Right
	}
	return nums[len(nums)-k]
}