package main

/**
给你一个只包含 '('和 ')'的字符串，找出最长有效（格式正确且连续）括号子串的长度。
示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"
示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
示例 3：
输入：s = ""
输出：0
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
 */

/**
思路：利用辅助栈。保持栈底元素是"最后一个未匹配的有括号的下标"，为什么呢，一个边界的问题，就是当这个又括号没有匹配元素了，
那么说明，这个右括号左边的内些，已经不能够和它右边的组成连续的了，所以，栈底放置的是左边界
对于遇到的每个 "（ " ，我们将它的下标放入栈中
 */
func longestValidParentheses(s string) int {
	maxAns := 0
	var stack []int
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if maxAns < i - stack[len(stack)-1] {
					maxAns = i - stack[len(stack)-1]
				}
			}
		}
	}
	return maxAns
}
