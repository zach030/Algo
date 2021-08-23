package dp

func strStr(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		var j int
		for j = 0; j < n; j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == n {
			return i
		}
	}
	return -1
}

func kmp(s, p string) int {
	if p == "" {
		return 0
	}
	n, m := len(s), len(p)
	s, p = " "+s, " "+p
	next := make([]int, m+1)
	i := 2
	for j := 0; i <= m; i++ {
		for j > 0 && p[i] != p[j+1] {
			j = next[j]
		}
		if p[i] == p[j+1] {
			j++
		}
		next[i] = j
	}
	j := 0
	for i := 1; i <= n; i++ {
		for j > 0 && s[i] != p[j+1] {
			j = next[j]
		}
		if s[i] == p[j+1] {
			j++
		}
		if j == m {
			return i - m
		}
	}
	return -1
}

