package main

/**
在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
示例:
s = "abaccdeff"
返回 "b"
s = ""
返回 " "

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof
 */


func firstUniqChar(s string) byte {
	m := make(map[byte]bool)
	for i:=0;i<len(s);i++ {
		if _,ok := m[s[i]];ok{
			m[s[i]] = false
		}else {
			m[s[i]] = true
		}
	}

	for j:=0;j<len(s);j++ {
		if m[s[j]] {
			return s[j]
		}
	}
	return ' '
}