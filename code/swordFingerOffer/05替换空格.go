package main

/**
请实现一个函数，把字符串 s 中的每个空格替换成"%20"。
示例 1：
输入：s = "We are happy."
输出："We%20are%20happy."

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ti-huan-kong-ge-lcof
 */

/**
思路：这个我觉得不需要啥思路了，放心，面试肯定不会考这白给题
 */

func replaceSpace(s string) string {
	var res []byte
	for i:=0; i<len(s); i++ {
		if s[i] == ' ' {
			res = append(res,'%','2','0')
		}else {
			res = append(res,s[i])
		}
	}
	return string(res)
}