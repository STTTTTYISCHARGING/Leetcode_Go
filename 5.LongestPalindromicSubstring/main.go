/*
 * @Author       : STY
 * @Date         : 2022-08-04 12:42:39
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-04 16:01:20
 * @FilePath     : \5.LongestPalindromicSubstring\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-04 12:42:39
 */
/*
给你一个字符串 s，找到 s 中最长的回文子串。
输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
*/

package main

import "fmt"

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
	fmt.Println(longestPalindrome2(s))
}

//n的三次方，迭代
func longestPalindrome(s string) string {
	length := len(s)
	if length < 2 {
		return s
	}
	maxLen := 1
	begin := 0
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if j-i+1 > maxLen && validPalindromic(s, i, j) {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func validPalindromic(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

//动态规划
func longestPalindrome2(s string) string {
	// 1. 定义dp数组 dp[i][j]，表示s[i, j]是否为回文串，i，j闭区间

	// 2. 确定状态转移方程

	// 3. 初始化

	// 4. 确定遍历顺序 从后向前，从下向上
	// 5. 输出
	// 6. 对dp数组空间优化
	result := ""
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = 0
			}
			if s[i] == s[j] {
				if j-i == 1 {
					dp[i][j] = 1
				}
				if dp[i+1][j-1] == 1 {
					dp[i][j] = 1
				}
			}
			if dp[i][j] == 1 && len(result) < j-i+1 {
				result = s[i : j+1]
			}
		}
	}
	return result

}
