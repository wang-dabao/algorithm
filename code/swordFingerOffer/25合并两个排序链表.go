package main

import (
	"algorithm/code/const"
)

/**
输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。
示例1：
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/he-bing-liang-ge-pai-xu-de-lian-biao-lcof
 */

/**
思路：经典的合并有序链表.......这个没啥思路说的
 */

func mergeTwoLists(l1 *algorithm.ListNode, l2 *algorithm.ListNode) *algorithm.ListNode {
	res := &algorithm.ListNode{Val: 0}
	cur := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		}else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return res.Next
}
