package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	buy1 := math.MinInt32
	sell1 := 0

	buy2 := math.MinInt32
	sell2 := 0

	for _, price := range prices {

		// mua lần 1
		buy1 = max(buy1, -price)

		// bán lần 1
		sell1 = max(sell1, buy1+price)

		// mua lần 2
		buy2 = max(buy2, sell1-price)

		// bán lần 2
        
		sell2 = max(sell2, buy2+price)
	}

	return sell2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{3,3,5,0,0,3,1,4}
	fmt.Println(maxProfit(prices))
}