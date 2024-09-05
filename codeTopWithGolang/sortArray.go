package main

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

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, low, high int) {
	if low < high {
		// 获取分区点
		pivotIndex := partition(nums, low, high)

		// 递归排序左边部分
		quickSort(nums, low, pivotIndex-1)

		// 递归排序右边部分
		quickSort(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	// 选择最后一个元素作为基准（pivot）
	pivot := nums[high]

	// i 表示小于等于pivot的元素的最后位置,初始化一个虚拟的。
	i := low - 1

	// 遍历low到high-1的元素
	for j := low; j < high; j++ {
		// 如果当前元素小于等于pivot
		if nums[j] <= pivot {
			i++
			// 交换元素
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	// 将基准元素放到正确的位置
	nums[i+1], nums[high] = nums[high], nums[i+1]

	return i + 1
}
