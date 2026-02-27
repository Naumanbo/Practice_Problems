// Key Takeaways:
// 1. Two approaches: hash set (O(n) space) stores visited node pointers,
//    Floyd's cycle detection (O(1) space) uses fast/slow pointers.
//    Always prefer Floyd's when space is constrained.
//
// 2. In Go, use map[*ListNode]bool to store visited nodes (pointer as key).
//    Go maps support pointer keys natively — no special hashing needed.
//
// 3. Floyd's: slow moves 1 step, fast moves 2. If they meet, cycle exists.
//    Check fast != nil && fast.Next != nil before each step to avoid nil dereference.
//
// 4. This same Floyd's algorithm works on arrays (LeetCode #287) by treating
//    array values as next pointers — same pattern, different data structure.
//
// Complexity: Hash set — O(n) time, O(n) space. Floyd's — O(n) time, O(1) space.

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// HasCycleSet - hash set approach - Time: O(?) Space: O(?)
func HasCycleSet(head *ListNode) bool {
	// Your implementation
	seen := make(map[*ListNode]bool)

	for head != nil {
		_, ok := seen[head]
		if !ok {
			seen[head] = true
		} else {
			return true
		}
		head = head.Next

	}

	return false
}

// HasCycleFloyd - Floyd's tortoise and hare - Time: O(?) Space: O(?)
func HasCycleFloyd(head *ListNode) bool {
	// Your implementation
	if head == nil || head.Next == nil {
		return false

	}
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}

	}
	return false
}

func makeListWithCycle(values []int, cyclePos int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	if cyclePos >= 0 && cyclePos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[cyclePos]
	}
	return nodes[0]
}

func makeRange(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = i
	}
	return s
}

type testCase struct {
	values   []int
	cyclePos int
	expected bool
	desc     string
}

func main() {
	tests := []testCase{
		{[]int{3, 2, 0, -4}, 1, true, "cycle at pos 1"},
		{[]int{1, 2}, 0, true, "cycle at head"},
		{[]int{1, 2, 3, 4}, 0, true, "cycle back to head"},
		{[]int{1, 2, 3, 4}, 2, true, "cycle in middle"},
		{[]int{1}, -1, false, "single node no cycle"},
		{[]int{1, 2}, -1, false, "two nodes no cycle"},
		{[]int{1, 2, 3, 4, 5}, -1, false, "five nodes no cycle"},
		{[]int{}, -1, false, "empty list"},
		{[]int{1}, 0, true, "self-loop"},
		{[]int{1, 2, 3}, -1, false, "three nodes no cycle"},
		{makeRange(100), -1, false, "large list no cycle"},
		{makeRange(100), 50, true, "large list cycle at 50"},
		{makeRange(100), 0, true, "large list cycle at head"},
		{makeRange(100), 99, true, "large list self-loop tail"},
		{[]int{1, 2, 3, 4, 5}, 4, true, "tail self-loop"},
		{[]int{1, 2, 3, 4, 5}, 3, true, "cycle at second to last"},
		{[]int{1, 2, 3, 4, 5}, 1, true, "cycle at pos 1 five nodes"},
		{[]int{-1, -2, -3}, -1, false, "negative values no cycle"},
		{[]int{-1, -2, -3}, 0, true, "negative values with cycle"},
		{[]int{1, 2}, 1, true, "two nodes tail self-loop"},
	}

	type approach struct {
		name string
		fn   func(*ListNode) bool
	}
	approaches := []approach{
		{"HashSet", HasCycleSet},
		{"Floyd's", HasCycleFloyd},
	}

	fmt.Println("======================================================================")
	fmt.Println("LINKED LIST CYCLE - Test Results")
	fmt.Println("======================================================================")

	total := len(tests)
	for _, a := range approaches {
		fmt.Printf("\n%s:\n", a.name)
		passed := 0
		for i, tc := range tests {
			head := makeListWithCycle(tc.values, tc.cyclePos)
			result := a.fn(head)
			ok := result == tc.expected
			if ok {
				passed++
			}
			status := "FAIL"
			if ok {
				status = "PASS"
			}
			vals := fmt.Sprintf("%v", tc.values)
			if len(tc.values) > 8 {
				vals = fmt.Sprintf("%v...", tc.values[:5])
			}
			fmt.Printf("  %2d. [%s] %s: vals=%s, pos=%d -> %v\n", i+1, status, tc.desc, vals, tc.cyclePos, result)
		}
		fmt.Printf("  Result: %d/%d\n", passed, total)
	}

	fmt.Println("\n======================================================================")
	fmt.Println("Questions:")
	fmt.Println("1. Why does Floyd's algorithm guarantee fast meets slow?")
	fmt.Println("2. What's the trade-off between hash set and Floyd's?")
	fmt.Println("3. Follow-up: How do you find WHERE the cycle begins? (LeetCode #142)")
}
