package main

import algorithm "algorithm/const"

/**
输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

示例 1：
输入：head = [1,3,2]
输出：[2,3,1]
 */

/**
思路：第一时间想到的肯定是，先反转链表.之后遍历反转链表，之后输出到数组中....
但是也可以用栈的先进后出的特点，但是我个人感觉没啥差别....因为go中本身没有栈的概念都是用slice实现的
 */

func reversePrint(head *algorithm.ListNode) []int {
	var pre,next *algorithm.ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	var res []int
	for pre != nil {
		res = append(res,pre.Val)
		pre = pre.Next
	}
	return res
}
