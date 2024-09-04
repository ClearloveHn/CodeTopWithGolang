package main

import "fmt"

/**
题目描述：
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

示例1：
输入: [3,2,1,5,6,4], k = 2
输出: 5

"数组中的第K个最大元素"这个问题要求我们在一个未排序的数组中找到第K大的元素。
例如，在数组 [3,2,1,5,6,4] 中，第1大的元素是6，第2大的元素是5，第3大的元素是4，以此类推。
*/

/**
解题思路：
1.基于快排的思想：每次选择一个枢纽元素，将数组分为小于枢纽的部分和大于枢纽的部分。然后根据枢纽的位置，可以确定第k个最大的元素在哪个半区。
2.我们的目标是找第 k 大的元素。但是快速选择算法是设计用来找第 x 小的元素。所以我们需要在这之间建立一个关系。
3.通用公式第 k 大的元素 = 第 (n-k+1) 小的元素，由于数组索引从 0 开始，我们需要再减 1,第 k 大的元素的索引 = (n-k+1) - 1 = n-k
*/

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

// quickSelect 快速选择算法的核心实现
// l,r是区间节点，搜索范围，k为寻找的元素索引
func quickSelect(nums []int, l, r, k int) int {
	// 如果区间只有一个元素，直接返回
	if l == r {
		return nums[k]
	}

	// 选择第一个元素作为分区点（pivot）
	partition := nums[l]

	// i和j分别从区间的左右两端向中间移动，用于分区操作。
	i := l - 1
	j := r + 1

	for i < j {
		// 向右移动i，直到找到一个大于等于partition的元素
		for i++; nums[i] < partition; i++ {
		}

		// 向左移动j，直到找到一个小于等于partition的元素
		for j--; nums[j] > partition; j-- {
		}

		// 如果i和j没有相遇，交换这两个元素
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 根据k的位置决定向左还是向右递归
	if k <= j {
		// k在左半部分，继续在左半部分寻找
		return quickSelect(nums, l, j, k)
	} else {
		// k在右半部分，继续在右半部分寻找
		return quickSelect(nums, j+1, r, k)
	}
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}
