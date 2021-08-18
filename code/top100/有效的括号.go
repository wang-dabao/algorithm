package main

import "fmt"

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 示例 1：
//输入：s = "()"
//输出：true
// 示例 2：
//输入：s = "()[]{}"
//输出：true
// 示例 3：
//输入：s = "(]"
//输出：false
// 示例 4：
//输入：s = "([)]"
//输出：false
// 示例 5：
//输入：s = "{[]}"
//输出：true
// 1 <= s.length <= 104
// s 仅由括号 '()[]{}' 组成

func isValid(s string) bool {
	//思路：辅助栈，利用栈的先进后出的特点，我们遍历字符串按顺序 将符合'(' '{' '[' 放入栈中，之后 当遇到括号的右半部分，就取出最后一个 （其实数组也是可以的，下面就用数组做）
	var nums []uint8
	//用map来存储每个符号对应的符号，以便比较
	var stMap = map[uint8]uint8{'(':')','[':']','{':'}'}
	for i := 0; i < len(s); i++ {
		//如果map存在，也就是当字符为 '(' '[' '{' 时，将他对应的字符存储在数组中，
		if stMap[s[i]] > 0  {
			nums = append(nums,stMap[s[i]])
		}else {
			//反之，如果是')' ']' '}' 字符，那就需要比较 如果这时数组是空的，说明上来就是反字符，肯定false 数组不是空的，但是此时的字符和数组的最后一个不相等，那也是false
			if len(nums) == 0 || s[i] != nums[len(nums)-1] {
				return false
			}
			//每一次比较完之后，如果不返回，说明匹配上了，需要将数组最后一位字符去掉，以便下一次比较，（模拟弹栈）
			nums = nums[:len(nums)-1]
		}
	}
	if  len(nums) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Print(isValid("{[]}"))
}

