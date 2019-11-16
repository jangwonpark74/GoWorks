package max_profit_test

import (
	"fmt"
	"max_profit"
)

func Example_MaxProfit() {

	prices1 := []int{7, 1, 5, 3, 6, 4}
	prices2 := []int{1, 2, 3, 4, 5}
	p1 := max_profit.MaxProfit(prices1)
	p2 := max_profit.MaxProfit(prices2)

	fmt.Println(p1)
	fmt.Println(p2)
	//Output:
	//7
	//4
}
