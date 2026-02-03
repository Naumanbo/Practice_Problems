package main

import "fmt"

// Tests: Arrays, single pass, tracking minimum
//
// Best Time to Buy and Sell Stock (LeetCode 121)
// Given prices where prices[i] is stock price on day i,
// return maximum profit from one buy-sell transaction.
// Return 0 if no profit possible.

func maxProfit(prices []int) int {
	// TODO: Implement solution
	// Hint: Track minimum price seen so far and maximum profit
	return 0
}

func main() {
	type testCase struct {
		prices   []int
		expected int
		desc     string
	}

	tests := []testCase{
		{[]int{7, 1, 5, 3, 6, 4}, 5, "basic case"},
		{[]int{7, 6, 4, 3, 1}, 0, "decreasing prices"},
		{[]int{1, 2, 3, 4, 5}, 4, "increasing prices"},
		{[]int{2, 4, 1}, 2, "buy early sell middle"},
		{[]int{3, 3, 3}, 0, "flat prices"},
		{[]int{1}, 0, "single element"},
		{[]int{2, 1}, 0, "two elements decreasing"},
		{[]int{1, 2}, 1, "two elements increasing"},
	}

	allPassed := true
	for _, tc := range tests {
		result := maxProfit(tc.prices)
		if result != tc.expected {
			fmt.Printf("FAIL [%s]: maxProfit(%v) = %d, expected %d\n",
				tc.desc, tc.prices, result, tc.expected)
			allPassed = false
		} else {
			fmt.Printf("PASS [%s]\n", tc.desc)
		}
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
