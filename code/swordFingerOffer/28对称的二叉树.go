package main

import algorithm "algorithm/const"

/**
请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。
例如，二叉树[1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
 2    2
  \    \
   3    3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof
 */


/**
思路：递归法 很简单，写一个比较左子树和右子树是否是对称的方法，之后从根节点一直递归下去，将整个树变成最小粒度，判断其值是否相等
当然这种破题就一定有迭代法，和二叉树便利有所不同的是，需要注意压栈的顺序
 */

func isSymmetricDigui(root *algorithm.TreeNode) bool {
	var compare func(left,right *algorithm.TreeNode) bool
	compare = func(left, right *algorithm.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left != nil && right != nil{
			return left.Val == right.Val && compare(left.Left,right.Right) && compare(left.Right,right.Left)
		}
		return false
	}
	if root == nil {
		return true
	}
	return compare(root.Left,root.Right)
}

func isSymmetricDiedai(root *algorithm.TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []*algorithm.TreeNode
	stack = append(stack,root.Right,root.Left)
	for len(stack) > 0 {
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack = append(stack,right.Left,left.Right,right.Right,left.Left)
	}
	return true
}