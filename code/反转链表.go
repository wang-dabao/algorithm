package main
/**
描述
输入一个链表，反转链表后，输出新链表的表头。

示例1
输入：
{1,2,3}
返回值：
{3,2,1}
 */

func ReverseList(head *ListNode) *ListNode {
	/**
	思路： 反转列表嘛， 主要核心思想就是 三个指针 pre、 cur、 next  详细说下解题思路
	pre: 指向新的链表的尾结点，（因为反转链表，我们遍历的是原始链表，那么最开始的数是反转后的链表的最后一位）初始是null
	cur: 指向原始链表的头节点 cur = head
	next: 用于存储原始链表剩余的链表（因为每一次拿出一个数，剩余的链表需要找个位置存储）
	循环逻辑：（画图更好理解）
	首先，把原始列表中抛除首节点的剩余链表存储起来，供下一次使用。next = cur->next
	然后，将cur所表示的链表的后续赋值为pre (相当于在链表pre前面加上cur当前指针指的数) cur->next = pre
	然后移动 pre和cur两个指针， 将pre向前移动一个，因为现在pre是在cur后面一个位置，所以将pre指向cur就可以 即：pre = cur ； 将cur指向原始链表剩余的链表的头节点，即：cur = next
	以此循环，就相当于 cur在原始链表上从前往后一位一位的移动，pre在一个新链表从后往前一位一位的移动
	当 cur = nil 了说明反转完成
	 */
	cur := head
	var pre,next *ListNode
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}