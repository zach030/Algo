package window

// 无重复字符最长子串
func lengthOfLongestSubstring(s string) int {
	size := len(s)
	max := 0
	start, end := 0, 0
	filter := make(map[string]int, 0)
	for end < size {
		if nextIndex, ok := filter[string(s[end])]; ok {
			if nextIndex > start {
				start = nextIndex
			}
		}
		if end-start+1 > max {
			max = end - start + 1
		}
		filter[string(s[end])] = end + 1
		end++
	}
	return max
}
