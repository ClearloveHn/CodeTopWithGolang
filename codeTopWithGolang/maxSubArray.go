package main

import (
	"fmt"
	"math"
)

/**
题目描述：
给你一个整数数组nums，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
子数组是数组的一个连续部分。

示例1：
输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
输出：6
解释：连续子数组[4,-1,2,1]的和最大为6
*/

/**
解题思路：
1.遍历数组。
2.对于每个元素，我们有两个选择：a.将当前元素作为新子数组的开始。b.将当前元素加入到之前的子数组中。
3.选择两个选项中较大的一个，同时，持续跟踪遇到的最大和。
*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]     // 记录全局最大和
	currentSum := nums[0] // 记录到当前位置的最大和

	for i := 1; i < len(nums); i++ {
		// 在下面两种选择中取较大值：
		// 1. 仅包含当前元素 nums[i]
		// 2. 将当前元素加入到之前的子数组中 (currentSum + nums[i])
		currentSum = maxNum(currentSum+nums[i], nums[i])

		// 更新全局最大和
		maxSum = maxNum(maxSum, currentSum)
	}

	return maxSum

}

func maxNum(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func main() {
	// 示例输入数组
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	// 调用 maxSubArray 函数计算最大子数组和
	result := maxSubArray(nums)

	// 打印结果
	fmt.Printf("最大子数组和: %d\n", result)
}
