package main

import "fmt"

// Tests: Floyd's cycle detection on arrays, two pointers, O(1) space
//
// Find the Duplicate Number (LeetCode #287)
// Array of n+1 integers where each is in [1,n]. Find the one duplicate.
// Must not modify the array and use only O(1) extra space.

func findDuplicate(nums []int) int {
	return 0
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	nums     []int
	expected int
	desc     string
}

func main() {
	tests := []testCase{
		{[]int{1, 3, 4, 2, 2},                    2, "basic case"},
		{[]int{3, 1, 3, 4, 2},                    3, "duplicate not adjacent"},
		{[]int{1, 1},                              1, "two elements"},
		{[]int{1, 1, 2},                           1, "three elements"},
		{[]int{1, 2, 3, 4, 4},                    4, "duplicate at end"},
		{[]int{1, 2, 3, 1},                       1, "duplicate wraps"},
		{[]int{2, 1, 2, 3},                       2, "duplicate early"},
		{[]int{6, 2, 4, 1, 3, 5, 6},              6, "six elements"},
		{[]int{9, 7, 4, 6, 3, 2, 8, 5, 1, 1},    1, "duplicate is 1"},
		{[]int{2, 5, 9, 6, 3, 8, 7, 1, 4, 9},    9, "duplicate is 9"},
		{[]int{3, 4, 8, 5, 9, 1, 6, 8, 7, 2},    8, "duplicate in middle"},
		{[]int{5, 1, 2, 3, 4, 5},                 5, "duplicate at boundaries"},
		{[]int{1, 2, 3, 2, 4},                    2, "five elements"},
		{[]int{4, 3, 2, 1, 4},                    4, "duplicate at front and back"},
		{[]int{1, 2, 1, 3},                       1, "four elements"},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 9}, 9, "reverse sorted with dup"},
	}

	fmt.Println("======================================================================")
	fmt.Println("FIND DUPLICATE NUMBER - Test Results")
	fmt.Println("======================================================================")

	passed := 0
	for i, tc := range tests {
		result := findDuplicate(tc.nums)
		ok := result == tc.expected
		if ok {
			passed++
		}
		status := "FAIL"
		if ok {
			status = "PASS"
		}
		fmt.Printf("  %2d. [%s] %s\n", i+1, status, tc.desc)
	}

	fmt.Println("======================================================================")
	fmt.Printf("Summary: %d/%d passed\n", passed, len(tests))
	fmt.Println("======================================================================")
}
