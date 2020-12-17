package main

import (
	"fmt"
	"math"
	"sort"
)

/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。



示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

提示：

3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
*/
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums)-2; i++ {
		m, n := i+1, len(nums)-1
		for m < n {
			sums := nums[i] + nums[m] + nums[n]
			if math.Abs(float64(sums-target)) < math.Abs(float64(result-target)) {
				result = sums
			}
			if sums > target {
				n--
			} else if sums < target {
				m++
			} else {
				return sums
			}
		}
	}
	return result
}
