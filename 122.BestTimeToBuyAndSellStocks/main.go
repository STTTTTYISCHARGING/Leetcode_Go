/*
 * @Author       : STY
 * @Date         : 2022-08-03 08:42:38
 * @LastEditors  : STY
 * @LastEditTime : 2022-08-03 10:58:30
 * @FilePath     : \12X.BestTimeToBuyAndSellStocks\main.go
 * @Description  :
 * Copyright 2022 OBKoro1, All Rights Reserved.
 * 2022-08-03 08:42:38
 */
package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(prices)
	fmt.Println(maxProfit(prices))
	fmt.Println(maxProfit2(prices))
	fmt.Println(maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}

// leetcode 121，只能买其中一支股票
func maxProfit(prices []int) int {

	// 1.确定dp数组含义
	have := make([]int, len(prices))    //当前持有股票的情况下最大利益
	nothave := make([]int, len(prices)) //当前不持有股票情况下最大利益

	// 3. dp数组初始化
	have[0] = -prices[0]
	nothave[0] = 0

	// 4. 确定遍历顺序
	for i := 1; i < len(prices); i++ {

		// 2. 确定递推公式
		have[i] = max(have[i-1], -prices[i])
		nothave[i] = max(nothave[i-1], have[i-1]+prices[i])
	}

	return nothave[len(prices)-1]

}

// leetcode 122，只能持有一支，但是这段时间可以多次购买
func maxProfit2(prices []int) int {

	have := make([]int, len(prices))
	nothave := make([]int, len(prices))

	have[0] = -prices[0]
	nothave[0] = 0

	for i := 1; i < len(prices); i++ {
		have[i] = max(have[i-1], nothave[i-1]-prices[i])
		nothave[i] = max(nothave[i-1], have[i-1]+prices[i])
	}

	return nothave[len(prices)-1]
}

// leetcode 123. 只能持有一支，但是这段时间至多购买两次(0, 1, 2)
func maxProfit3(prices []int) int {
	/*
		1. 需要重新设计dp数组以满足这个题目要求的状态
			第i天的状态：
				buy1[i]:只进行了一次买
				sell1[i]:完成了一笔交易
				buy2[i]:在完成一比交易的前提下，进行了第二次购买
				sell2[i]:完成了两笔交易
				由于初始状态是0，可以不做记录
				剩下的四种状态，用四个数组表示其当前状态的最大利润
		2. 确定递推公式
				buy1[i] = -prices[i]
				sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i])
				buy2[i] = max(buy2[i-1], sell1[i-1]-prices[i])
				sell2[i] = max(buy2[i-1]+prices[i], sell2[i-1])
		3. 初始化
				隐含条件: 无论题目中是否允许「在同一天买入并且卖出」这一操作, 最终的答案都不会受到影响, 这是因为这一操作带来的收益为零, 所以为了方便初始化, 这里默认是可以的.
				于是就有了下面的初始化过程:
				buy1[0] := -prices[0]
				sell1[0] := 0
				buy2[0] := -prices[0]
				sell2[0] := 0

				如果题目不允许的话, 那初始化就麻烦一些, 比如sell1[i]就只能在第二天进行sell1[1]的初始化, 同理: buy2[i]在第三天进行buy2[2]的初始化, sell2[i]在第四天进行sell2[3]的初始化, 这就是能否「在同一天买入并且卖出」的不同之处.
		4. 关于最终返回结果是哪个
				在动态规划结束后,由于我们可以进行不超过两笔交易,因此最终的答案在0,sell1[i],sell2[i]中, 且为三者中的最大值
				由于在边界条件中sell1[i],sell2[i]的值已经为0, 并且在状态转移的过程中我们维护的是最大值, 因此sell1[i],sell2[i]最终一定大于等于0.

				如果最优的情况对应的是恰好一笔交易, 那么它也会因为我们在转移时允许在同一天买入并且卖出这一宽松的条件, 从sell1[i]转移至sell2[i], 可以理解假如最优是 sell1[i] 的话, 可以当天再做一次买卖, 也就是完成两次交易, 所以无论如何都可以是 sell2[i] 最优.
		5. 关于空间优化
				代码在空间优化上和我们之前说过的直接将数组用变量来替换就可以了, 但是它在理解上还是有一定难度的.
				例如在计算sell1[i]时, 我们直接使用buy1[i]而不是buy1[i-1]进行转移。buy1[i]比 buy[i-1]多考虑的是在第i天买入股票的情况, 而转移到sell1[i]时, 考虑的是在第i天卖出股票的情况, 这样在同一天买入并且卖出收益为零, 不会对答案产生影响, 所以我们才能这样进行优化.
	*/
	n := len(prices)
	buy1, buy2 := -prices[0], -prices[0]
	sell1, sell2 := 0, 0

	for i := 1; i < n; i++ {
		// buy1[i] = -prices[i]
		// sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i])
		// buy2[i] = max(buy2[i-1], sell1[i-1]-prices[i])
		// sell2[i] = max(buy2[i-1]+prices[i], sell2[i-1])
		buy1 = max(buy1, -prices[i])
		sell1 = max(sell1, buy1+prices[i])
		buy2 = max(buy2, sell1-prices[i])
		sell2 = max(sell2, buy2+prices[i])
	}
	return sell2

}

// leetcode 124. 只能持有一支，但是这段时间至多可以购买k次
func maxProfit4(prices []int) int {
	/*
		1. dp数组的定义
		首先要确定一共有多少个状态。前面几道题目的训练, 我们可以确定的状态有: 当天天数i, 当天是否持有股票(have not). 但本题的特殊之处在于: 我们必须还要确定当前进行了几笔交易j, 因为前面的几道题中, k要么是1, 要么是正无穷, 所以我们也不需要将它单独作为一个状态, 但本题就不同了, k可以取任意的值, 所以它也必须作为一个状态.

		所以本题必须用3个状态才能完整表示dp数组, 也就是三维dp数组, 按照我们之前说过的方法, 可以通过命名来优化一个状态, 也就是have表示持有状态, no表示不持有状态(我觉得这样的命名比官方的buy, sell更明确)。

		所以最终的dp数组就是二维dp数组:
			have[i][j]表示进行恰好 j 笔交易, 并且当前手上持有一支股票, 这种情况下的最大利润.
			no[i][j] 表示进行恰好 j 笔交易, 并且当前手上不持有股票, 这种情况下的最大利润.


	*/

}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
