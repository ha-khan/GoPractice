package main

import (
	"fmt"
	"math"
)

/*
The following sequence of stock prices: (310,315,275,295,260,270,290,230,255,250).
The maximum profit that can be made with one buy and one sell is 30—buy at 260 and sell at 290.
Note that 260 is not the lowest price, nor 290 the highest price.
Write a program that takes an array denoting the daily stock price, and returns the
maximum profit that could be made by buying and then selling one share of that stock.
*/
func main() {

	fmt.Println(computeBestBuySellSlow(310, 315, 275, 295, 260, 270, 290, 230, 255, 250))
	fmt.Println(computeBestBuySellFast(310, 315, 275, 295, 260, 270, 290, 230, 255, 250))
}

/*
O(n^2) will have two iterators where there is a slow (p1) and fast(p2)

*/
func computeBestBuySellSlow(stockprices ...int) int {

	var (
		largestSale int
	)

	for idx, p1 := range stockprices {
		for _, p2 := range stockprices[idx:] {
			if largestSale < p2-p1 {
				largestSale = p2 - p1
			}
		}
	}

	return largestSale
}

/*
Need to keep track of both the smallest seen stock price thus far && the largest delta seen

space/time ~ O(n)

potentially could also use a PriorityQueue <min> instead of our optimization

so O(n^2) Time/Space -> O(n) Time/Space -> O(n) Time, O(1) Space
*/
func computeBestBuySellFast(stockprices ...int32) int32 {
	var (
		smallestBuy = int32(math.MaxInt32)
		largestSell int32
	)

	for _, val := range stockprices {

		if val < smallestBuy {
			smallestBuy = val
		}

		if largestSell < val-smallestBuy {
			largestSell = val - smallestBuy
		}
	}

	return largestSell
}
