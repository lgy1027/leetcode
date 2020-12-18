package main

import "fmt"

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母
*/
func main() {
	digits := "233"
	fmt.Println(letterCombinations(digits))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	s := m[digits[0]]
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, s[j]+m[digits[i]][k])
			}
		}
		s = temp
	}
	return s
}
