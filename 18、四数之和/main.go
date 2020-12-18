package main

import (
	"fmt"
	"sort"
)

/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：
答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
*/
func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}

func fourSum(nums []int, target int) [][]int {
	var resArr [][]int
	// 从小到大排序
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		newTarget := target - nums[i]
		// 转化成了三数之和，和为新的target
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 转化为两数之和
			newTargetPro := newTarget - nums[j]
			for indexHed, indexTal := j+1, len(nums)-1; indexHed < indexTal; {
				if indexHed > j+1 && nums[indexHed] == nums[indexHed-1] {
					indexHed++
					continue
				}
				if indexTal < len(nums)-1 && nums[indexTal] == nums[indexTal+1] {
					indexTal--
					continue
				}
				tmpSum := nums[indexHed] + nums[indexTal]
				if tmpSum == newTargetPro {
					resArr = append(resArr, []int{nums[i], nums[j], nums[indexHed], nums[indexTal]})
					indexHed++
					indexTal--
				} else if tmpSum > newTargetPro {
					indexTal--
				} else {
					indexHed++
				}
			}
		}
	}
	return resArr
}
