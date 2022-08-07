/*
 * @Author       : STY
 * @Date         : 2022-08-04 10:44:26
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-04 10:55:40
 * @FilePath     : \992.atMostK\main.go
 * @Description  : https://leetcode.cn/problems/subarrays-with-k-different-integers/solution/992-pu-gua-xing-mo-ban-miao-sha-exactlyk-i5sl/
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-04 10:44:26
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 1, 2, 3}
	k := 2
	fmt.Println(subarraysWithKDistinct(nums, k))
}

func subarraysWithKDistinct(nums []int, k int) int {
	return atMost(nums, k) - atMost(nums, k-1)
}

func atMost(nums []int, k int) int {

	// 1. 边界处理
	length := len(nums)
	if length == 0 {
		return 0
	}

	// 2. 定义数据结构
	window := make(map[int]int, length) // k:nums[i]，v:counts
	res, cnt := 0, 0                    //res结果集，cnt计算当前窗口中不同元素个数
	left, right := 0, 0

	// 3. 执行过程
	for right < length {

		//如果map中没有，cnt+1
		if window[nums[right]] == 0 {
			cnt += 1
		}
		window[nums[right]] += 1

		//修改对应数据，左移滑动窗口
		for cnt > k {
			window[nums[left]]--
			if window[nums[left]] == 0 {
				cnt -= 1 //减到0，才改cnt
			}
			left += 1
		}

		//这一步，需要理解一下
		res += right - left + 1 //例如 [2,4,3,5] 满足条件，要累加的其实就是 [2,4,3,5], [4,3,5], [3,5], [5] 四个子区间。

		right++
	}

	return res
}
