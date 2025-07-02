package main

import (
	"fmt"
	"math"
)

func calculateStock(arr []int) int {
	min := math.MaxInt32
	maxProfit := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
			fmt.Printf("min price is: %v \n", min)
		} else {
			if arr[i]-min > maxProfit {
				maxProfit = arr[i] - min
			}
		}

	}
	return maxProfit
}

func main() {
	arr := []int{
		7, 1, 5, 3, 6, 4,
	}
	maxprofirt := calculateStock(arr)
	fmt.Println("Your Profit is:", maxprofirt)
}
