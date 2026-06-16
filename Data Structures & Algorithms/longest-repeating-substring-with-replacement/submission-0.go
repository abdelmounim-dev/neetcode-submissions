func characterReplacement(s string, k int) int {
	replaced := make(map[int]bool)
	max,l  := 0, 0
	start := 0
	for end:= 0; end < len(s); {
		if s[start] == s[end] {
			l++ 
			if l > max {
				max = l
			}
		} else if k > 0 {
			k--
			replaced[end] = true
			l++ 
			if l > max {
				max = l
			}
		} else {
			for !replaced[start] {
				start++
				l--
			}
			k++
			replaced[start] = false
			continue
		}
		end++
	}
	return max
}
