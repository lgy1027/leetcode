package main

import "fmt"

/*
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回 -1。

例如：
输入: haystack = "hello", needle = "ll"
输出: 2
输入: haystack = "aaaaa", needle = "bba"
输出: -1

说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

func main() {
	haystack := "hello"
	needle := "llo"
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack, needle string) int {
	//特殊情况，当needle是空字符串，那就返回0
	if needle == "" {
		return 0
	}

	if haystack == needle {
		return 1
	}

	//整段判断是否相同
	//判断的时候需要注意数组越界，所以小于haystack长度减去needle的长度
	//判断的时候也需要注意两个长度是一样的
	nlen := len(needle)
	for i := 0; i <= len(haystack)-nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}
