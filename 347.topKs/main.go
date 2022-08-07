/*
 * @Author       : STY
 * @Date         : 2022-08-02 08:11:07
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-02 19:02:13
 * @FilePath     : \1\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-02 08:11:07
 */
package main

import "fmt"

func main() {
	nums1 := []int{12, 323, 54, 12, 23, 234, 44, 55, 66, 32, 1}
	nums2 := []int{1, 1, 2, 2, 3, 4, 4, 4, 4, 2, 1, 4, 5, 7, 8, 67, 4, 5, 7, 9, 3, 2, 3, 1, 2}
	heapSort(nums1)
	res := topKFrequent(nums2, 2)
	res2 := topKHeapSort(nums2, 2)
	fmt.Println(nums1)
	fmt.Println(nums2)
	fmt.Println(res)
	fmt.Println(res2)
}

func heapSort(nums []int) {

	lens := len(nums) - 1

	// 建立初始堆
	// lens/2是最后一个非叶子节点的位置，0是第一个非叶子节点
	for i := lens / 2; i >= 0; i-- {
		down(nums, i, lens)
	}

	// 执行堆排序
	// 堆顶元素与堆尾元素交换位置，然后重新建堆
	for i := lens; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		lens--              //堆元素-1
		down(nums, 0, lens) //重新建堆
	}

}

// 堆的下沉，递归
func down(nums []int, start, lens int) {
	min := start
	if start<<1+1 <= lens && nums[start<<1+1] < nums[min] {
		min = start<<1 + 1
	}

	if start<<1+2 <= lens && nums[start<<1+2] < nums[min] {
		min = start<<1 + 2
	}

	if min != start {
		nums[min], nums[start] = nums[start], nums[min]
		down(nums, min, lens)
	}
}

func topKFrequent(nums []int, k int) []int {

	// 1.构建初始数据结构
	//make(map[keyType] valueType, cap)，cap表示初始存储能力
	cnts := make(map[int]int, k)
	for i := range nums {
		cnts[nums[i]]++
	}
	//构建元素出现次数的数组,方便进行堆操作
	//make([]type, len, cap)，其中，type表示数组元素类型，len表示长度，cap表示容量
	heapCnt := make([][2]int, 0, len(cnts))
	for num, cnt := range cnts {
		heapCnt = append(heapCnt, [2]int{num, cnt})
	}
	// fmt.Printf("cnts=%+v,heapCnt=%+v\n", cnts, heapCnt)

	// 2. 执行堆算法
	// 1. 堆的下沉操作，迭代not递归
	/*
			1. 自顶向下调整堆小顶堆
			2. 小顶堆的堆顶为出现次数第k多的元素,堆里面的元素出现次数都>=堆顶元素出现的次数
		            0
		        1       2
		     3     4  5   6
	*/
	heapify := func(heapCnt [][2]int, start, count int) {
		for start < count {
			pos := start
			left, right := 2*pos+1, 2*pos+2
			if left < count && heapCnt[left][1] < heapCnt[pos][1] {
				pos = left
			}
			if right < count && heapCnt[right][1] < heapCnt[pos][1] {
				pos = right
			}
			if pos == start {
				break
			}
			heapCnt[pos], heapCnt[start] = heapCnt[start], heapCnt[pos]
			start = pos
		}
	}

	// 2. k大小的堆，初始化
	//堆的大小为k,从最后一个非叶子节点开始堆化数组
	for p := (k >> 1) - 1; p >= 0; p-- {
		heapify(heapCnt, p, k)
	}

	// 3. 从第k+1个元素，到第n个元素，一次比较入堆
	for i := k; i < len(heapCnt); i++ {
		/*
			和堆顶元素比较,如果出现次数大于堆顶元素,就和堆顶元素交换,
			相当于删除堆顶元素,然后堆顶元素下沉
		*/
		if heapCnt[i][1] > heapCnt[0][1] {
			heapCnt[i], heapCnt[0] = heapCnt[0], heapCnt[i]
			heapify(heapCnt, 0, k)
		}
	}

	ret := make([]int, 0, k)
	//堆中元素就是要求的结果
	for i := 0; i < k; i++ {
		ret = append(ret, heapCnt[i][0])
	}

	return ret
}

func topKHeapSort(nums []int, k int) []int {
	// 1. 定义好初始的数据结构
	counts := make(map[int]int, k)
	for _, v := range nums {
		counts[v]++
	}
	NewNums := make([][2]int, 0, len(nums))
	for num, count := range counts {
		NewNums = append(NewNums, [2]int{num, count})
	}
	fmt.Printf("counts=%+v, NewNums=%+v\n", counts, NewNums)

	// 2. 初始化数组
	for i := (k - 1) / 2; i >= 0; i-- {
		down3(NewNums, i, k)
	}

	// 3. 遍历k+1 to n
	for i := k; i < len(NewNums); i++ {
		if NewNums[i][1] > NewNums[0][1] {
			NewNums[i], NewNums[0] = NewNums[0], NewNums[i]
			down2(NewNums, 0, k)
		}
	}

	// 4. 取数组
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, NewNums[i][0])
	}
	return res
}

// logk 迭代写法
func down2(NewNums [][2]int, start, count int) {
	for start < count {
		pos := start
		left, right := 2*pos+1, 2*pos+2
		if left < count && NewNums[left][1] < NewNums[pos][1] {
			pos = left
		}
		if right < count && NewNums[right][1] < NewNums[pos][1] {
			pos = right
		}
		if pos == start {
			break
		}
		NewNums[pos], NewNums[start] = NewNums[start], NewNums[pos]
		start = pos
	}
}

// 递归写法
func down3(heapNums [][2]int, start, end int) {
	pos := start
	left, right := pos*2+1, pos*2+2
	if left < end && heapNums[left][1] < heapNums[pos][1] {
		pos = left
	}
	if right < end && heapNums[right][1] < heapNums[pos][1] {
		pos = right
	}
	if pos != start {
		heapNums[pos], heapNums[start] = heapNums[start], heapNums[pos]
		down3(heapNums, pos, end)
	}

}
