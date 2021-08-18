package main

/**
用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
示例
输入：
["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
输出：[null,null,3,-1]

输入：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
输出：[null,-1,null,null,5,2]
 */

/**
思路：用两个栈实现队列，队列的特性是 先进先出，栈的特性是先进后出， 我们玩过汉诺塔的都可以想象的到，每根棍子相当于一个栈，
我们拿一个栈用于存储插入模拟队列队尾的数据(s1)，一个栈用于删除模拟队列头部的数据(s2)，
当需要做删除指令的时候，如果s2中没有元素，把s1中的所有元素取出来，放入s2中，之后弹栈s2，如果s2中有元素就弹栈，这样就将最开始进入队列的数据删除
至此，就完美的用两个栈实现了队列的功能
 */

// CQueue 模拟队列
type CQueue struct {
	s1,s2 []int
}

func Constructor() CQueue {
	return CQueue{
		s1: []int{},
		s2: []int{},
	}
}

// AppendTail 添加队尾元素
func (cqueue *CQueue) AppendTail(value int)  {
	cqueue.s1 = append(cqueue.s1,value)
}

// DeleteHead 删除头部元素
func (cqueue *CQueue) DeleteHead() int {
	res := -1
	if len(cqueue.s2) == 0 {
		for len(cqueue.s1) != 0 {
			cqueue.s2 = append(cqueue.s2,cqueue.s1[len(cqueue.s1)-1])
			cqueue.s1 = cqueue.s1[:len(cqueue.s1)-1]
		}
	}
	if len(cqueue.s2) > 0 {
		res = cqueue.s2[len(cqueue.s2)-1]
		cqueue.s2 = cqueue.s2[:len(cqueue.s2)-1]
	}
	return res
}

