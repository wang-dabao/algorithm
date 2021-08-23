package main

import (
	algorithm "algorithm/const"
	"fmt"
)

/**
输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
B是A的子结构， 即 A中有出现和B相同的结构和节点值。
例如:
给定的树 A:
         3
        / \
       4   5
      / \
     1   2
给定的树 B：
      4
     /
    1
返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。
示例 1：
输入：A = [1,2,3], B = [3,1]
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof
 */

/**
思路： 将树A进行前序遍历，从A的根节点开始，每一个子树作为一个新树，去跟B 比较，是否完全包含于跟B
 */

func isSubStructure(A *algorithm.TreeNode, B *algorithm.TreeNode) bool {
	var include func(A *algorithm.TreeNode, B *algorithm.TreeNode) bool
	include = func(A *algorithm.TreeNode, B *algorithm.TreeNode) bool {
		if B == nil {
			return true
		}
		if A == nil || A.Val != B.Val {
			return false
		}
		return include(A.Left,B.Left) && include(A.Right,B.Right)
	}
	if A == nil || B == nil {
		return false
	}
	return include(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}

func main() {
	fmt.Println(isSubStructure(&algorithm.TreeNode{Val: 1, Left:&algorithm.TreeNode{Val: 2,Left:&algorithm.TreeNode{Val: 4}},Right: &algorithm.TreeNode{Val: 3}},
	&algorithm.TreeNode{Val: 3}))
}