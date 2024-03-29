package Controllers

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		// Update the minimum price
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else {
			// Calculate the profit if selling at the current price
			currentProfit := prices[i] - minPrice

			// Update the maximum profit if the current profit is greater
			if currentProfit > maxProfit {
				maxProfit = currentProfit
			}
		}
	}

	return maxProfit
}

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

func ProfitMain() {
	// Example 1
	prices1 := []int{7, 1, 5, 3, 6, 4}
	fmt.Println("Example 1:", maxProfit(prices1)) // Output: 5

	// Example 2
	prices2 := []int{7, 6, 4, 3, 1}
	fmt.Println("Example 2:", maxProfit(prices2)) // Output: 0

	// Additional Test Case
	prices3 := []int{2, 4, 1}
	fmt.Println("Additional Test Case:", maxProfit(prices3)) // Output: 2
}
