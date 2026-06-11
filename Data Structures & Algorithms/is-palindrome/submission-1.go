func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; {
		for s[i] < 'a' || s[i] > 'z' {
			i++
		}
		for s[j] < 'a' || s[j] > 'z' {
			j--
		}
		if s[i] != s[j] {
			return false 
		}
		i++
		j--
	}
	return true
}
