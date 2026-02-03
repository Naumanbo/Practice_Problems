package main

import "fmt"

// Tests: Dynamic programming, Kadane's algorithm
//
// Maximum Subarray (LeetCode 53)
// Find the contiguous subarray with the largest sum.

func maxSubArray(nums []int) int {
	// TODO: Implement Kadane's algorithm
	// Hint: At each position, decide: extend current subarray or start new?
	return 0
}

func main() {
	type testCase struct {
		nums     []int
		expected int
		desc     string
	}

	tests := []testCase{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, "mixed values"},
		{[]int{1}, 1, "single element"},
		{[]int{5, 4, -1, 7, 8}, 23, "mostly positive"},
		{[]int{-1}, -1, "single negative"},
		{[]int{-2, -1}, -1, "all negative"},
		{[]int{-2, -3, -1, -5}, -1, "all negative longer"},
		{[]int{1, 2, 3, 4}, 10, "all positive"},
		{[]int{-1, 0, -2}, 0, "zero is max"},
		{[]int{8, -19, 5, -4, 20}, 21, "recovery after negative"},
	}

	allPassed := true
	for _, tc := range tests {
		result := maxSubArray(tc.nums)
		if result != tc.expected {
			fmt.Printf("FAIL [%s]: maxSubArray(%v) = %d, expected %d\n",
				tc.desc, tc.nums, result, tc.expected)
			allPassed = false
		} else {
			fmt.Printf("PASS [%s]\n", tc.desc)
		}
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
