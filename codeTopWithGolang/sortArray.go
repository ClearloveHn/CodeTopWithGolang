package main

import "math/rand"

/**
题目描述：
给你一个整数数组 nums，请你将该数组升序排列。

示例1：
输入：nums = [5,2,3,1]
输出：输出：[1,2,3,5]
*/

/**
解题思路：
1.选择一个基准元素（pivot）
2.将数组分区，使得基准元素左边的所有元素都小于等于它，右边的所有元素都大于等于它
3.递归地对左右两个子数组进行相同的操作
*/

// sortArray 是主要的排序函数
func sortArray(nums []int) []int {
	quick(nums, 0, len(nums)-1) // 调用快速排序函数
	return nums
}

// quick 函数实现快速排序的递归过程
func quick(arr []int, i, j int) {
	// 如果左边界大于等于右边界,说明已经排序完成,直接返回
	if i >= j {
		return
	}

	// 对数组进行分区,返回基准元素的最终位置
	mid := partition(arr, i, j)

	quick(arr, i, mid-1) // 递归排序左半部分
	quick(arr, mid+1, j) // 递归排序右半部分
}

// partition 函数实现数组的分区操作
func partition(arr []int, i int, j int) int {
	// 随机选择一个元素作为基准,这有助于避免最坏情况的时间复杂度
	p := rand.Intn(j-i+1) + i
	nums := arr

	// 将基准元素交换到数组的开始位置
	nums[i], nums[p] = nums[p], nums[i]

	for i < j {
		// 从右向左找第一个小于 nums[i] 的元素
		for nums[i] < nums[j] && i < j {
			j--
		}

		// 将找到的元素交换到左侧
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}

		// 从左向右找第一个大于等于 nums[j] 的元素
		for nums[i] < nums[j] && i < j {
			i++
		}

		// 将找到的元素交换到右侧
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}

	return i
}
