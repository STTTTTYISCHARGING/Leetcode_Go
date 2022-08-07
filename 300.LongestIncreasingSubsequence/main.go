/*
 * @Author       : STY
 * @Date         : 2022-08-07 11:02:36
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-07 11:12:22
 * @FilePath     : \300.LongestIncreasingSubsequence\main.go
 * @Description  : 最长递增子序列，动态规划
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-07 11:02:36
 */
package main

import "fmt"

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(nums)
	fmt.Println(lengthOfLIS(nums))
}

func lengthOfLIS(nums []int) int {
	res := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
