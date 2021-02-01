package main

import "fmt"

/*
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须 原地 修改，只允许使用额外常数空间。

 TODO 参考全排列

示例 1：

输入：nums = [1,2,3]
输出：[1,3,2]
示例 2：

输入：nums = [3,2,1]
输出：[1,2,3]
示例 3：

输入：nums = [1,1,5]
输出：[1,5,1]
示例 4：

输入：nums = [1]
输出：[1]
*/
func main() {
	array := []int{5, 4, 3, 2, 1}

	k := len(array) - 1
	i := k
	for i >= 0 {
		if i < k && array[i] < array[i+1] {
			j := i
			for j <= k && array[i] <= array[j] {
				j++
			}
			array[i], array[j-1] = array[j-1], array[i]
			break
		}
		i--
	}
	reverse(array, i+1, k)
	fmt.Println(array)
}

func reverse(array []int, begin int, end int) {
	for begin < end {
		array[begin], array[end] = array[end], array[begin]
		end--
		begin++
	}
}
