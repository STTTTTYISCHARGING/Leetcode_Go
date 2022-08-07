/*
 * @Author       : STY
 * @Date         : 2022-08-07 09:42:50
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-07 10:59:58
 * @FilePath     : \42.CatchRain\main.go
 * @Description  : 经典问题，接雨水，用单调递减栈。 图解:https://leetcode.cn/problems/trapping-rain-water/solution/yi-miao-jiu-neng-du-dong-de-dong-hua-jie-o9sv/
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-07 09:42:50
 */
package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(height))
	fmt.Println(trap2(height))
}

func trap(height []int) int {
	// 1. 准备数据结构
	sum := 0
	stack := make([]int, len(height))
	stack = append(stack, height[0]) //栈中存的是height的位置下标

	// 2. 遍历height，用单调栈存储单调递减栈
	for i := 1; i < len(height); i++ {
		if len(stack) != 0 && height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if len(stack) != 0 && height[i] == height[stack[len(stack)-1]] {
			stack[len(stack)-1] = i
		} else {
			for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					left := stack[len(stack)-1]
					high := min(height[left], height[i]) - height[mid]
					width := i - left - 1
					sum += high * width
				}
			}
			stack = append(stack, i)
		}
	}

	return sum
}

func trap2(height []int) int {
	res := 0
	stack := make([]int, len(height))
	stack = append(stack, 0)

	for i := 0; i < len(height); i++ {
		if len(stack) != 0 && height[i] < height[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if height[i] == height[stack[len(stack)-1]] {
			stack[len(stack)-1] = i
		} else {
			for len(stack) != 0 && height[i] > height[stack[len(stack)-1]] {
				mid := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					left := stack[len(stack)-1]
					high := min(height[i], height[left]) - height[mid]
					width := i - left - 1
					res += high * width
				}
			}
		}
		stack = append(stack, i)
	}
	return res
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
