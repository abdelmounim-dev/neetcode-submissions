func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	max := 0
	for left < right {
		if heights[left] < heights[right] {
			if v := heights[left] * (right - left); v > max {
				max = v
			}
			left++
		} else {
			if v := heights[right] * (right - left); v > max {
				max = v
			}
			right--
		}
	}
	return max

}
