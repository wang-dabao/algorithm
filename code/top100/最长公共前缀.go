package main
//编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 示例 1：
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 示例 2：
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
// 0 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] 仅由小写英文字母组成

func longestCommonPrefix(strs []string) string {
	//思路： 两种办法，横向查找，和纵向查找
	//横向查找：循环遍历这个数组，默认以第一个认为是最开始的公共前缀，让其和第二个字符串相比较，得出一个公共前缀，在与第三个字符串比，这样最后得出
	//纵向查找：双重循环，外层循环默认是遍历第一个字符串，之后内层循环从数组的下标1开始，向后，依次比较每个字符串相同位置的字符是否一致，当存在不一致的时候，或者有字符串遍历完了，那么就得出了
	//我用纵向查找
	if len(strs) == 0 {
		return ""
	}
	str := strs[0]
	for i := 0; i < len(str); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || str[i] != strs[j][i] {
				return str[0:i]
			}
		}
	}
	return str  //时间复杂度O(m*n)
}

