package main

import algorithm "algorithm/const"

/**
输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
例如：
给定二叉树 [3,9,20,null,null,15,7]，
    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度3
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-de-shen-du-lcof
 */

/**
思路：二叉树的层序遍历。广度优先遍历..之后记录树的深度就可以啦
 */

func maxDepth(root *algorithm.TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*algorithm.TreeNode
	deep := 0
	queue = append(queue,root)
	for len(queue) > 0 {
		deep++
		var p []*algorithm.TreeNode
		for _, node := range queue {
			if node.Left != nil {
				p = append(p,node.Left)
			}
			if node.Right != nil {
				p = append(p,node.Right)
			}
		}
		queue = p
	}
	return deep
}
