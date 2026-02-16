package main

import (
	"fmt"
	"strings"
)

// IsAnagramSort - Time: O(?) Space: O(?)
func IsAnagramSort(s string, t string) bool {
	// Your implementation
	return false
}

// IsAnagramMap - Time: O(?) Space: O(?)
func IsAnagramMap(s string, t string) bool {
	// Your implementation
	return false
}

type testCase struct {
	s        string
	t        string
	expected bool
	desc     string
}

func main() {
	tests := []testCase{
		{"anagram", "nagaram", true, "classic anagram"},
		{"rat", "car", false, "not anagram"},
		{"listen", "silent", true, "listen/silent"},
		{"hello", "world", false, "different letters"},
		{"a", "a", true, "single char same"},
		{"a", "b", false, "single char different"},
		{"ab", "ba", true, "two chars swapped"},
		{"ab", "cd", false, "two chars different"},
		{"abc", "ab", false, "different lengths"},
		{"a", "ab", false, "subset string"},
		{"aaa", "aaa", true, "all same chars"},
		{"aab", "baa", true, "repeated chars anagram"},
		{"aacc", "ccac", false, "same chars wrong count"},
		{"aabb", "abab", true, "interleaved duplicates"},
		{"abcde", "abcdf", false, "one char different"},
		{"abcd", "abce", false, "last char differs"},
		{"aaab", "aaba", true, "rearranged with repeats"},
		{strings.Repeat("a", 100), strings.Repeat("a", 100), true, "long same string"},
		{strings.Repeat("a", 99) + "b", "b" + strings.Repeat("a", 99), true, "long anagram"},
		{strings.Repeat("a", 100), strings.Repeat("a", 99) + "b", false, "long near miss"},
		{"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true, "full alphabet reversed"},
	}

	fmt.Println("======================================================================")
	fmt.Println("VALID ANAGRAM - Test Results")
	fmt.Println("======================================================================")

	type approach struct {
		name string
		fn   func(string, string) bool
	}
	approaches := []approach{
		{"Sorting", IsAnagramSort},
		{"HashMap", IsAnagramMap},
	}

	total := len(tests)
	for _, a := range approaches {
		fmt.Printf("\n%s:\n", a.name)
		passed := 0
		for i, tc := range tests {
			result := a.fn(tc.s, tc.t)
			ok := result == tc.expected
			if ok {
				passed++
			}
			status := "FAIL"
			if ok {
				status = "PASS"
			}
			sDisp := tc.s
			if len(sDisp) > 15 {
				sDisp = sDisp[:12] + "..."
			}
			tDisp := tc.t
			if len(tDisp) > 15 {
				tDisp = tDisp[:12] + "..."
			}
			fmt.Printf("  %2d. [%s] %s: '%s','%s' -> %v\n", i+1, status, tc.desc, sDisp, tDisp, result)
		}
		fmt.Printf("  Result: %d/%d\n", passed, total)
	}

	fmt.Println("\n======================================================================")
	fmt.Println("Questions:")
	fmt.Println("1. Which approach is better if inputs contain Unicode?")
	fmt.Println("2. Can you solve this with a single [26]int array?")
	fmt.Println("3. What's the space complexity of the sorting approach?")
}
