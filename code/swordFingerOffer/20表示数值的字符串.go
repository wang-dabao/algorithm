package main

import "fmt"

/**
请实现一个函数用来判断字符串是否表示数值（包括整数和小数）。数值（按顺序）可以分成以下几个部分：
1.若干空格
2.一个小数或者整数
3.（可选）一个'e'或'E'，后面跟着一个整数
4.若干空格
小数（按顺序）可以分成以下几个部分：
	1.（可选）一个符号字符（'+' 或 '-'）
	2.下述格式之一：
		至少一位数字，后面跟着一个点 '.'
		至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
		一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：
	1.（可选）一个符号字符（'+' 或 '-'）
	2.至少一位数字
部分数值列举如下：
["+100", "5e2", "-123", "3.1416", "-1E-16", "0123"]
部分非数值列举如下：
["12e", "1a3.14", "1.2.3", "+-5", "12e+5.4"]

示例 1：
输入：s = "0"
输出：true
示例 2：
输入：s = "e"
输出：false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof
 */

/**
思路：自己写了好久，都要吐了....还是直接看答案吧
 */

type State int
type CharType int

const (
	StateInitial State = iota
	StateIntSign
	StateInteger
	StatePoint
	StatePointWithoutInt
	StateFraction
	StateExp
	StateExpSign
	StateExpNumber
	StateEnd
)

const (
	CharNumber CharType = iota
	CharExp
	CharPoint
	CharSign
	CharSpace
	CharIllegal
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CharNumber
	case 'e', 'E':
		return CharExp
	case '.':
		return CharPoint
	case '+', '-':
		return CharSign
	case ' ':
		return CharSpace
	default:
		return CharIllegal
	}
}

func isNumber(s string) bool {
	transfer := map[State]map[CharType]State{
		StateInitial: {
			CharSpace:  StateInitial,
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
			CharSign:   StateIntSign,
		},
		StateIntSign: {
			CharNumber: StateInteger,
			CharPoint:  StatePointWithoutInt,
		},
		StateInteger: {
			CharNumber: StateInteger,
			CharExp:    StateExp,
			CharPoint:  StatePoint,
			CharSpace:  StateEnd,
		},
		StatePoint: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StatePointWithoutInt: {
			CharNumber: StateFraction,
		},
		StateFraction: {
			CharNumber: StateFraction,
			CharExp:    StateExp,
			CharSpace:  StateEnd,
		},
		StateExp: {
			CharNumber: StateExpNumber,
			CharSign:   StateExpSign,
		},
		StateExpSign: {
			CharNumber: StateExpNumber,
		},
		StateExpNumber: {
			CharNumber: StateExpNumber,
			CharSpace:  StateEnd,
		},
		StateEnd: {
			CharSpace: StateEnd,
		},
	}
	state := StateInitial
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == StateInteger || state == StatePoint || state == StateFraction || state == StateExpNumber || state == StateEnd
}

func main() {
	fmt.Println(isNumber("3.  "))
}