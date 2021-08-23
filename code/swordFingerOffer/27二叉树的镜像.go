package main

import algorithm "algorithm/const"

/**
请完成一个函数，输入一个二叉树，该函数输出它的镜像。
例如输入：
        4
       / \
      2  7
    / \  / \
   1  3  6  9
镜像输出：
        4
       / \
      7  2
    / \  / \
   9  6  3  1
示例 1：
输入：root = [4,2,7,1,3,6,9]
输出：[4,7,2,9,6,3,1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-de-jing-xiang-lcof
 */

/**
思路：最简单的方法是递归....也就是一次一次调用，当最小粒度的时候，把左右指针互换位置。
还有就是辅助栈迭代的思路。利用栈的现金后出，层序遍历整个树，之后每一次弹栈构建一个新的树
分别写一下，这也算比较经典的二叉树的题
 */
//递归法
func mirrorTreeDigui(root *algorithm.TreeNode) *algorithm.TreeNode {
	if root == nil {
		return root
	}
	left := mirrorTreeDigui(root.Left)
	right := mirrorTreeDigui(root.Right)
	root.Left = right
	root.Right = left
	return root
}
//迭代法
func mirrorTreeDiedai(root *algorithm.TreeNode) *algorithm.TreeNode {
	if root == nil {
		return root
	}
	var stack []*algorithm.TreeNode
	stack = append(stack,root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack,node.Left)
		}
		if node.Right != nil {
			stack = append(stack,node.Right)
		}
		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp
	}
	return root
}