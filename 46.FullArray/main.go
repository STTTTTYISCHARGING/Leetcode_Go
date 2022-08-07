/*
 * @Author       : STY
 * @Date         : 2022-08-05 08:53:17
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-05 09:31:11
 * @FilePath     : \46.FullArray\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-05 08:53:17
 */
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute((nums)))
}

// 全排列，区别于全组合，元素有序
func permute(nums []int) [][]int {

	// 1. 准备数据结构
	res := [][]int{}

	// 2. 回溯函数，很模板化，DFS+找条件回溯
	var backTrack func(nums []int, length int, path []int)
	backTrack = func(nums []int, length int, path []int) {
		// 1. 递归终止条件
		if length == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
		}
		// 2. 单层遍历执行，树的横向遍历
		for i := 0; i < length; i++ {
			cur := nums[i]
			path = append(path, cur)
			// a = append(a, 1, 2, 3) // 追加多个元素, 手写解包方式
			// a = append(a, []int{1,2,3}...) // 追加一个切片, 切片需要解包
			nums = append(nums[:i], nums[i+1:]...)
			backTrack(nums, len(nums), path) //纵向遍历
			// 回溯过程
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	backTrack(nums, len(nums), []int{})
	return res
}
