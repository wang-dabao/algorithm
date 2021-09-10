package main

import (
	"algorithm/code/const"
)

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
示例:
给定如下二叉树，以及目标和 target = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]

*/

func pathSum(root *algorithm.TreeNode, targetSum int) [][]int {
	//思路：dfs 深度优先遍历 并且用 tag 记录每一条路径的和，当等于目标值的时候，将路径放入二维数组中
	var ret [][]int
	var dfs func(*algorithm.TreeNode,int)
	var path []int
	dfs = func(root *algorithm.TreeNode,tag int) {
		if root == nil {
			return
		}
		tag += root.Val
		path = append(path,root.Val)
		//注意这个是回溯法的关键，当返回的时候，一定要回到上个节点，那么对应的path要去掉最后一位
		defer func() {path = path[:len(path)-1]}()
		if root.Right == nil && root.Left == nil && tag == targetSum {
			ret = append(ret,append([]int{},path...))
			return
		}
		dfs(root.Left,tag)
		dfs(root.Right,tag)
	}
	dfs(root,0)
	return ret
}