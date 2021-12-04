package daily

// 把字符串 s 中的每个空格替换成"%20"
func replaceSpace(s string) string {
	o := []byte(s)
	n := make([]byte, 0, len(o))
	for i := 0; i < len(o); i++ {
		if o[i] == ' ' {
			n = append(n, []byte("%20")...)
		} else {
			n = append(n, o[i])
		}
	}
	return string(n)
}

func reverseLeftWords(s string, n int) string {
	buf := []byte(s)
	buf = append(buf, buf[:n]...)
	str := string(buf[n:])
	return str
}
