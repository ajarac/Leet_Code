package best_time_to_buy_and_sell_stock

import "testing"

type Test struct {
	input    []int
	expected int
}

func TestMaxProfit(t *testing.T) {
	var tests = []Test{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, test := range tests {
		if output := maxProfit(test.input); output != test.expected {
			t.Error("TestMaxProfit Failed: {} inputted, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}
