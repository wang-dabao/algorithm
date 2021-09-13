package main

import "fmt"

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：=
对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
如果 s 中存在这样的子串，我们保证它是唯一的答案。
示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"
示例 2：

输入：s = "a", t = "a"
输出："a"
示例 3:

输入: s = "a", t = "aa"
输出: ""
解释: t 中两个字符 'a' 均应包含在 s 的子串中，
因此没有符合条件的子字符串，返回空字符串。=
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-window-substring
 */

/**
思路：使用滑动窗口，先用tmap将t字符存储起来,重复的累加,之后双指针，后指针移动，将走过的s子串记录下来，也存在smap中，如果遍历tmap其中的值都=tmap中的数量，这就说明已经
 */

func minWindow(s string, t string) string {
	smap,tmap := make(map[byte]int),make(map[byte]int)
	check := func() bool{
		for i,v := range tmap {
			if sv,ok := smap[i]; !ok || sv < v {
				return true
			}
		}
		return false
	}
	for i := 0; i<len(t);i++ {
		tmap[t[i]]++
	}
	start,end := 0,0
	res := ""
	for start<=end && start < len(s) {
		for check() && end < len(s){
			smap[s[end]]++
			end++
		}
		if !check() {
			if start == 0 {
				res = s[start:end]
			}else if end-start < len(res) && end <= len(s) {
				res = s[start:end]
			}
		}
		smap[s[start]]--
		start++
	}
	return res
}

func main() {
	fmt.Println(minWindow("a","aa"))
}