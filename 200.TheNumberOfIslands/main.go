/*
 * @Author       : STY
 * @Date         : 2022-08-05 07:40:25
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-05 08:17:15
 * @FilePath     : \200.TheNumberOfIslands\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-05 07:40:25
 */
package main

import "fmt"

func main() {
	grid := [][]byte{{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}
	// for _, v := range grid {
	// 	fmt.Println(v)
	// }
	fmt.Println(numIslands(grid))

}

//计算岛屿数量
func numIslands(grid [][]byte) int {
	nums := 0

	//遍历，先找到一个1， 作为根
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == '1' {
				nums++
				dfs(grid, i, j)
			}
		}
	}
	return nums
}

//参考树的dfs
func dfs(grid [][]byte, row, col int) {
	// 1. 边界处理
	if !IsInArea(grid, row, col) {
		return
	}

	// 2. 处理当前节点
	if grid[row][col] != '1' {
		return
	}
	grid[row][col] = '2' //标记为已经浏览过

	// 3. 执行递归操作，四个方向
	dfs(grid, row-1, col)
	dfs(grid, row+1, col)
	dfs(grid, row, col-1)
	dfs(grid, row, col+1)

}

//是否在区域内
func IsInArea(grid [][]byte, row, col int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid)
}
