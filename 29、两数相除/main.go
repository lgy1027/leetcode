package main

import "fmt"

/*

给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
解释: 10/3 = truncate(3.33333..) = truncate(3) = 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
解释: 7/-3 = truncate(-2.33333..) = -2

*/
func main() {
	dividend := 12
	divisor := 3
	fmt.Println(divide(dividend, divisor))
}

func divide(dividend int, divisor int) int {
	if divisor == 0 {
		return 0
	}
	//处理符号和绝对值
	symbol := 1
	if dividend < 0 {
		dividend = -dividend
		symbol = -symbol
	}
	if divisor < 0 {
		symbol = -symbol
		divisor = -divisor
	}

	//被除数小于除数
	if dividend < divisor {
		return 0
	}

	//计算除法
	result := operation(dividend, divisor)
	//溢出时返回值
	if result == -1 {
		if symbol > 0 {
			return 1<<31 - 1
		} else {
			return -1 << 31
		}
	}

	//计算结果
	if symbol > 0 {
		return result
	} else {
		return -result
	}
}

func operation(m int, n int) int {
	move := 0
	for {
		if n<<move <= m && n<<(move+1) > m {
			break
		}
		move++
	}
	if move >= 31 {
		return -1
	}

	if m-n<<move >= n {
		return 1<<move + operation(m-n<<move, n)
	} else {
		return 1 << move
	}
}
