package main

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例 1：
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
// 示例 2：
//输入：l1 = [], l2 = []
//输出：[]
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//思路：采用一个空余指针，指向合并后的链表的开头，这里将合并后的链表开头复制0 存放初始指针位置，之后分别比较两个链表的每一位，将小的内一位复制到合并链表的位置
	pre := &ListNode{Val: 0}
	cur := pre
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}
		//比较大小,
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		}else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	//时间复杂度O(M+N)
	return pre.Next
}

