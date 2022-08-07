/*
 * @Author       : STY
 * @Date         : 2022-08-04 12:23:50
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-04 12:27:49
 * @FilePath     : \236.nearestBinaryTreeFather\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-04 12:23:50
 */
//https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solution/236-er-cha-shu-de-zui-jin-gong-gong-zu-xian-hou-xu/
package main

import "fmt"

func main() {
	fmt.Println(1)
}

/*

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归前序遍历

	// 1.终止条件
	if root == p || root == q || root == nil {
		return root
	}

	// 2. 递推遍历，需要返回值
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 3.根据返回值讨论，有四种情况
	if left != nil && right != nil { // ① 同时不为空，说明p、q在异侧
		return root
	}
	if left != nil { // ② 说明p、q都在这一侧
		return left
	}
	if right != nil {
		return right
	}
	return nil // ④ 说明左右都不包含p，q，返回nil
}

*/
