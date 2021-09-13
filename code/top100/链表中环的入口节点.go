package main

import algorithm "algorithm/code/const"

/**
给定一个链表，返回链表开始入环的第一个节点。 从链表的头节点开始沿着 next 指针进入环的第一个节点为环的入口节点。如果链表无环，则返回null。
为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
说明：不允许修改给定的链表。
示例 1：
输入：head = [3,2,0,-4], pos = 1
输出：返回索引为 1 的链表节点
解释：链表中有一个环，其尾部连接到第二个节点。
示例2：
输入：head = [1,2], pos = 0
输出：返回索引为 0 的链表节点
解释：链表中有一个环，其尾部连接到第一个节点。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/c32eOV
 */

/**
思路：快慢指针..当第一次相遇的时候，说明是有环的，如果有环，就把快指针指向头，之后快指针一步一步走，再次相遇的时候，就是入口节点
 */

func detectCycle(head *algorithm.ListNode) *algorithm.ListNode {
	fast,slow := head,head
	if head == nil{
		return nil
	}
	for fast != nil || slow != nil {
		if fast == nil {
			return nil
		}else {
			fast = fast.Next
		}
		if slow == nil || fast == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
