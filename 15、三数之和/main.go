package main

import (
	"fmt"
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
func main() {
	nums := []int{-4, -1, -1, 0, 1, 2}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	//先对数组排序
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		z := len(nums) - 1
		for z > j {
			b := nums[j]
			c := nums[z]
			if nums[i]+b+c > 0 {
				z--
			} else if nums[i]+b+c < 0 {
				j++
			} else {
				item := []int{nums[i], b, c}
				result = append(result, item)
				for j < z && nums[j] == nums[j+1] {
					j++
				}
				for j < z && nums[z] == nums[z-1] {
					z--
				}
				j++
				z--
			}
		}
	}
	return result
}
