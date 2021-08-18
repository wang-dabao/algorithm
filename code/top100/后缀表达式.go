package main

import (
	"fmt"
	"strconv"
)

//根据 逆波兰表示法，求该后缀表达式的计算结果。
// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。
//
// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
// 示例 1：
//输入：tokens = ["2","1","+","3","*"]
//输出：9
//解释：该算式转化为常见的中缀算术表达式为：((2 + 1) * 3) = 9
// 示例 2：
//输入：tokens = ["4","13","5","/","+"]
//输出：6
//解释：该算式转化为常见的中缀算术表达式为：(4 + (13 / 5)) = 6
// 示例 3：
//输入：tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
//输出：22
//解释：
//该算式转化为常见的中缀算术表达式为：
//  ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
//= ((10 * (6 / (12 * -11))) + 17) + 5
//= ((10 * (6 / -132)) + 17) + 5
//= ((10 * 0) + 17) + 5
//= (0 + 17) + 5
//= 17 + 5
//= 22

func evalRPN(tokens []string) int {
	//思路：使用辅助栈，根据栈的先进后出的特性，以此遍历tokens将第一个运算符前面的数字都存入栈中，这样当遇到第一个运算符时，将栈顶的两个元素拿出，做运算后的值再塞回去。参与下一次运算
	//由于go中没有栈的概念，可以拿数组做
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		//每一次运算完之后，都应该把运算后的值重新塞回去，注意：需要把已经参与运算的数从数组中移除
		case "+":
			stack = append(stack[:len(stack)-2], stack[len(stack) - 2] + stack[len(stack) - 1])
			break
		case "-":
			stack = append(stack[:len(stack)-2], stack[len(stack) - 2] - stack[len(stack) - 1])
			break
		case "*":
			stack = append(stack[:len(stack)-2], stack[len(stack) - 2] * stack[len(stack) - 1])
			break
		case "/":
			stack = append(stack[:len(stack)-2], stack[len(stack) - 2] / stack[len(stack) - 1])
			break
		default:
			atoi, _ := strconv.Atoi(tokens[i])
			stack = append(stack,atoi)
		}
	}
	return stack[len(stack)-1]
}

func main() {
	fmt.Print(evalRPN([]string{"4","13","5","/","+"}))
}

