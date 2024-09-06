package main

/**
题目描述：
给你一个字符串s,找到s中最长的回文子串（如果字符串向前和向后读都相同，则它满足 回文性。）。
示例1：
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

/**
解题思路：
1.创建一个二维布尔数组dp，dp[i][j]表示字符串s从索引i到j的子串是否为回文串。
2.初始化所有长度为1的子串为回文串（dp[i][i] = true）。
3.从长度2开始，逐步扩大子串长度，判断每个子串是否为回文串。
4.如果s[i] == s[j]，且dp[i+1][j-1]为true（或者子串长度<=3），则dp[i][j]为true。
*/

func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	// 初始化动态规划数组
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	// start 记录最长回文子串的起始索引
	// maxLen 记录最长回文子串的长度
	start, manLen := 0, 1

	// 所有长度为1的子串都是回文串
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 检查长度为2及以上的子串
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1

			// 如果首尾字符相同，且中间子串是回文串（或长度<=3），则整个子串是回文串
			if s[i] == s[j] && (length <= 3 || dp[i+1][j-1]) {
				dp[i][j] = true
				// 如果找到更长的回文子串，更新start和maxLen
				if length > manLen {
					start = i
					manLen = length
				}
			}
		}
	}

	// 返回最长回文子串
	return s[start : start+manLen]
}
