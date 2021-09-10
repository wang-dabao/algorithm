package main

import (
	"algorithm/code/const"
)

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。
例如，一个链表有 6 个节点，从头节点开始，它们的值依次是 1、2、3、4、5、6。这个链表的倒数第 3 个节点是值为 4 的节点。
示例：

给定一个链表: 1->2->3->4->5, 和 k = 2.
返回链表 4->5.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof
 */

/**
思路：暴力破解，就是先遍历一遍得到链表的长度，之后再遍历输出倒数第k个
当然还有别的更好的办法，快慢指针的办法，两个指针都指向头，之后快指针先走K步，之后再一起移动，
 */

func getKthFromEnd(head *algorithm.ListNode, k int) *algorithm.ListNode {
	fast,slow := head,head
	for ; k > 0 && fast !=nil ; k-- {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}