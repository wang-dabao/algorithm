package main

/**
请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
示例1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof
 */

/**
思路：这题最先想到是是 双指针滑动窗口去做，喜欢双指针，后节点一点一点向后走，用map存放走过的值
 */

func lengthOfLongestSubstring(s string) int {
	end,max := 0,0
	sMap := make(map[byte]int,len(s))
	for start:=0; start<len(s); start++ {
		for end<len(s) && sMap[s[end]] == 0 {
			sMap[s[end]] = 1
			end++
		}
		if max < end-start {
			max = end-start
		}
		delete(sMap,s[start])
	}
	return max
}

/**
当然还有动态规划，跟内个最长连续子数组最大和 一个思路
dp[i] i是字符串的长度 d[i] 一定是以s[i-1]为结尾的子字符串的最大长度
状态转移方程 分为两种，
1.如果s[i-1]字符和之前的已知字符串没有重复的，那么dp[i] = dp[i-1] + 1
2.如果有重复的 假设s[j-1]和s[i-1] 重复，那么 i-j 如果大于dp[i-1] 说明重复的字符已经在已知最大连续字符串之外了，dp[i] = dp[i-1] + 1 如果小于了 dp[i] = j-i了
初始值 dp[1] = 1
*/

func lengthOfLongestSubstring1(s string) int {
	var dp []int
	dp = append(dp,1)
	max := 0
	for i:=1;i<len(s)+1;i++{
		j := i-1
		for j > 0 && s[j-1] != s[i-1] {
			j--
		}
		if i-j > dp[i-1] {
			dp = append(dp,dp[i-1] + 1)
		}else {
			dp = append(dp,i-j)
		}
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}