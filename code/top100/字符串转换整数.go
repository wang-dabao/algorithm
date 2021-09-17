package main

import (
	"fmt"
	"math"
)

//请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
// 读入字符串并丢弃无用的前导空格
// 检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
// 读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
// 将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤
//2 开始）。
// 如果整数数超过 32 位有符号整数范围 [−231, 231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固
//定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
// 返回整数作为最终结果。
// 注意：
// 本题中的空白字符只包括空格字符 ' ' 。
// 除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
// 示例 1：
//输入：s = "42"
//输出：42
//解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//第 1 步："42"（当前没有读入字符，因为没有前导空格）
//         ^
//第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//         ^
//第 3 步："42"（读入 "42"）
//           ^
//解析得到整数 42 。
//由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。


func myAtoi(s string) int {
	qiandao,zhengfuTag := true,true
	zhengfu := 1
	result := 0
	for i := 0; i < len(s); i++ {
		//判断是否是签到空格
		if i == 0 || qiandao{
			if s[i] != ' ' {
				qiandao = false
			}else {
				continue
			}
		}
		//判断正负号
		if s[i] == '-' && zhengfuTag{
			zhengfu = -1
			zhengfuTag = false
			continue
		}
		if s[i] == '+' && zhengfuTag {
			zhengfuTag = false
			continue
		}
		//读数了
		if s[i] >= '0' && s[i] <= '9' {
			qiandao = false
			zhengfuTag = false
			result = result * 10 + int(s[i]-'0')
			if result * zhengfu >= math.MaxInt32 {
				result =  math.MaxInt32
			}
			if result * zhengfu <= math.MinInt32 {
				result = math.MinInt32 * -1
			}
			continue
		}else {
			return result * zhengfu
		}
	}
	return result * zhengfu
}
func main() {
	s := "-91283472332"
	fmt.Printf("%+v,%T",myAtoi(s),myAtoi(s))
}

