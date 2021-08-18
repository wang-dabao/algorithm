package main

/**
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof
 */

/**
思路：也是属于栈的变形题，我们在一个minStack中构建两个栈，一个用于存储元数据，一个用于存储最小的min值
 */

type MinStack struct {
	stack1 []int
	stack2 []int
}

// Constructor  initialize your data structure here.
func constructor() MinStack {
	return MinStack{stack1: []int{},stack2: []int{}}
}

// Push 压栈
func (minStack *MinStack) Push(x int)  {
	if len(minStack.stack1) == 0 {
		minStack.stack2 = append(minStack.stack2,x)
	}else {
		if minStack.stack2[len(minStack.stack2)-1] >= x {
			minStack.stack2 = append(minStack.stack2,x)
		}
	}
	minStack.stack1 = append(minStack.stack1,x)
}

// Pop 弹栈
func (minStack *MinStack) Pop()  {
	pop := minStack.stack1[len(minStack.stack1)-1]
	minStack.stack1 = minStack.stack1[:len(minStack.stack1)-1]
	if pop <= minStack.stack2[len(minStack.stack2)-1] {
		minStack.stack2 = minStack.stack2[:len(minStack.stack2)-1]
	}
}


func (minStack *MinStack) Top() int {
	return minStack.stack1[len(minStack.stack1)-1]
}


func (minStack *MinStack) Min() int {
	return minStack.stack2[len(minStack.stack2)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */