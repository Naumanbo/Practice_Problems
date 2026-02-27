package main

import "fmt"

// Tests: Sliding window, hash map/set, two pointers
//
// Longest Substring Without Repeating Characters (LeetCode #3)
// Find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	return 0
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	s        string
	expected int
	desc     string
}

func main() {
	tests := []testCase{
		{"abcabcbb",  3, "basic case"},
		{"bbbbb",     1, "all same char"},
		{"",          0, "empty string"},
		{"pwwkew",    3, "repeat in middle"},
		{"a",         1, "single char"},
		{"au",        2, "two unique"},
		{"dvdf",      3, "overlap window"},
		{"abcdefg",   7, "all unique"},
		{"aab",       2, "repeat at start"},
		{"tmmzuxt",   5, "repeat at both ends"},
		{" ",         1, "single space"},
		{"abba",      2, "palindrome"},
		{"abcbda",    4, "complex window"},
		{"aababcabcd",4, "increasing unique suffix"},
		{"abcdeabcde",5, "repeated block"},
		{"ohvhjdml",  6, "long unique suffix"},
	}

	fmt.Println("======================================================================")
	fmt.Println("LONGEST SUBSTRING - Test Results")
	fmt.Println("======================================================================")

	passed := 0
	for i, tc := range tests {
		result := lengthOfLongestSubstring(tc.s)
		ok := result == tc.expected
		if ok {
			passed++
		}
		status := "FAIL"
		if ok {
			status = "PASS"
		}
		fmt.Printf("  %2d. [%s] %s: %q -> %d\n", i+1, status, tc.desc, tc.s, result)
	}

	fmt.Println("======================================================================")
	fmt.Printf("Summary: %d/%d passed\n", passed, len(tests))
	fmt.Println("======================================================================")
}
