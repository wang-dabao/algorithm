package main

import algorithm "algorithm/const"

//给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
// 示例 1：
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
// 示例 2：
//输入：l1 = [0], l2 = [0]
//输出：[0]
// 示例 3：
//输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
//输出：[8,9,9,9,0,0,0,1]
// 提示：
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
// Related Topics 递归 链表 数学

func addTwoNumbers(l1 *algorithm.ListNode, l2 *algorithm.ListNode) *algorithm.ListNode {
	//思路：循环递归，把每个连链表的首节点拿出来，做加法，进位用index存储起来，参与下一位的运算。这里主要是链表的一个输出，用一个tag指针纸箱结果链表，一次一次移动，直到最后
	var result *algorithm.ListNode
	var tag *algorithm.ListNode
	index := 0
	for l1 != nil || l2 != nil {
		l1Value := 0
		l2Value := 0
		if l1 != nil {
			l1Value = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2Value = l2.Val
			l2 = l2.Next
		}
		v := (l1Value + l2Value + index) % 10
		index = (l1Value + l2Value + index) / 10
		if result == nil {
			result = &algorithm.ListNode{Val: v}
			tag = result
		}else {
			tag.Next = &algorithm.ListNode{Val: v}
			tag = tag.Next
		}
	}
	if index > 0 {
		tag.Next =  &algorithm.ListNode{Val: index}
	}
	return result
}

