package main
//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
// 示例 4:
//输入: s = ""
//输出: 0
// 提示：
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
// Related Topics 哈希表 字符串 滑动窗口

func lengthOfLongestSubstring(s string) int {
	//采用滑动窗口 首先定义起始指针和结束指针，最初双双都指向头，结束指针一点一点向后移动，用map记录路过的字符，
	//当map中存在当前结束指针指向的字符时，说明有重复的字符，记录此时起始指针和结束指针的间隔，就是字符串的距离
	//当有重复字符时，将起始指针向后推移一个，同时将之前位置的字符从map中剔除，结束节点继续
	end, sum := 0,0
	lenght := len(s)
	remap := make(map[byte]int)
	//前节点向后一步一步走
	for start := 0 ; start < lenght ; start++ {
		for end < lenght && remap[s[end]] == 0{
			remap[s[end]]++
			end++
		}
		sum = max(sum , end-start)
		//将map中的起始指针当前值去掉
		delete(remap,s[start])
	}
	return sum
}
func max(sum1,sum2 int) int {
	if sum1 > sum2{
		return sum1
	}
	return sum2
}

