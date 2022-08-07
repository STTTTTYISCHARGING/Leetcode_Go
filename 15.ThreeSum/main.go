/*
 * @Author       : STY
 * @Date         : 2022-08-02 21:35:57
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-02 21:58:13
 * @FilePath     : \15.ThreeSum\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-02 21:35:57
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := [][]int{}

	// 1. 排序
	sort.Ints(nums)

	// 2. for循环先取a
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			continue
		}

		//去重a
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 3. 左右指针遍历有序数组，寻找合适的b、c
		left, right := i+1, len(nums)-1
		for left < right {
			target := nums[i] + nums[left] + nums[right]
			if target == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				//去重b
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				//去重c
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if target > 0 {
				right--
			} else {
				left++
			}
		}
	}
	fmt.Println(res)
}
