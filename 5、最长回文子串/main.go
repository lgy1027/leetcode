package main

import "fmt"

/*

	给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。


示例 2：

输入: "cbbd"
输出: "bb"
*/
func main() {
	str := "adghdfgfdhddd"
	result := longestPalindrome(str)
	fmt.Println(result)
}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	if isPalindrome(s) {
		return s
	}

	pdMap := make(map[string]int)
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i+1; j-- {
			tmp := s[i:j]
			//fmt.Println(tmp)
			if isPalindrome(tmp) {
				pdMap[tmp] = len(tmp)
			}
		}
	}
	max := 0
	rst := ""
	for k, v := range pdMap {
		if v > max {
			max = v
			rst = k
		}
	}
	if rst == "" {
		return s[0:1]
	}
	return rst
}

func isPalindrome(s string) bool {
	if len(s) < 1 {
		return false
	}
	if len(s) == 2 && s[0] == s[1] {
		return true
	}
	for i := 0; i < len(s)/2; i++ { // 3/2 = 1    4/2 = 2
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
