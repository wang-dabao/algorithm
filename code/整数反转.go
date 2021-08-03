package main

import (
	"fmt"
	"math"
)

//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
// 如果反转后整数超过 32 位的有符号整数的范围 [−231, 231 − 1] ，就返回 0。
//假设环境不允许存储 64 位整数（有符号或无符号）。
// 示例 1：
//输入：x = 123
//输出：321
// 示例 2：
//输入：x = -123
//输出：-321
// 示例 3：
//输入：x = 120
//输出：21
// 示例 4：
//输入：x = 0
//输出：0
// 提示：
// -231 <= x <= 231 - 1

func reverse(x int) int {
	//思路：数学方法，一个数反转，这个数直接除10得出余数，这个余数作为新数的起始位置 每次*10   点睛之笔（result = result * 10 + rever ）
	rever := 0
	result := 0
	for x != 0 {
		if result < math.MinInt32/10 || result > math.MaxInt32/10 {
			return 0
		}
		rever = x % 10
		x = x / 10
		result = result * 10 + rever
	}
	return result
}

func main()  {
	fmt.Print(reverse(123))
}

