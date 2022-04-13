package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	gain := 0
	low := int(^uint(0) >> 1)

	for _, price := range prices {
		low = Min(low, price)
		gain = Max(gain, price-low)
	}

	return gain
}

func Max(i1 int, i2 int) int {
	if i1 > i2 {
		return i1
	} else {
		return i2
	}
}
func Min(i1 int, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}
