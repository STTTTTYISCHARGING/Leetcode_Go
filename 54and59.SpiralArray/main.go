/*
 * @Author       : STY
 * @Date         : 2022-08-05 09:41:33
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-05 10:09:59
 * @FilePath     : \54and59.SpiralArray\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-05 09:41:33
 */
package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	for _, v := range matrix {
		fmt.Println(v)
	}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	counts := len(matrix[0]) * len(matrix)
	for counts != 0 {

		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
			counts--
		}
		if counts == 0 {
			break
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
			counts--
		}
		if counts == 0 {
			break
		}
		right--

		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
			counts--
		}
		if counts == 0 {
			break
		}
		bottom--

		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
			counts--
		}
		if counts == 0 {
			break
		}
		left++
	}

	return res
}
