package main

import "fmt"

/*
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。



示例 1：

输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：

输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"
*/

func main() {

	str := "()()())("
	fmt.Println(longestValidParentheses(str))
}

func longestValidParentheses(str string) int {
	longest := 0
	if len(str) < 2 {
		return longest
	}

	dp := make([]int, len(str))
	dp[0] = 0
	if str[1] == '(' {
		dp[1] = 0
	} else if str[0] == '(' {
		dp[1] = 2
		longest = 2
	}
	for i := 2; i < len(str); i++ {
		if str[i] == '(' {
			dp[i] = 0
			continue
		}

		if str[i-1] == '(' { // ...()
			dp[i] = dp[i-2] + 2
		} else { // ...))
			if i-1-dp[i-1] < 0 {
				dp[i] = 0
			} else if str[i-1-dp[i-1]] == '(' { // ...[(](...))
				if i-1-dp[i-1]-1 < 0 {
					dp[i] = dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2 + dp[i-1-dp[i-1]-1]
				}
			}
		}
		if dp[i] > longest {
			longest = dp[i]
		}
	}

	return longest
}
