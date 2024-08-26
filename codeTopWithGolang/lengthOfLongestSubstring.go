package main

import "fmt"

/**
题目描述：
给定一个字符串S，请你找出其中不含有重复字符的最长子串的长度。

示例1：
输入：S = “abcabcbb”。
输出：3。
解释：因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

/**
解题思路:
1.例如，求连续子串，连续子数组，连续子序列，可以考虑使用滑动窗口。
2.初始化，设置窗口左边界为0，无重复子串的最大长度为0。
3.遍历字符串，检查当前的字符是否重复，如果重复，检查是否在窗口内，如果在，移动窗口左边界到上次那个位置的下一位置。更新当前字符的位置。
计算窗口长度。注意窗口长度的计算里面的下标问题(从0开始)。
4.使用map来辅助滑动窗口的条件检查，比如，检查字符是否重复，是否在窗口里面，确定窗口的左边界位置。
*/

func lengthOfLongestSubstring(s string) int {

	// 创建一个map来存储每个字符最后出现的索引。
	charIndex := make(map[rune]int)

	// 用于存储最长无重复子串的长度。
	maxLength := 0

	// 当前无重复子串的起始索引，即滑动窗口的左边界。
	start := 0

	// 遍历字符串,i为滑动窗口的右边界。
	for i, char := range s {

		// 检查当前字符是否在map中
		if lastIndex, ok := charIndex[char]; ok {
			// 如果在map中，查找它上次出现的位置是否在窗口内
			if lastIndex >= start {
				// 如果在窗口内，窗口左边界移动到其上次出现的索引位置的下一个位置
				start = lastIndex + 1
			}
		}

		// 无论字符是否重复，都更新它在map中的位置为当前索引
		charIndex[char] = i

		// 计算当前无重复子串的长度
		currentLength := i - start + 1

		// 如果当前长度大于已知的最大长度，则更新最大长度
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

func main() {
	testCases := []string{"abcabcbb", "bbbbb", "pwwkew"}
	for _, s := range testCases {
		fmt.Printf("输入: %s\n输出: %d\n\n", s, lengthOfLongestSubstring(s))
	}
}
