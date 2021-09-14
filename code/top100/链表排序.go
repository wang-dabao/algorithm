package main

import (
	"algorithm/code/const"
)

//给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
// 示例 1：
//输入：head = [4,2,1,3]
//输出：[1,2,3,4]
// 示例 2：
//输入：head = [-1,5,3,4,0]
//输出：[-1,0,3,4,5]


func sortList(head *algorithm.ListNode) *algorithm.ListNode {
	//思路:链表排序，核心思想是归并排序，把链表的每个节点，当做是最小的粒度，之后合并两个有序链表
	//合并两个有序链表方法
	merge := func(l1,l2 *algorithm.ListNode) *algorithm.ListNode {
		//定义2个指针
		pre := &algorithm.ListNode{0,nil}
		cur := pre
		for l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
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
		return pre.Next
	}

	//排序..需要找到mid 分区点，这里就需要知道链表的中心点，可以用快慢指针
	var sort func(head,tail *algorithm.ListNode) *algorithm.ListNode
	sort = func(head, tail *algorithm.ListNode) *algorithm.ListNode {
		if head == nil  {
			return head
		}
		//这是为什么呢？ 主要是因为区间的确认是左闭右开的，不包括tail的
		if head.Next == tail  {
			head.Next = nil	//mark： 特别注意这个处理，将大单链表拆分成小的单链表
			return head
		}
		//快慢指针找中间值：左指针每次向后移1个位置；右指针移2个位置（直到尾结点）；
		left,right := head,head
		for right != tail {
			left = left.Next
			right = right.Next
			if right != tail {
				right = right.Next
			}
		}
		mid := left
		return merge(sort(head,mid),sort(mid,tail))
	}

	return sort(head,nil)
}

