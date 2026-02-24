// Key Takeaways:
// 1. Kadane's algorithm is identical across languages: at each element, decide
//    whether to extend or start fresh: `currSum = max(nums[i], currSum + nums[i])`.
//
// 2. Initialize to nums[0], not 0 — all-negative inputs would return 0 incorrectly.
//
// 3. In Go, use `len(slice)` to get the size (vs `nums.size()` in C++, `len(nums)` in Python).
//    Go's built-in `max()` is available since Go 1.21 — no import needed, unlike
//    C++ which requires <algorithm> for std::max.
//
// 4. Loop from index 1 since index 0 seeds both currSum and maxSum.
//
// Complexity: Time O(n), Space O(1)

package main

import "fmt"

// Tests: Dynamic programming, Kadane's algorithm
//
// Maximum Subarray (LeetCode 53)
// Find the contiguous subarray with the largest sum.

func maxSubArray(nums []int) int {
	// TODO: Implement Kadane's algorithm
	// Hint: At each position, decide: extend current subarray or start new?
	currSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(currSum, maxSum)
	}
	return maxSum
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
		{[]int{1, -1, 1, -1, 1}, 1, "alternating"},
		{[]int{5, 6, -100, 7, 8}, 15, "large negative breaks subarray"},
		{[]int{10, -1, -1, -1, -1}, 10, "subarray at start"},
		{[]int{-1, -1, -1, -1, 10}, 10, "subarray at end"},
		{[]int{5, -10, 5}, 5, "two equal subarrays"},
		{[]int{100}, 100, "single large positive"},
		{[]int{10, -5, 10, -5, 10}, 20, "worth keeping negatives"},
		{[]int{1, -1, 1, -1, 2}, 2, "max at end"},
		{[]int{3, -1, 2, -1, 4}, 7, "entire array is max"},
		{[]int{0}, 0, "single zero"},
		{[]int{-10000}, -10000, "single large negative"},
		{[]int{1, 2, -1, 3, -2, 4}, 7, "scattered negatives"},
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
