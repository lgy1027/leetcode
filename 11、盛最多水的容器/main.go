package main

import "fmt"

/*
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}

func maxArea(height []int) int {
	s := make([]int, 1)
	for i := 0; i <= len(height); i++ {
		for j := 0; j < len(height); j++ {
			if i < j {
				area := (j - i) * (minNum(height[i], height[j]))
				if area > s[0] {
					s[0] = area
				}
			}
		}
	}
	return s[0]
}

func minNum(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}
