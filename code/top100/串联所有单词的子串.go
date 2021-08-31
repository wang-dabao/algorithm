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

/**
思路：
分析题，words的所有连续的排列组合是s的子字符串。那么也就是说，我们可以将s分割一下。利用滑动窗口的思想，窗口的大小应该是所有words数组中的单词组成的字符串（也就是子串应该有的长度）
这样我们每一次将起始位置向后挪一个，移动窗口，直到尾结点到字符串的最后。这样就遍历完了所有的s字符串。
再来看每一次窗口中的字符串处理，将窗口中的字符串分割成长度为len(word)的小段，将小段存储在map中，遍历map，如果和words中的元素一致，那就符合返回true
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