package main

import "fmt"

// Tests: Arrays, prefix/suffix products, no-division constraint
//
// Product of Array Except Self (LeetCode #238)
// Return array where output[i] = product of all nums except nums[i].
// Must run in O(n) time without division.

func productExceptSelf(nums []int) []int {
	return nil
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	nums     []int
	expected []int
	desc     string
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	tests := []testCase{
		{[]int{1, 2, 3, 4},       []int{24, 12, 8, 6},        "basic case"},
		{[]int{2, 3, 4, 5},       []int{60, 40, 30, 24},       "all positive"},
		{[]int{-1, 1, 0, -3, 3},  []int{0, 0, 9, 0, 0},        "contains zero"},
		{[]int{1, 2},             []int{2, 1},                  "two elements"},
		{[]int{3, 3, 3},          []int{9, 9, 9},               "all same"},
		{[]int{1, 1, 1, 1},       []int{1, 1, 1, 1},            "all ones"},
		{[]int{0, 0},             []int{0, 0},                  "all zeros"},
		{[]int{1, 0},             []int{0, 1},                  "one zero"},
		{[]int{-1, 2, 3},         []int{6, -3, -2},             "negative value"},
		{[]int{2, 2, 2, 2},       []int{8, 8, 8, 8},            "all twos"},
		{[]int{1, 2, 3, 4, 5},    []int{120, 60, 40, 30, 24},   "five elements"},
		{[]int{-1, -2, -3, -4},   []int{-24, -12, -8, -6},      "all negative"},
		{[]int{100, 1, 2},        []int{2, 200, 100},           "large first element"},
		{[]int{0, 1, 2, 3},       []int{6, 0, 0, 0},            "zero at start"},
		{[]int{1, 2, 0, 4},       []int{0, 0, 8, 0},            "zero in middle"},
		{[]int{2, 3},             []int{3, 2},                  "two elements v2"},
	}

	fmt.Println("======================================================================")
	fmt.Println("PRODUCT EXCEPT SELF - Test Results")
	fmt.Println("======================================================================")

	passed := 0
	for i, tc := range tests {
		result := productExceptSelf(tc.nums)
		ok := slicesEqual(result, tc.expected)
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
