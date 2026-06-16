func characterReplacement(s string, k int) int {
	replaced := make(map[int]bool)
	max := 0
	start := 0
	for end:= 0; end < len(s); {
		if s[start] == s[end] {
			if end-start +1 > max {
				max = end-start +1
			}
		} else if k > 0 {
			k--
			replaced[end] = true
			if end-start > max {
				max = end-start +1
			}
		} else {
			for !replaced[start] {
				start++
			}
			k++
			replaced[start] = false
			continue
		}
		end++
	}
	return max
}
