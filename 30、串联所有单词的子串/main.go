package main

import "fmt"

/*
给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。

注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

示例 1：

输入：
  s = "barfoothefoobarman",
  words = ["foo","bar"]
输出：[0,9]
解释：
从索引 0 和 9 开始的子串分别是 "barfoo" 和 "foobar" 。
输出的顺序不重要, [9,0] 也是有效答案。
示例 2：

输入：
  s = "wordgoodgoodgoodbestword",
  words = ["word","good","best","word"]
输出：[]
*/
func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	var res []int
	if s == "" || len(s) == 0 || words == nil || len(words) == 0 {
		return res
	}
	// words数组中每个单词字符串长度一致, 取第一个单词
	length := len(words[0])
	// 定义一个wordsMap用于存放words数组中每个单词出现的次数
	wordsMap := make(map[string]int)
	for _, w := range words {
		put(wordsMap, w)
	}
	for i := 0; i < len(s)-len(words)*length+1; i++ {
		// 定义一个window用于存放匹配的单词出现的次数
		window := make(map[string]int)
		// 定义一个num维护当前匹配的次数
		var num int
		for num < len(words) {
			// 根据查找取子串
			word := s[i+num*length : i+(num+1)*length]
			_, ok := wordsMap[word]
			if ok {
				put(window, word)
				// window中word出现的次数超过words数组中，则跳出本次循环
				if window[word] > wordsMap[word] {
					break
				}
			} else {
				// 没有查询到, 跳出本次循环, 查找下一个字符
				break
			}
			num++
		}
		if num == len(words) {
			res = append(res, i)
		}
	}
	return res
}

func put(myMap map[string]int, key string) {
	_, ok := myMap[key]
	if ok {
		myMap[key] += 1
	} else {
		myMap[key] = 1
	}
}
