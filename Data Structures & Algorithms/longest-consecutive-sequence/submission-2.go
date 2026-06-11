func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _,v := range nums {
		m[v] = false
	}
	max := 0
	for k,_ := range m {
		if _,ok := m[k-1]; !ok{
			lm := 0
			i := k
			for {
				lm++
				i++
				if _,ok := m[i]; !ok {
					break
				} 
			}
			if lm > max {
				max = lm
			}
		}
	}
	return max
}
