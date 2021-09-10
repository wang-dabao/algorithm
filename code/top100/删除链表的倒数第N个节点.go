package main

import (
	"algorithm/code/const"
)

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 进阶：你能尝试使用一趟扫描实现吗？
// 示例 1：
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//	 示例 2：
//输入：head = [1], n = 1
//输出：[]
// 示例 3：
//输入：head = [1,2], n = 1
//输出：[1]
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
// Related Topics 链表 双指针

func removeNthFromEnd(head *algorithm.ListNode, n int) *algorithm.ListNode {
	pre := &algorithm.ListNode{0,head}
	cur := pre
	for i := 0; i < getLength(head)-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return pre.Next
}
func getLength(head *algorithm.ListNode) int {
	length := 0
	for head != nil {
		head = head.Next
		length++
	}
	return length
}
