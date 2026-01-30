package main

/*
DSA Problem 3: Reverse Linked List

Tests: Linked list traversal, pointer manipulation, iterative vs recursive

Difficulty: Easy
Source: LeetCode #206

Problem:
Given the head of a singly linked list, reverse the list and return it.
*/

import "fmt"

// ListNode represents a node in a singly linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseIterative reverses using iterative approach with three pointers
// Time: O(?)
// Space: O(?)
func ReverseIterative(head *ListNode) *ListNode {
	// Your implementation
	return nil
}

// ReverseRecursive reverses using recursive approach
// Time: O(?)
// Space: O(?) - consider call stack
func ReverseRecursive(head *ListNode) *ListNode {
	// Your implementation
	return nil
}

// =============================================================================
// Helper functions
// =============================================================================

func sliceToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, v := range arr[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
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

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	input    []int
	expected []int
	desc     string
}

func runTests(fn func(*ListNode) *ListNode, name string, testCases []testCase) int {
	fmt.Printf("\n%s:\n", name)
	passed := 0

	for i, tc := range testCases {
		head := sliceToList(tc.input)
		result := listToSlice(fn(head))
		ok := slicesEqual(result, tc.expected)
		if ok {
			passed++
		}

		status := "FAIL"
		if ok {
			status = "PASS"
		}

		displayIn := fmt.Sprintf("%v", tc.input)
		displayOut := fmt.Sprintf("%v", result)
		if len(tc.input) > 8 {
			displayIn = fmt.Sprintf("%v...", tc.input[:5])
		}
		if len(result) > 8 {
			displayOut = fmt.Sprintf("%v...", result[:5])
		}

		fmt.Printf("  %2d. [%s] %s: %s -> %s\n", i+1, status, tc.desc, displayIn, displayOut)
	}

	return passed
}

func main() {
	testCases := []testCase{
		// Base cases
		{[]int{}, []int{}, "Empty list"},
		{[]int{1}, []int{1}, "Single element"},
		{[]int{1, 2}, []int{2, 1}, "Two elements"},

		// Basic cases
		{[]int{1, 2, 3}, []int{3, 2, 1}, "Three elements"},
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, "Five elements"},

		// Edge cases - duplicates
		{[]int{1, 1, 1}, []int{1, 1, 1}, "All same values"},
		{[]int{1, 2, 2, 1}, []int{1, 2, 2, 1}, "Palindrome"},
		{[]int{1, 1, 2, 2}, []int{2, 2, 1, 1}, "Pairs"},

		// Edge cases - negative numbers
		{[]int{-1, -2, -3}, []int{-3, -2, -1}, "Negative values"},
		{[]int{-1, 0, 1}, []int{1, 0, -1}, "Mixed signs"},

		// Edge cases - large values
		{[]int{1000000, 2000000}, []int{2000000, 1000000}, "Large values"},

		// Longer lists
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, "0-9 sequence"},
	}

	fmt.Println("======================================================================")
	fmt.Println("REVERSE LINKED LIST - Test Results")
	fmt.Println("======================================================================")

	total := len(testCases)
	iterPassed := runTests(ReverseIterative, "Iterative", testCases)
	recPassed := runTests(ReverseRecursive, "Recursive", testCases)

	fmt.Println("\n======================================================================")
	fmt.Printf("Summary: Iterative %d/%d | Recursive %d/%d\n", iterPassed, total, recPassed, total)
	fmt.Println("======================================================================")
	fmt.Println("\nQuestions:")
	fmt.Println("1. Draw pointer changes for [1,2,3] step by step (iterative).")
	fmt.Println("2. What's the space complexity difference between approaches?")
	fmt.Println("3. Why prefer iterative in production code?")
}
