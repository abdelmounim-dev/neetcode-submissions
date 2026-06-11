func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i <= j; {
		for i <= j && (s[i] < 'a' || s[i] > 'z') {
			i++
		}
		for  i <= j && (s[j] < 'a' || s[j] > 'z') {
			j--
		}
		if  i <= j && s[i] != s[j] {
			return false 
		}
		i++
		j--
	}
	return true
}
