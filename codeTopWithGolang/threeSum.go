package main

import (
	"fmt"
	"sort"
)

/**
题目描述：
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。

输入：nums = [-1,0,1,2,-1,4]
输出：[[-1,-1,2],[-1,0,1]
*/

/**
解题思路：
1.首先对数组进行排序，这样可以方便去重和使用双指针。
2.遍历排序后的数组，固定第一个数nums[i]
3.使用双指针left和right，分别指向i+1和数组末尾。
4.计算三数之和sum = nums[i] + nums[left] + nums[right]
5.如果sum == 0，我们找到了一组解，将其加入结果集，然后移动左右指针。
6.如果 sum < 0,说明和太小,left 右移。如果 sum > 0,说明和太大,right 左移。
7.注意去重:对于第一个数和左右指针,如果和前一个数相同,直接跳过,避免重复。
*/

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		// 跳过重复的第一个数
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的左指针
				for left < right && nums[left] == nums[left+1] {
					left++
				}

				// 跳过重复的右指针
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
