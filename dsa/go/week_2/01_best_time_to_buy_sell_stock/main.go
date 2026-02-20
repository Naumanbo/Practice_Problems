// Key Takeaways:
// 1. Classic single-pass greedy: track the minimum price seen so far, and at each
//    step compute the profit if you sold today. No need to look ahead.
//
// 2. In Go, use `math.MaxInt` (from "math" package) to initialize a minimum tracker.
//    Equivalent to sys.maxsize in Python or INT_MAX in C++.
//
// 3. Initialize max_profit to 0 so unprofitable cases naturally return 0
//    without extra checks — the zero value works in your favor here.
//
// Complexity: Time O(n), Space O(1)

package main

import (
	"fmt"
	"math"
)

// Tests: Arrays, single pass, tracking minimum
//
// Best Time to Buy and Sell Stock (LeetCode 121)
// Given prices where prices[i] is stock price on day i,
// return maximum profit from one buy-sell transaction.
// Return 0 if no profit possible.

func maxProfit(prices []int) int {
	// TODO: Implement solution
	// Hint: Track minimum price seen so far and maximum profit
	minimum_price := math.MaxInt
	maximum_profit := 0

	for _, price := range prices {
		if price < minimum_price {
			minimum_price = price
		} else if price-minimum_price > maximum_profit {
			maximum_profit = price - minimum_price
		}
	}
	return maximum_profit
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
		{[]int{10, 1, 10}, 9, "V-shape recovery"},
		{[]int{3, 8, 2, 5}, 5, "peak then valley"},
		{[]int{1, 5, 2, 8}, 7, "multiple peaks best at end"},
		{[]int{5, 4, 3, 2, 1, 10}, 9, "min at end then spike"},
		{[]int{0, 0, 0}, 0, "all zeros"},
		{[]int{5, 5}, 0, "two same"},
		{[]int{3, 1, 100, 2, 4}, 99, "spike in middle"},
		{[]int{1, 10, 1, 10, 1}, 9, "oscillating"},
		{[]int{0, 1}, 1, "zero then one"},
		{[]int{10000, 1, 10000}, 9999, "near constraint max"},
		{[]int{2, 1, 2, 0, 1}, 1, "multiple dips"},
		{[]int{1, 1, 1, 1, 2}, 1, "flat then rise"},
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
