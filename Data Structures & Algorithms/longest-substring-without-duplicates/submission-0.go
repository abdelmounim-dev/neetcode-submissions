func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	l, max := 0, 0
	j := 0
	for i := 0; i < len(s) - 1; i++ {
		if !m[s[i]] {
			m[s[i]] = true
			l++
			if l > max {
				max = l
			}
		} else {
			for ; s[j] != s[i]; j++ {
				m[s[j]] = false
				l--
			}
			j++
		}
	}
	return max
}
