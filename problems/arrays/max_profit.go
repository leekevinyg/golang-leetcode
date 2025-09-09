package problems

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	maxProfitSoFar := 0
	buyIndex := 0
	sellIndex := 1

	for sellIndex < len(prices) {
		if prices[sellIndex] > prices[buyIndex] {
			profit := prices[sellIndex] - prices[buyIndex]
			if profit > maxProfitSoFar {
				maxProfitSoFar = profit
			}
		} else {
			buyIndex = sellIndex
		}
		sellIndex++
	}

	return maxProfitSoFar
}
