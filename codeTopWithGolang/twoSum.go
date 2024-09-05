package main

/**
题目描述：
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
你可以按任意顺序返回答案。

示例 1：
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1]
*/

/**
解题思路：
1.创建一个map来存储遍历过的数字及其索引。key是数字,value是索引。
2.遍历数组,对每个数字num,我们做以下操作：
计算它需要的配对数complement = target - num
检查complement是否在map中
如果找到配对,直接返回两个数的索引
如果没找到,将当前数字及其索引加入map
*/

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if num, ok := numMap[complement]; ok {
			return []int{num, i}
		}
		numMap[num] = i
	}
	return nil
}
