package main

/*
DSA Problem 1: Two Sum

Tests: Hash map usage, time/space complexity analysis, array traversal

Difficulty: Easy
Source: LeetCode #1

Problem:
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

Constraints:
  - 2 <= len(nums) <= 10^4
  - -10^9 <= nums[i] <= 10^9
  - Only one valid answer exists.
  - You may not use the same element twice.
*/

import "fmt"

// TwoSumBrute solves using brute force approach
// Time: O(?)
// Space: O(?)
func TwoSumBrute(nums []int, target int) []int {
	// Your implementation
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}

	}
	return nil
}

// TwoSumOptimal solves using hash map approach
// Time: O(?)
// Space: O(?)
func TwoSumOptimal(nums []int, target int) []int {
	// Your implementation
	seen := make(map[int]int, 0)
	for i, val := range nums {
		requiredValue := target - val
		requiredIndex, ok := seen[requiredValue]
		if ok {
			return []int{i, requiredIndex}
		}
		seen[val] = i
	}

	return nil
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	nums   []int
	target int
	desc   string
}

func validate(result []int, nums []int, target int) bool {
	if result == nil || len(result) != 2 {
		return false
	}
	if result[0] == result[1] {
		return false
	}
	if result[0] < 0 || result[0] >= len(nums) || result[1] < 0 || result[1] >= len(nums) {
		return false
	}
	return nums[result[0]]+nums[result[1]] == target
}

func copySlice(s []int) []int {
	c := make([]int, len(s))
	copy(c, s)
	return c
}

func main() {
	testCases := []testCase{
		// Basic cases
		{[]int{2, 7, 11, 15}, 9, "Basic - first two elements"},
		{[]int{3, 2, 4}, 6, "Basic - middle elements"},
		{[]int{3, 3}, 6, "Duplicate values"},

		// Edge cases - array size
		{[]int{1, 2}, 3, "Minimum size (2 elements)"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19, "Larger array"},

		// Edge cases - negative numbers
		{[]int{-1, -2, -3, -4}, -6, "All negative"},
		{[]int{-3, 4, 3, 90}, 0, "Negative + positive = 0"},
		{[]int{5, -5, 10}, 0, "Sum to zero"},

		// Edge cases - zeros
		{[]int{0, 4, 3, 0}, 0, "Two zeros"},
		{[]int{0, 1, 2}, 2, "Zero + positive"},

		// Edge cases - large numbers
		{[]int{1000000000, 2, 1000000000}, 2000000000, "Large numbers"},

		// Edge cases - position variations
		{[]int{1, 5, 8, 3}, 4, "Answer at start and end"},
		{[]int{4, 5, 1, 2}, 6, "Answer in middle"},
		{[]int{5, 5, 5, 5}, 10, "All same values"},
	}

	fmt.Println("======================================================================")
	fmt.Println("TWO SUM - Test Results")
	fmt.Println("======================================================================")

	brutePassed, optimalPassed := 0, 0
	total := len(testCases)

	for i, tc := range testCases {
		bruteCopy := copySlice(tc.nums)
		optimalCopy := copySlice(tc.nums)

		bruteResult := TwoSumBrute(bruteCopy, tc.target)
		optimalResult := TwoSumOptimal(optimalCopy, tc.target)

		bruteOk := validate(bruteResult, tc.nums, tc.target)
		optimalOk := validate(optimalResult, tc.nums, tc.target)

		if bruteOk {
			brutePassed++
		}
		if optimalOk {
			optimalPassed++
		}

		bruteStatus := "FAIL"
		if bruteOk {
			bruteStatus = "PASS"
		}
		optimalStatus := "FAIL"
		if optimalOk {
			optimalStatus = "PASS"
		}

		fmt.Printf("\n%d. %s\n", i+1, tc.desc)
		fmt.Printf("   Input: nums=%v, target=%d\n", tc.nums, tc.target)
		fmt.Printf("   Brute:   [%s] %v\n", bruteStatus, bruteResult)
		fmt.Printf("   Optimal: [%s] %v\n", optimalStatus, optimalResult)
	}

	fmt.Println("\n======================================================================")
	fmt.Printf("Summary: Brute %d/%d | Optimal %d/%d\n", brutePassed, total, optimalPassed, total)
	fmt.Println("======================================================================")
	fmt.Println("\nQuestions:")
	fmt.Println("1. What is the time complexity of brute force? Why?")
	fmt.Println("2. What is the time complexity of the hash map approach? Why?")
	fmt.Println("3. Can this be solved with two pointers? What's the prerequisite?")
}
