package main

/**
字符串的左旋转操作是把字符串前面的若干个字符转移到字符串的尾部。请定义一个函数实现字符串左旋转操作的功能。比如，输入字符串"abcdefg"和数字2，该函数将返回左旋转两位得到的结果"cdefgab"。
示例 1：
输入: s = "abcdefg", k = 2
输出:"cdefgab"
示例 2：
输入: s = "lrloseumgh", k = 6
输出:"umghlrlose"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof
 */

/**
思路：emm....没啥具体思路要写，就是既然是把前面n个拿到后面，那我们遍历字符串，从n+1个开始从新构建字符串 这样的时间复杂度是O(N) N是字符串长度..
 */

func reverseLeftWords(s string, n int) string {
	if len(s) == 0 || n == 0 {
		return s
	}
	var nums []byte
	for j := n ;j < len(s); j++ {
		nums = append(nums,s[j])
	}
	for i := 0; i < n; i++ {
		nums = append(nums,s[i])
	}
	return string(nums)
}