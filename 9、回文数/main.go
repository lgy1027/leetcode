package main

import (
	"fmt"
	"strconv"
)

/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/
func main() {
	x := 123454321
	fmt.Println(isPalindrome(x))
}

//首先将整数转换成字符串
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	j := len(str) - 1
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[j] {
			return false
		}
		j--
	}
	return true
}

func isPalindromeTwo(x int) bool {
	if x >= 0 {
		if x > 0 && x%10 == 0 {
			return false
		}
		var y, m int
		for {
			m = x % 10
			y = y*10 + m
			if y >= x || y >= x/10 {
				if y == x || y == x/10 {
					return true
				}
				return false
			}
			x = x / 10
		}
	}
	return false
}
