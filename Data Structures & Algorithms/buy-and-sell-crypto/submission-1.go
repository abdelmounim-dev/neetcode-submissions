func maxProfit(prices []int) int {
	min, max, profit := prices[0], prices[0], 0
	for _,v := range prices {
		if v < min {
			min = v
			max = v
			continue
		}
		if v > max {
			max = v
			if max - min > profit {
				profit = max - min	
			}
		}
	}
	return profit
}
