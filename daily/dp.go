package daily

import "math"

// 斐波拉契数列，简单的动态规划
// f(n)=f(n-1)+f(n-2)，因为只需要前两个数相加，可以只保留两个变量
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b := 0, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}
	return sum
}

// 跳上n个台阶有f(n)种方法，最后一步只有两种情况：跳一步or跳两步
// 跳一步的情况：之前已经跳了n-1步，因此有f(n-1)方案
// 跳两步：之前跳了n-2步，有f(n-2)方案
// f(n)=f(n-1)+f(n-2), f(0)=1,f(1)=1,f(2)=2
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	a, b := 1, 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = (a + b) % 1000000007
		a = b
		b = sum
	}
	return sum
}

// 买卖股票最大利润，dp[i]表示前i天的最大利润
// dp[i]=max(dp[i-1],prices[i]-lowest)
// 第i天有两种选择：不卖出，则利润仍是dp[i-1]；卖出：当天的价格-前i天最低价
// 因此在循环时需要记录前i天的最低价
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// [7,1,5,3,6,4]
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices)+1)
	dp[0], dp[1] = 0, 0
	profit, low := 0, math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < low {
			low = prices[i]
		}
		dp[i+1] = max(dp[i], prices[i]-low)
		if dp[i+1] > profit {
			profit = dp[i+1]
		}
	}
	return profit
}
