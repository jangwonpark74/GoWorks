package max_profit

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).

Example 1:

Input: [7,1,5,3,6,4]
Output: 7
Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
             Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.


*/

func MaxProfit(prices []int) int {

	total_profit := 0
	start := 0
	end := 0

	if len(prices) < 2 {
		return 0
	}

	for i := 1; i < len(prices); i++ {
		if prices[i-1] > prices[i] {
			total_profit += prices[end] - prices[start]
			start = i
		}
		end += 1

		if i == (len(prices) - 1) {
			total_profit += prices[end] - prices[start]
		}
	}
	return total_profit
}
