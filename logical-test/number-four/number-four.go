package main

import (
	"fmt"
	"math"
)

func findBestProfit(prices []int) int {
	if len(prices) < 2 {
		return -1
	}

	minPrice := prices[0]
	maxProfit := math.MinInt64
	bestBuyPrice := prices[0]

	for _, price := range prices {
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
			bestBuyPrice = minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}

	return bestBuyPrice
}

func main() {
	examples := [][]int{
		{7, 8, 3, 10, 8},
		{5, 12, 11, 12, 10},
		{7, 18, 27, 10, 29},
		{20, 17, 15, 14, 10},
	}

	for i, prices := range examples {
		fmt.Printf("Soal %d: Harga saham: %v\n", i+1, prices)
		bestBuyPrice := findBestProfit(prices)
		if bestBuyPrice != -1 {
			fmt.Printf("Harga beli terbaik: %d\n\n", bestBuyPrice)
		} else {
			fmt.Println("Tidak cukup data untuk menghitung keuntungan\n")
		}
	}
}
