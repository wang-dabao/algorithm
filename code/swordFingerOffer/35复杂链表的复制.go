package main

import algorithm "algorithm/const"

/**
请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]
示例 2：
输入：head = []
输出：[]
解释：给定的链表为空（空指针），因此返回 null。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof
 */


/**
思路：采用 回溯法+hashmap 将链表的每一个节点单独复制..之后再一点一点拼接，hashmap中存储节点的值，key是原链表的指针
 */

/**
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *algorithm.Node) *algorithm.Node {
	chacheMap := make(map[*algorithm.Node]*algorithm.Node)
	var deepcopy func(head *algorithm.Node) *algorithm.Node
	deepcopy = func(head *algorithm.Node) *algorithm.Node {
		if head == nil {
			return nil
		}
		if value,ok := chacheMap[head]; ok {
			return value
		}
		newNode := &algorithm.Node{Val: head.Val}
		chacheMap[head] = newNode
		newNode.Next = deepcopy(head.Next)
		newNode.Random = deepcopy(head.Random)
	    return newNode
	}
	return deepcopy(head)
}