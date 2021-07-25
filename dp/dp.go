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
