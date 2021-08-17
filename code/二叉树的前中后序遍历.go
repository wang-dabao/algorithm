package main

import algorithm "algorithm/const"

/**
给你二叉树的根节点 root ，返回它节点值的 前序 中序 后序 遍历。
        1
       / \
      2   3
     /   / \
    4   5   6
 */

/**
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//思路：二叉树的遍历 基本的DFS(深度优先遍历) 分为递归和非递归. 这里分别用两种解法解题

//前序遍历 递归方法
func qianXuDiGui (root *algorithm.TreeNode) []int {
	//前序遍历 根——左——右
	var r []int
	var dfs func(*algorithm.TreeNode)
	dfs = func(node *algorithm.TreeNode){
		if node == nil {
			return
		}
		r = append(r,node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return r
}
//前序遍历 非递归
func qianXuFeiDiGui (root *algorithm.TreeNode) []int {
	//前序遍历 根——左——右 非递归，采用显示栈的方法，先将整个树的所有左孩子节点压入栈，利用栈的先进后出的原理
	var stack []*algorithm.TreeNode
	var r []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			r = append(r,root.Val)
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return r
}
//中序遍历 递归
func zhongXuDiGui(root *algorithm.TreeNode) []int {
	var r []int
	var dfs func(node *algorithm.TreeNode)
	dfs = func(node *algorithm.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		r = append(r,node.Val)
		dfs(node.Right)
	}
	return r
}
//中序遍历 非递归
func zhongXuFeiDiGui(root *algorithm.TreeNode) []int {
	var stack []*algorithm.TreeNode
	var r []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r,root.Val)
		root = root.Right
	}
	return r
}
//后续遍历 递归
func houXuDiGui(root *algorithm.TreeNode) []int {
	var r []int
	var dfs func(node *algorithm.TreeNode)
	dfs = func(node *algorithm.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		r = append(r,node.Val)
	}
	return r
}
//后序遍历 非递归
func houXuFeiDiGui(root *algorithm.TreeNode) []int {
	var stack []*algorithm.TreeNode
	var r []int
	var per *algorithm.TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack,root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == per {
			r = append(r,root.Val)
			per = root
			root = nil
		}else {
			stack = append(stack,root)
			root = root.Right
		}
	}
	return r
}


