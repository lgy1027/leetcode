package main

import (
	"fmt"
)

/*
给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗？

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

示例 3：

输入：nums1 = [0,0], nums2 = [0,0]
输出：0.00000
*/

func main() {
	num1 := []int{2, 3, 49}
	num2 := []int{3, 5, 6, 7}
	result := findMedianSortedArrays(num1, num2)
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 获取两个数组的长度
	nums1Length := len(nums1)
	nums2Length := len(nums2)
	// 总长度
	sumLength := nums1Length + nums2Length
	halfLength := 0
	// 确定半长，因为如果总长度是偶数，那么中位数就是两个数字的和除以2，如果总长度是奇数，那么中位数就是一位数字
	if sumLength%2 != 0 {
		halfLength = (sumLength + 1) / 2
	} else {
		halfLength = sumLength / 2
	}
	// 这三个用来存储临时变量，也就是和中位数有关的一位或者两位数字
	tmpNum0, tmpNum1, tmpNum2 := 0, 0, 0
	// 题目要求时间复杂度是O(m+n)，所以使用两个下标来遍历两个数组
	nums1Index, nums2Index := 0, 0
	// 这个循环就是最后找到中位数，循环比总长度的一半多找了后面的一个数字，
	// 这样的目的是当中位数是两个数字的时候，一次性就找出这两个数字，代码看起来会简单一点
	// 最开始我循环的临界是i<halfLength，这样的结果是当中位数是两个数字的时候，后面会有很大一部分重复代码
	for i := 0; i <= halfLength; i++ {
		// 第一个条件判断是当两个数组都还没有越界的时候，所以需要取两个数组中较小的那个数字
		if nums1Length > nums1Index && nums2Length > nums2Index {
			// 取了某个数组中的数字，那个数组的下标就向后加1
			// 使用tmpNum0保存临时值
			if nums1[nums1Index] <= nums2[nums2Index] {
				tmpNum0 = nums1[nums1Index]
				nums1Index++
			} else {
				tmpNum0 = nums2[nums2Index]
				nums2Index++
			}
		} else if nums1Length <= nums1Index && nums2Length > nums2Index {
			// 当数组nums1已经发生了越界，那么此时只剩nums2，所以取nums2的数字
			tmpNum0 = nums2[nums2Index]
			nums2Index++
		} else if nums1Length > nums1Index && nums2Length <= nums2Index {
			// 当数组nums2已经发生了越界，那么此时只剩nums1，所以取nums1的数字
			tmpNum0 = nums1[nums1Index]
			nums1Index++
		}
		// 之前取出的数字保存在临时变量tmpNum0中
		// 之前的值一直保存在tmpNum1中，tmpNum1一直在更新，直到更新到最中间的数字，或者中间的两个数字中较小的那个
		// 只有最后一个数字，也就超过半数的那个数字保存tmpNum2中
		if i < halfLength {
			tmpNum1 = tmpNum0
		} else {
			tmpNum2 = tmpNum0
		}
	}
	if sumLength%2 != 0 {
		// 此时中位数就只有一个
		return float64(tmpNum1)
	} else {
		// 此时中位数是中间的两个数字的和除以2
		return float64(float64((tmpNum1 + tmpNum2)) / 2)
	}
}
