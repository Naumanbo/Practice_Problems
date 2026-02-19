package main

/*
DSA Problem 2: Valid Parentheses

Tests: Stack data structure, matching pairs, string traversal

Difficulty: Easy
Source: LeetCode #20

Problem:
Given a string s containing just '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

Valid if:
1. Open brackets closed by same type
2. Open brackets closed in correct order
3. Every close bracket has corresponding open bracket
*/

import (
	"fmt"
	"strings"
)

// IsValid checks if parentheses string is valid
// Time: O(?)
// Space: O(?)
// Hint: Use a slice as a stack
func IsValid(s string) bool {
	// Your implementation
	stack := make([]rune, 0)
	matching := make(map[rune]rune, 0)
	matching['('] = ')'
	matching['['] = ']'
	matching['{'] = '}'

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) > 0 && c == matching[stack[len(stack)-1]] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	s        string
	expected bool
	desc     string
}

func main() {
	testCases := []testCase{
		// Basic valid cases
		{"()", true, "Single pair - parentheses"},
		{"[]", true, "Single pair - brackets"},
		{"{}", true, "Single pair - braces"},
		{"()[]{}", true, "Multiple pairs sequential"},
		{"{[]}", true, "Nested brackets"},
		{"([{}])", true, "Deeply nested"},

		// Basic invalid cases
		{"(]", false, "Mismatched types"},
		{"([)]", false, "Wrong order - interleaved"},
		{"{[}]", false, "Wrong order - interleaved v2"},

		// Edge cases - empty
		{"", true, "Empty string"},

		// Edge cases - single bracket
		{"(", false, "Single open paren"},
		{")", false, "Single close paren"},
		{"[", false, "Single open bracket"},
		{"}", false, "Single close brace"},

		// Edge cases - unbalanced
		{"(()", false, "Extra open at start"},
		{"())", false, "Extra close at end"},
		{"(())", true, "Balanced nested"},
		{"((()))", true, "Triple nested"},

		// Edge cases - long strings
		{strings.Repeat("()", 100), true, "Long valid string"},
		{strings.Repeat("(", 50) + strings.Repeat(")", 50), true, "Many nested"},
		{strings.Repeat("(", 50) + strings.Repeat(")", 49), false, "Off by one"},

		// Edge cases - complex patterns
		{"{[()()]}", true, "Complex valid"},
		{"[({})]", true, "All types nested"},
		{"[(])", false, "Complex invalid"},

		// Edge case - close before open
		{")(", false, "Close before open"},
		{"}{", false, "Close before open v2"},
	}

	fmt.Println("======================================================================")
	fmt.Println("VALID PARENTHESES - Test Results")
	fmt.Println("======================================================================")

	passed := 0
	total := len(testCases)

	for i, tc := range testCases {
		result := IsValid(tc.s)
		ok := result == tc.expected
		if ok {
			passed++
		}

		status := "FAIL"
		if ok {
			status = "PASS"
		}

		displayS := tc.s
		if len(displayS) > 20 {
			displayS = displayS[:17] + "..."
		}

		fmt.Printf("%2d. [%s] %s\n", i+1, status, tc.desc)
		fmt.Printf("    Input: '%s' | Got: %v | Expected: %v\n", displayS, result, tc.expected)
	}

	fmt.Println("\n======================================================================")
	fmt.Printf("Summary: %d/%d passed\n", passed, total)
	fmt.Println("======================================================================")
	fmt.Println("\nQuestions:")
	fmt.Println("1. Why is a stack the right data structure for this problem?")
	fmt.Println("2. What's the time/space complexity?")
	fmt.Println("3. Could you solve this without a stack? What's the tradeoff?")
}
