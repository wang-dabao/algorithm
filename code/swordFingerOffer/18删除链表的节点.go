package main

import algorithm "algorithm/const"

/**
给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。返回删除后的链表的头节点。
注意：此题对比原题有改动
示例 1:
输入: head = [4,5,1,9], val = 5
输出: [4,1,9]
解释: 给定你链表中值为5的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
示例 2:
输入: head = [4,5,1,9], val = 1
输出: [4,5,9]
解释: 给定你链表中值为1的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shan-chu-lian-biao-de-jie-dian-lcof
 */

func deleteNode(head *algorithm.ListNode, val int) *algorithm.ListNode {
	per := &algorithm.ListNode{}
	cur := per
	for head != nil {
		if head.Val != val {
			cur.Next = &algorithm.ListNode{}
			cur.Next.Val = head.Val
			cur = cur.Next
		}
		head = head.Next
	}
	return per.Next
}