package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度

示例 1:

输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:

输入: s = ""
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

func main() {
	s := "jhyuy"
	result := lengthOfLongestSubstring(s)
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	str := []byte(s)
	dict := make([]int, 256)
	start := 0
	result := 0
	dict[0] = 0
	for i := 0; i < len(str); i++ {
		if dict[str[i]] >= start {
			start = dict[str[i]]
		}

		if result < i-start+1 {
			result = i - start + 1
		}
		dict[str[i]] = i + 1
	}
	return result
}
