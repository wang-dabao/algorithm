package main

import algorithm "algorithm/const"

/**
输入两个链表，找出它们的第一个公共节点。
两个链表相交点
示例 1：
输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
输出：Reference of the node with value = 8
输入解释：相交节点的值为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof
 */

/**
思路：我们第一时间能想到的是 将A遍历一遍，拿map装起来，之后再遍历B 如果B的某一个节点在map中存在，那么他们就相交，不存在就不想交 这样时间复杂度是o(A+B) 空间复杂度o(A)
还有一种比较巧妙的双指针方式，空间复杂度可以降到o(1)
双指针，假设AB两个链表相交 那么A链表相交之前的长度为a,相交部分的长度为c B链表相交之前的长度为b 这样！A=a+c B=b+c
巧妙的点来了！我们遍历AB两个链表，当A遍历完了我们将A指针指向B链表的开头 同理 B链表遍历完了 我们将B指针指向A链表的开头，如果当A指针和B指针的值相等，说明链表相交。如果最后A指针和B指针都指向nil，说明不想交
为什么？ 因为如果链表相交，那么一定有 a+c+b=b+c+a，两指针在相交点相遇了，如果不想交那么 a+c+b+c=b+c+a+c也就是都把AB链表遍历一遍，最后都指向nil 时间复杂度还是o(A+B)
 */

func getIntersectionNode(headA, headB *algorithm.ListNode) *algorithm.ListNode {
	cura,curb := headA,headB
	if headA == nil || headB == nil{
		return nil
	}
	for cura != nil || curb != nil {
		if cura == nil {
			cura = headB
		}
		if curb == nil {
			curb = headA
		}
		if cura == curb {
			return cura
		}
		cura = cura.Next
		curb = curb.Next
	}
	return nil
}
