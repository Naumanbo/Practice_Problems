package main

import "fmt"

// Tests: Linked lists, two pointers, iteration
//
// Merge Two Sorted Lists (LeetCode 21)
// Merge two sorted linked lists into one sorted list.

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// TODO: Implement solution
	// Hint: Use a dummy head node to simplify edge cases
	return nil
}

// Helper: convert slice to linked list
func sliceToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	current := head
	for _, val := range arr[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Helper: convert linked list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// Helper: compare slices
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
	type testCase struct {
		list1    []int
		list2    []int
		expected []int
		desc     string
	}

	tests := []testCase{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}, "basic merge"},
		{[]int{}, []int{}, []int{}, "both empty"},
		{[]int{}, []int{0}, []int{0}, "first empty"},
		{[]int{1}, []int{}, []int{1}, "second empty"},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}, "interleaved"},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}, "no overlap"},
		{[]int{5}, []int{1, 2, 3}, []int{1, 2, 3, 5}, "single vs multiple"},
		{[]int{1, 1, 1}, []int{1, 1}, []int{1, 1, 1, 1, 1}, "all duplicates"},
		{[]int{-3, -1, 0}, []int{-2, 5}, []int{-3, -2, -1, 0, 5}, "negative values"},
		{[]int{1}, []int{2}, []int{1, 2}, "single element each"},
		{[]int{2}, []int{1}, []int{1, 2}, "single elements reversed"},
		{[]int{1}, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "one short one long"},
		{[]int{5, 5, 5}, []int{5, 5, 5}, []int{5, 5, 5, 5, 5, 5}, "all same values"},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}, "identical lists"},
		{[]int{1000000}, []int{-1000000}, []int{-1000000, 1000000}, "large values"},
		{[]int{-100, -50, 0, 50, 100}, []int{-75, -25, 25, 75}, []int{-100, -75, -50, -25, 0, 25, 50, 75, 100}, "interleaved with negatives"},
		{[]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "no overlap longer"},
	}

	allPassed := true
	for _, tc := range tests {
		l1 := sliceToList(tc.list1)
		l2 := sliceToList(tc.list2)
		result := listToSlice(mergeTwoLists(l1, l2))

		// Handle nil result for empty expected
		if result == nil {
			result = []int{}
		}

		if !slicesEqual(result, tc.expected) {
			fmt.Printf("FAIL [%s]: merge(%v, %v) = %v, expected %v\n",
				tc.desc, tc.list1, tc.list2, result, tc.expected)
			allPassed = false
		} else {
			fmt.Printf("PASS [%s]\n", tc.desc)
		}
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
