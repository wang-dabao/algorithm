package main

import algorithm "algorithm/const"

/**
从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。
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
  [9,20],
  [15,7]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
 */


func levelOrder2(root *algorithm.TreeNode) [][]int {
	var res [][]int
	var queue []*algorithm.TreeNode
	queue = append(queue,root)
	for len(queue) > 0 {
		var p []*algorithm.TreeNode
		var s []int
		for _, node := range queue {
			if node != nil {
				s = append(s,node.Val)
				p = append(p,node.Left,node.Right)
			}
		}
		if len(s) != 0 {
			res = append(res,s)
		}
		queue = p
	}
	return res
}