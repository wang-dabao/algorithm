package main

import "fmt"

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 示例 1：
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//输入：n = 1
//输出：["()"]
var res  []string
var cur string
func generateParenthesis(n int) []string {
	//思路，回溯法，
	res = []string{}
	pinjie(cur,0,0,n)
	return res
}
func pinjie(cur string, left,right,n int) {
	if len(cur) == 2 * n {
		res = append(res,cur)
		cur = ""
	}
	if left < n {
		cur = cur + "("
		pinjie(cur, left+1, right, n)
		cur = cur[:len(cur)-1]
	}
	if right < left {
		cur = cur + ")"
		pinjie(cur,left,right+1,n)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	fmt.Print(generateParenthesis(3))
}

