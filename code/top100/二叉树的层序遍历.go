package main

import algorithm "algorithm/const"

/**
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
   3
  / \
 9  20
   /  \
  15   7
返回其层序遍历结果：
[
 [3],
 [9,20],
 [15,7]
]
 */

/**
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *algorithm.TreeNode) [][]int {
	//思路： 层序遍历，典型的广度优先遍历。核心点事哪一个队列，存储二叉树的每一层，之后先进从队列取出来每一层，再放入二维数组中，
	var re [][]int
	if root == nil{
		return re
	}
	queen := []*algorithm.TreeNode{root}
	index := 0
	for len(queen) > 0 {
		re = append(re, []int{})
		var per []*algorithm.TreeNode
		for j := 0; j < len(queen); j++ {
			root = queen[j]
			re[index] = append(re[index],root.Val)
			if root.Left != nil {
				per = append(per,root.Left)
			}
			if root.Right != nil {
				per = append(per,root.Right)
			}
		}
		index++
		queen = per
	}
	return re
}

