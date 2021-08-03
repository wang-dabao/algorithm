package main

//给你一个字符串 s，找到 s 中最长的回文子串。

// 示例 1：
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
// 示例 2：
//输入：s = "cbbd"
//输出："bb"
// 示例 3：
//输入：s = "a"
//输出："a"
// 示例 4：
//输入：s = "ac"
//输出："a"
// 1 <= s.length <= 1000
// s 仅由数字和英文字母（大写和/或小写）组成
// Related Topics 字符串 动态规划

func longestPalindrome(s string) string {
	//思路： 采用中心扩展法去解题 因为回文字符串前后对称的特性，分别遍历字符串的每个位置，将其作为中心点，用连个指针分别向前后移动，比较字符是否相等，注意：要兼容奇数回文和偶数回文
	result := ""
	for i := 0 ; i < len(s) ; i++ {
		//假设此中心点为i的回文串是奇数回文
		s1 := yidong(s, i, i)
		//假设此中心点为i的回文串是偶数回文
		s2 := yidong(s, i, i+1)
		if len(s1) > len(result) {
			result = s1
		}
		if len(s2) > len(result) {
			result = s2
		}
	}
	return result  //时间复杂度O(n^2) 空间复杂度O(1)
}
func yidong(s string, left,right int) string {
	for (left >= 0 && right < len(s)) && s[left] == s[right] {
		left--
		right++
	}
	//当遇到不一致的字符之后，左右指针应该返回上一次的位置
	return s[left+1:right-1+1]
}

