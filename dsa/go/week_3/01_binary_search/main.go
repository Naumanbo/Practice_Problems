package main

import "fmt"

// BinarySearchIterative - Time: O(?) Space: O(?)
func BinarySearchIterative(nums []int, target int) int {
	// Your implementation
	return -1
}

// BinarySearchRecursive - Time: O(?) Space: O(?)
func BinarySearchRecursive(nums []int, target int) int {
	// Your implementation
	return -1
}

type testCase struct {
	nums     []int
	target   int
	expected int
	desc     string
}

func makeRange(start, end int) []int {
	s := make([]int, end-start+1)
	for i := range s {
		s[i] = start + i
	}
	return s
}

func runTests(fn func([]int, int) int, name string, tests []testCase) int {
	fmt.Printf("\n%s:\n", name)
	passed := 0
	for i, tc := range tests {
		cp := make([]int, len(tc.nums))
		copy(cp, tc.nums)
		result := fn(cp, tc.target)
		ok := result == tc.expected
		if ok {
			passed++
		}
		status := "FAIL"
		if ok {
			status = "PASS"
		}
		fmt.Printf("  %2d. [%s] %s: target=%d -> %d (expected %d)\n", i+1, status, tc.desc, tc.target, result, tc.expected)
	}
	return passed
}

func main() {
	tests := []testCase{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4, "target in middle"},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1, "target not found"},
		{[]int{1, 2, 3, 4, 5}, 1, 0, "target at start"},
		{[]int{1, 2, 3, 4, 5}, 5, 4, "target at end"},
		{[]int{1, 2, 3, 4, 5}, 3, 2, "target in center"},
		{[]int{5}, 5, 0, "single element found"},
		{[]int{5}, 3, -1, "single element not found"},
		{[]int{1, 2}, 1, 0, "two elements first"},
		{[]int{1, 2}, 2, 1, "two elements second"},
		{[]int{1, 2}, 3, -1, "two elements not found"},
		{[]int{-10, -5, 0, 5, 10}, -5, 1, "negative target"},
		{[]int{-10, -5, 0, 5, 10}, 0, 2, "zero target"},
		{[]int{-100, -50, -10, -1}, -100, 0, "all negative first"},
		{[]int{-100, -50, -10, -1}, -1, 3, "all negative last"},
		{[]int{-100, -50, -10, -1}, 5, -1, "all negative not found"},
		{[]int{1, 2, 3, 4, 5}, 0, -1, "target below range"},
		{[]int{1, 2, 3, 4, 5}, 6, -1, "target above range"},
		{makeRange(1, 100), 50, 49, "larger array middle"},
		{makeRange(1, 100), 1, 0, "larger array first"},
		{makeRange(1, 100), 100, 99, "larger array last"},
		{makeRange(1, 100), 101, -1, "larger array not found"},
		{[]int{-9999, 0, 9999}, 9999, 2, "near constraint bounds"},
	}

	fmt.Println("======================================================================")
	fmt.Println("BINARY SEARCH - Test Results")
	fmt.Println("======================================================================")

	total := len(tests)
	iterPassed := runTests(BinarySearchIterative, "Iterative", tests)
	recPassed := runTests(BinarySearchRecursive, "Recursive", tests)

	fmt.Println("\n======================================================================")
	fmt.Printf("Summary: Iterative %d/%d | Recursive %d/%d\n", iterPassed, total, recPassed, total)
	fmt.Println("======================================================================")
	fmt.Println("\nQuestions:")
	fmt.Println("1. Why must the array be sorted for binary search?")
	fmt.Println("2. What happens if you use (left + right) / 2 vs left + (right-left)/2?")
	fmt.Println("3. When would you use sort.Search from the standard library?")
}
