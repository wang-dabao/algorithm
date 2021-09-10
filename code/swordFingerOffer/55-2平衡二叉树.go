package main

import algorithm "algorithm/const"

/**
输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
示例 1:
给定二叉树 [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
返回 true 。
示例 2:
给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回false 。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
 */

/**
思路：递归判断整个树... 判断每一个节点的左树和右树的深度..相差超过1 为false
 */

func isBalanced(root *algorithm.TreeNode) bool {
	deep := func(root *algorithm.TreeNode) int{
		if root == nil {
			return 0
		}
		var queue []*algorithm.TreeNode
		d := 0
		queue = append(queue,root)
		for len(queue) > 0 {
			d++
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
		return d
	}
	var check func(root *algorithm.TreeNode) bool
	check = func(root *algorithm.TreeNode) bool {
		if  root == nil || (root.Right == nil && root.Left == nil) {
			return true
		}
		leftdeep := deep(root.Left)
		rightdeep := deep(root.Right)
		diff := 0
		if rightdeep > leftdeep {
			diff = rightdeep - leftdeep
		}else {
			diff = leftdeep - rightdeep
		}
		return diff < 1  && check(root.Right) && check(root.Left)
	}
	return check(root)
}