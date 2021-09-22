package main

import "strings"

/**
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:
num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。
输入/输出示例：
Example 1:
Input:
num = "1432219", k = 3
Output:
"1219"
Explanation: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。
 */

/**
思路 贪心+ 单调栈
遍历字符串，逐个入栈,若新来的比栈顶小，则栈顶出栈 ,出栈相当于移除动作，需次数并与 k 比较，别删多了
移除头部零有 '0',全空，return "0",构建字符串，顺序返回整个单调栈元素.
时间复杂度：o(n+k)
空间复杂度：o(n)
 */


func removeKdigits(num string, k int) string {
	var stack []byte
	for i := range num {
		digit := num[i]
		for k > 0 && len(stack) > 0 && digit < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, digit)
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		ans = "0"
	}
	return ans
}