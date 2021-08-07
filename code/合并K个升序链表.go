package main
//给你一个链表数组，每个链表都已经按升序排列。
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
// 示例 1：
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6


/**
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	//思路：在合并两个有序链表的前提下，升级版，可以用一个数组存储每一次合并后的链表，之后遍历这个数组，每取出一个链表就和数组中最后一个链表相比较。然后得到一个新的链表放入到数组中，
	//笨方法，当然还有分治法什么的..不会
	var ans *ListNode
	for i := 0; i < len(lists); i++ {
		//当第一个的时候，将他直接赋值给ans
		if i == 0 {
			ans = lists[i]
			continue
		}
		ans = twoListNode(ans, lists[i])
	}
	return ans
}
func twoListNode(l1,l2 *ListNode) *ListNode {
	per := &ListNode{0,nil}
	cur := per
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}
		//比较
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
	return per.Next
}

