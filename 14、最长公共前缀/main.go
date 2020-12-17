package main

import "fmt"

func main() {
	str := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(str))
}

func longestCommonPrefix(strs []string) string {
	var res string
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		j := 1
		for ; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[0][i] != strs[j][i] {
				return res
			}
		}
		if j == len(strs) {
			res += string(strs[0][i])
		}
	}
	return res

}
