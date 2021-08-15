package dp

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
