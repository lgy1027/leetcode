package main

import "fmt"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
*/

func main() {
	x := -345
	result := reverse(x)
	fmt.Println(result)
}

func reverse(x int) int {
	var nums, newnums int
	for x != 0 { //直到x等于0，跳出循环
		a := x % 10
		newnums = nums*10 + a

		nums = newnums
		x = x / 10
		MaxInt32 := 1<<31 - 1
		MinInt32 := -1 << 31
		if nums > MaxInt32 || nums < MinInt32 {
			return 0
		}
	}

	return nums
}
