package main

import (
	algorithm "algorithm/const"
)

/**
请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
  [3],
  [20,9],
  [15,7]
]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
 */

/**
思路：二叉树层序遍历的基础上，记录下层数，奇数层数放入队里里的顺序是左——右，偶数是右——左 之后遍历队里都是从后往前遍历
 */

func levelOrder3(root *algorithm.TreeNode) [][]int {
	var res [][]int
	var queue []*algorithm.TreeNode
	queue = append(queue,root)
	depth := 1
	for len(queue) > 0 {
		var p []*algorithm.TreeNode
		var s []int
		for i := len(queue) - 1; i >= 0; i-- {
			node := queue[i]
			if node != nil {
				s = append(s,node.Val)
				if depth % 2 == 0 {
					p = append(p,node.Right,node.Left)
				}else {
					p = append(p,node.Left,node.Right)
				}
			}
		}
		depth++
		if len(s) != 0 {
			res = append(res,s)
		}
		queue = p
	}
	return res
}