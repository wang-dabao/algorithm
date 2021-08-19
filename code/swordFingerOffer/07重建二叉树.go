package main

import algorithm "algorithm/const"

/**
输入某二叉树的前序遍历和中序遍历的结果，请构建该二叉树并返回其根节点。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof
 */

/**
思路：递归法 我们先分析，前序遍历 【根——左——右】 中序遍历【左——根——右】 也就是说，我们通过前序遍历可以轻易的得出整个树的根节点。
得出根节点后。在中序遍历中可以找到根节点。这样把中序遍历的数组分成两个部分 左子树中序遍历 和 右子树中序遍历 这样就又可以把这两个部分作为一个新的树，来进行递归，但是我们不知道新树的根节点，怎么办呢？
不管中序也好前序也好，子树的节点数量是一致的，那么。把中序遍历分割成两部分后，根据这两部分的长度，我们可以吧前序遍历也分成两部分 左子树前序遍历 和 右子树前序遍历 这样新树的根节点就轻易得出了

还有一种迭代的方法，用栈去做....有兴趣了解一下
 */

func buildTree(preorder []int, inorder []int) *algorithm.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &algorithm.TreeNode{Val: preorder[0]}
	var index int
	for index < len(inorder) {
		if inorder[index] == preorder[0] {
			break
		}
		index++
	}
	root.Left = buildTree(preorder[1:len(inorder[:index])+1],inorder[:index])
	root.Right = buildTree(preorder[len(inorder[:index])+1:],inorder[index+1:])
	return root
}