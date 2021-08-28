package dp

import (
	"fmt"
	"math"
)

// 最长回文子串
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	size := len(s)
	left, right, length := 0, 0, 1
	maxStart, maxLen := 0, 0
	for i := 0; i < size; i++ {
		left = i - 1
		right = i + 1
		for left > 0 && s[left] == s[i] {
			length++
			left--
		}
		for right < size && s[right] == s[i] {
			length++
			right++
		}
		for left >= 0 && right < size && s[right] == s[left] {
			length += 2
			left--
			right++
		}
		if length > maxLen {
			maxLen = length
			maxStart = left
		}
		length = 1
	}
	return s[maxStart+1 : maxStart+maxLen+1]
}

// 打家劫舍 198
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	size := len(nums)
	dp := make([]int, size+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= size; i++ {
		dp[i] = max(dp[i-1], nums[i-1]+dp[i-2])
	}
	return dp[size]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 斐波拉契数列
func fib(n int) int64 {
	if n < 1 {
		return 0
	}
	var memo = make(map[int]int64)
	res := helper(memo, n)
	return res
}

func helper(arr map[int]int64, n int) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if arr[n] != 0 {
		return arr[n]
	}
	arr[n] = helper(arr, n-1)%1000000007 + helper(arr, n-2)%1000000007
	return arr[n]
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		}
		return min(dp(i-1, j)+1, dp(i, j-1)+1, dp(i-1, j-1)+1)
	}
	return dp(len(word1)-1, len(word2)-1)
}

func min(a, b, c int) int {
	m := a
	if b <= m {
		m = b
	}
	if c <= m {
		m = c
	}
	return m
}

func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	prev, curr := 1, 1
	for i := 3; i < n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	memo := make(map[int]int, 0)
	var dp func(n int) int
	dp = func(n int) int {
		if n == 0 {
			return 0
		}
		if n < 0 {
			return -1
		}
		if ans, ok := memo[n]; ok {
			return ans
		}
		res := math.MaxInt16
		for _, coin := range coins {
			sub := dp(n - coin)
			if sub == -1 {
				continue
			}
			if sub+1 < res {
				res = sub + 1
			}
		}
		if res != math.MaxInt16 {
			memo[n] = res
		} else {
			memo[n] = -1
		}
		return memo[n]
	}
	return dp(amount)
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	m := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > m {
			m = dp[i]
		}
	}
	return m
}

// 四键键盘
func maxA(n int) int {
	var dp func(n, num, copy int) int
	dp = func(n, num, copy int) int {
		if n <= 0 {
			return num
		}
		return max3(dp(n-1, num+1, num), dp(n-1, num+copy, copy), dp(n-2, num, num))
	}
	return dp(n, 0, 0)
}

func max3(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

// 650 ca+cc cv
func minSteps(n int) int {
	//var dp func(num) int
	//dp = func(num) int {
	//	if num == n {
	//		return op
	//	}
	//	return min2(dp(num), dp(num))
	//}
	//return dp(1)
	return 0
}

func min2(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 70
func climbStairs(n int) int {
	dp := make([]int, n+1)
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 740
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = min2(dp[i-1], dp[i-2]) + cost[i]
	}
	fmt.Println(dp)
	return min2(dp[len(cost)-1], dp[len(cost)-2])
}

// 62
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}

// 63 障碍物为1 空为0
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 343 整数拆分
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

// 704
func deleteAndEarn(nums []int) int {
	maxVal := 0
	for _, val := range nums {
		maxVal = max(maxVal, val)
	}
	sum := make([]int, maxVal+1)
	for _, val := range nums {
		sum[val] += val
	}
	return rob(sum)
}

// 55 跳跃游戏
func canJump(nums []int) bool {
	size := len(nums)
	end := size - 1
	for i := size - 2; i >= 0; i-- {
		if end-i <= nums[i] {
			end = i
		}
	}
	return end == 0
}

// 最少次数跳到终点
func jump(nums []int) int {
	//dp[i]:跳跃到i用的最少次数
	if len(nums) == 1 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums) && j <= i+nums[i]; j++ {
			dp[j] = min2(dp[i]+1, dp[j])
		}
	}
	return dp[len(dp)-1]
}

// 分割等和子集
func canPartition(nums []int) bool {
	var sum int
	size := len(nums)
	for _, i2 := range nums {
		sum+=i2
	}
	if sum%2!=0{
		return false
	}
	sum = sum / 2
	// dp[i][j] 选择前i个数字，容量为j的包，可以满足
	dp:=make([][]bool,size+1)
	for i := 0; i <= size; i++ {
		dp[i] = make([]bool, size+1)
	}
	for i := 0; i <= size; i++ {
		dp[i][0]=true
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1]<0{
				dp[i][j]=dp[i-1][j]
			}else{
				dp[i][j]=dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[size][sum]
}

// 最大子序和
func maxSubArray(nums []int) int {
	return 0
}