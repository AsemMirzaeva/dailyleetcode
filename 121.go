package main

import "fmt"

func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0

    for _, price := range prices {
        if price < minPrice {
            minPrice = price
        }
        profit := price - minPrice
        if profit > maxProfit {
            maxProfit = profit
        }
    }

    return maxProfit
}

func main() {
	// p1 := []int{7, 1, 5, 3, 6, 4}
	p2 := []int{3,2,6,5,0,3}
	// p3 := []int{2, 4, 1}

	// fmt.Println(maxProfit(p1))

	fmt.Println(maxProfit(p2))
	// fmt.Println(maxProfit(p3))
}
