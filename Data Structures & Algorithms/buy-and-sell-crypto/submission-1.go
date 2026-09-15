func maxProfit(prices []int) int {
	profit := 0
	bestBuy := prices[0]
	for _, sell := range prices {
		bestBuy = min(bestBuy, sell)
		profit = max(profit, sell-bestBuy)
	}
	return profit
}
