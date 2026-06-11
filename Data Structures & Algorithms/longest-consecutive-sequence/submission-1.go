func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _,v := range nums {
		m[v] = false
	}
	var r []int
	for k,_ := range m {
		if _,ok := m[k-1]; ok{
			m[k] = true
		} else {
			r = append(r, k)
		}
	}
	max := 0
	for _,v := range r {
		lm := 0
		i := v
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
	return max
}
