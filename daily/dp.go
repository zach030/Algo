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

// 连续子数组最大和,f(n)表示以num[n]结尾的最大子数组和
// f(n-1)<=0,f(n)=nums[n]
// f(n-1)>0,f(n)=f(n-1)+nums[n]
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var last = nums[0]
	var curr = 0
	var maxSum = last
	for i := 1; i < len(nums); i++ {
		if last <= 0 {
			curr = nums[i]
		} else {
			curr = last + nums[i]
		}
		if curr > maxSum {
			maxSum = curr
		}
		last = curr
	}
	return maxSum
}

// 礼物的最大价值
//输入:
//[
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
//]
//输出: 12
//解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

// dp[i][j]表示从[0,0]到[i,j]拿到的最多价值礼物
// 最后一步到达dp[i][j]有两条路:[i][j-1]--->[i][j], [i-1][j]--->[i][j]，取两者的最大值即可
// dp[i][j]=max(dp[i-1][j]+grid[i][j],dp[i][j-1]+grid[i][j])
func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
				continue
			}
			if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
				continue
			}
			dp[i][j] = max(dp[i-1][j]+grid[i][j], dp[i][j-1]+grid[i][j])
		}
	}
	return dp[m-1][n-1]
}
