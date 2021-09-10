package main

import (
	"algorithm/code/const"
)

/**
输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
示例:
给定如下二叉树，以及目标和target = 22，

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
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
 */

/**
思路：深度优先遍历 dfs 记录所走的路径和路径和，如果与目标值相等，放入数组中
 */

func pathSum(root *algorithm.TreeNode, target int) [][]int {
	var res [][]int
	var dfs func(root *algorithm.TreeNode,sum int, path []int)
	dfs = func(root *algorithm.TreeNode,sum int,path []int) {
		if root == nil {
			return
		}
		sum = sum + root.Val
		path = append(path,root.Val)
		if root.Right == nil && root.Left == nil && sum == target {
			res = append(res,append([]int{},path...))
			return
		}
		dfs(root.Left,sum,path)
		dfs(root.Right,sum,path)
	}
	dfs(root,0,[]int{})
	return res
}
