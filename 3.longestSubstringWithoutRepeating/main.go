/*
 * @Author       : STY
 * @Date         : 2022-08-02 19:53:50
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-04 10:44:19
 * @FilePath     : \3.longestSubstringWithoutRepeating\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-02 19:53:50
 */
package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	// fmt.Println(s[2]) //99
	res := window(s)
	fmt.Println(res)
}

func window(s string) int {

	// 1. 边界判定
	length := len(s)
	if length <= 1 {
		return length
	}

	// 2. 准备数据结构
	res := 1
	left, right := 0, 0
	window := [128]int{} //默认都是0，第一次遇到，改为当前位置的下一位

	// 3. 滑动窗口遍历
	for right < length {
		rightChar := s[right]
		rightCharIndex := window[rightChar] //从string切片取出来就是int表示

		if rightCharIndex > left {
			left = rightCharIndex
		}
		res = max(res, right-left+1)

		window[rightChar] = right + 1 //当前的位置就是right，如果再次碰到，left移动到下一个位置重新计算
		right++
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
