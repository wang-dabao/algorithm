package main

import "fmt"

/**
两个字符串，查找最长公共子串。
 */

func findStr(s1,s2 string) string {
	res := ""
	for i:=0; i<len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				a,b := i,j
				for a<len(s1) && b<len(s2) && s1[a]==s2[b] {
					a++
					b++
				}
				if len(res) < a-i{
					res = s1[i:a]
				}
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findStr("abcdefghijklmnop","abcsafjklmnopqrstuvw"))
}
