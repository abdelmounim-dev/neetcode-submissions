// (number of characters of sub string - maxFreq > k) when this is reached we check another substr
// (number of characters of sub string - maxFreq <= k) we can continue
// longest substr where character other than the most freq <= k
func characterReplacement(s string, k int) int {
	freq := make(map[byte]int)
	maxFreq := 0
	maxLen := 0
	start := 0
	for end := 0; end < len(s); {
		freq[s[end]]++
		if freq[s[end]] > maxFreq {
			maxFreq = freq[s[end]]
		}
		if end-start+1 - maxFreq > k {
			freq[s[start]]--
			start++
		} else {
			if end-start+1 > maxLen {
				maxLen++
			}
			end++
		}
	}
	return maxLen
}
