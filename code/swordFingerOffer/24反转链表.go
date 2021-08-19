package main

import algorithm "algorithm/const"

/**
定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。
示例:
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-lian-biao-lcof
 */

/**
思路：超级超级经典的链表题...也比较基础，大致思路就是三个指针：空余指针指向结果链表的尾部，当前指针指向待反转链表头部，next存储还未反转链表
详细的思路可以去看top100中的反转链表
 */

func reverseList(head *algorithm.ListNode) *algorithm.ListNode {
	var pre,next *algorithm.ListNode
	cur := head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}