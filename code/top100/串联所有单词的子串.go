package main

/**
给定一个字符串s和一些 长度相同 的单词words 。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
注意子串要与words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑words中单词串联的顺序。
示例 1：
输入：s = "barfoothefoobarman", words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：
输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
输出：[]
示例 3：
输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
输出：[6,9,12]
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words
 */

func findSubstring(s string, words []string) []int {
	wordsMap := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]]++
	}
	check := func(s string,lenth int) bool{
		cahcheMap := make(map[string]int, len(s))
		for i := 0; i < len(s); i += lenth {
			cahcheMap[s[i:i+lenth]]++
		}
		for key, value := range cahcheMap {
			if v,ok := wordsMap[key]; !ok || v != value {
				return false
			}
		}
		return true
	}
	lenth := len(words[0])
	start,end := 0,lenth*len(words)-1
	var r []int
	for end < len(s) {
		if check(s[start:end+1], lenth){
			r = append(r,start)
		}
		end++
		start++
	}
	return r
}