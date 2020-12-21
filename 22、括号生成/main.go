package main

import "fmt"

/*
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。



示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
*/
func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	nums := make([]string, 0)
	generate("", &nums, n, n)
	return nums
}

func generate(s string, nums *[]string, left int, right int) {
	if left == 0 && right == 0 {
		*nums = append(*nums, s)
		return
	}

	if left > 0 {
		generate(s+"(", nums, left-1, right)
	}
	if right > left {
		generate(s+")", nums, left, right-1)
	}
}
