package main

import "fmt"

/**
输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
示例 1：
输入: "the sky is blue"
输出:"blue is sky the"
示例 2：

输入: " hello world! "
输出:"world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good  example"
输出:"example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fan-zhuan-dan-ci-shun-xu-lcof
 */

/**
思路：倒叙遍历字符串 将所有的单词 找出来，之后存放在数组里，之后拼接字符串
 */

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	start,end := len(s)-1,len(s)-1
	var strs []string
	for start >= 0 {
		if s[start] == ' ' && start == end{
			start--
			end--
			continue
		}
		if s[start] != ' ' {
			if start == 0 {
				strs = append(strs,s[0:end+1])
			}
			start--
		}else {
			strs = append(strs,s[start+1:end+1])
			end = start
		}
	}
	var res string
	if len(strs) == 0 {
		return ""
	}
	for _, str := range strs {
		res = res + str + " "
	}
	return res[:len(res)-1]
}

func main() {
	fmt.Println(reverseWords(" 1"))
}