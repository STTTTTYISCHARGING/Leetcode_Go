/*
 * @Author       : STY
 * @Date         : 2022-08-02 19:03:19
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-02 19:36:47
 * @FilePath     : \topK\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-02 19:03:19
 */
package main

import (
	"fmt"
)

/*
[3,2,3,1,2,4,5,5,6]
4
*/
func main() {
	nums := []int{3, 2, 1, 5, 6, 4}

	//第k大元素
	k := 2
	target := len(nums) - k
	QuickSort(nums, 0, len(nums)-1, target)
	fmt.Println(nums[target])
	fmt.Println(nums)
}

func QuickSort(nums []int, left, right, k int) {
	if left > right {
		return
	}
	cur := parition(nums, left, right)
	if cur == k {
		return
	} else if cur > k {
		QuickSort(nums, left, cur-1, k)
	} else {
		QuickSort(nums, cur+1, right, k)
	}
}
func parition(nums []int, left, right int) int {
	if left > right {
		return -1
	}
	pos := nums[left]
	for left < right {
		for left < right && nums[right] >= pos {
			right--
		}
		nums[right], nums[left] = nums[left], nums[right]
		for left < right && nums[left] < pos {
			left++
		}
		nums[right], nums[left] = nums[left], nums[right]
	}
	return left
}
