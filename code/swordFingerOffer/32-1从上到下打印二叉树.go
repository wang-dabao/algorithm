package main

import (
	"algorithm/code/const"
)

/**
从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
例如:
给定二叉树:[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回：
[3,9,20,15,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
 */

/**
思路：二叉树的层序遍历.. bfs广度优先遍历，核心是用队列存储每一层的子树
 */

func levelOrder(root *algorithm.TreeNode) []int {
	var res []int
	var queue []*algorithm.TreeNode
	queue = append(queue,root)
	for len(queue) > 0 {
		var p []*algorithm.TreeNode
		for _, node := range queue {
			if node != nil {
				res = append(res,node.Val)
				p = append(p,node.Left)
				p = append(p,node.Right)
			}
		}
		queue = p
	}
	return res
}