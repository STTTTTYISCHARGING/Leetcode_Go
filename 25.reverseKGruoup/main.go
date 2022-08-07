/*
 * @Author       : STY
 * @Date         : 2022-08-02 20:41:30
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-02 21:21:34
 * @FilePath     : \25\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-02 20:41:30
 */
package main

import "fmt"

type Node struct {
	// 值
	Data interface{}
	// 后继节点指针
	Next *Node
}

// 链表是否为空
func IsEmpty(node *Node) bool {
	return node == nil
}

// 是否是最后一个节点
func IsLast(node *Node) bool {
	return node.Next == nil
}

// 查找前驱节点
func FindPrevious(index int, list *Node) *Node {
	temp := list
	if index == 0 {
		return temp
	}
	for i := 0; i < index-1; i++ {
		temp = temp.Next
	}
	return temp
}

// 插入节点
func InsertList(index int, list *Node, data int) {
	temp := list
	if index == 0 {
		temp.Data = data
		return
	}
	if index < 0 {
		fmt.Println("索引不合法")
		return
	}
	// 查找当前索引位置
	for i := 0; i < index; i++ {
		if IsEmpty(temp) {
			fmt.Println("索引越界")
			return
		}
		temp = temp.Next
	}

	// 创建新节点
	newNode := Node{Data: data}
	// 新节点的Next指针指向当前位置
	newNode.Next = temp
	// 查找前驱节点
	pre := FindPrevious(index, list)
	// 前驱节点的Next指针指向新节点
	pre.Next = &newNode
}

// 删除节点
func DeleteNode(index int, list *Node) {
	temp := list
	if index == 0 {
		return
	}
	// 查找当前索引位置
	for i := 0; i < index; i++ {
		temp = temp.Next
	}
	// 查找前驱节点
	pre := FindPrevious(index, list)
	// 前驱节点的Next指针指向后驱节点即可删除当前
	pre.Next = temp.Next

}

// 末尾添加节点
func Append(list *Node, node *Node) {

	for list != nil {
		temp := list.Next
		if temp == nil {
			list.Next = node
			break
		}
		list = list.Next
	}
}

// 打印所有的List
func PrintList(list *Node) {
	temp := list
	for temp != nil {
		fmt.Println(temp.Data)
		temp = temp.Next
	}
}

func main() {
	// 创建头节点
	headNode := &Node{
		Data: 0,
		Next: nil,
	}
	for i := 1; i < 10; i++ {
		node := &Node{Data: i}
		Append(headNode, node)
	}
	// InsertList(1, headNode, 10)
	// DeleteNode(3, headNode)
	PrintList(headNode)
	headNode2 := reverseKGruoup1(headNode, 3)
	PrintList(headNode2)

}

//递归
func reverseKGruoup1(head *Node, k int) *Node {
	curNode := head
	for i := 0; i < k; i++ {
		if curNode == nil {
			return head
		}
		curNode = curNode.Next
	}
	res := reverse(head, curNode)
	head.Next = reverseKGruoup1(curNode, k)
	return res
}
func reverse(left, right *Node) *Node {
	pre := right
	for left != right {
		next := left.Next
		left.Next = pre
		pre, left = left, next
	}
	return pre
}
