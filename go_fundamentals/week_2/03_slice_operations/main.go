package main

import "fmt"

// Tests: Slices, make, append, copy, range
//
// Implement the following slice operations:
// - Reverse(s []int) []int - returns a new slice with elements reversed
// - Filter(s []int, fn func(int) bool) []int - returns elements where fn returns true
// - Map(s []int, fn func(int) int) []int - applies fn to each element
// - Sum(s []int) int - returns sum of all elements
// - Contains(s []int, target int) bool - returns true if target exists in slice

// TODO: Implement Reverse
func Reverse(s []int) []int {
	return nil
}

// TODO: Implement Filter
func Filter(s []int, fn func(int) bool) []int {
	return nil
}

// TODO: Implement Map
func Map(s []int, fn func(int) int) []int {
	return nil
}

// TODO: Implement Sum
func Sum(s []int) int {
	return 0
}

// TODO: Implement Contains
func Contains(s []int, target int) bool {
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5}

	// Test Reverse
	fmt.Println("Original:", nums)
	fmt.Println("Reversed:", Reverse(nums)) // Expected: [5 4 3 2 1]
	fmt.Println("Original unchanged:", nums) // Should still be [1 2 3 4 5]

	// Test Filter (keep even numbers)
	isEven := func(n int) bool { return n%2 == 0 }
	fmt.Println("Even numbers:", Filter(nums, isEven)) // Expected: [2 4]

	// Test Map (double each number)
	double := func(n int) int { return n * 2 }
	fmt.Println("Doubled:", Map(nums, double)) // Expected: [2 4 6 8 10]

	// Test Sum
	fmt.Println("Sum:", Sum(nums)) // Expected: 15

	// Test Contains
	fmt.Println("Contains 3:", Contains(nums, 3)) // Expected: true
	fmt.Println("Contains 9:", Contains(nums, 9)) // Expected: false

	// Run test cases
	allPassed := true

	// Reverse tests
	if fmt.Sprint(Reverse([]int{1, 2, 3})) != "[3 2 1]" {
		fmt.Println("FAIL: Reverse([1,2,3])")
		allPassed = false
	}
	if fmt.Sprint(Reverse([]int{})) != "[]" {
		fmt.Println("FAIL: Reverse([])")
		allPassed = false
	}

	// Sum tests
	if Sum([]int{1, 2, 3}) != 6 {
		fmt.Println("FAIL: Sum([1,2,3])")
		allPassed = false
	}
	if Sum([]int{}) != 0 {
		fmt.Println("FAIL: Sum([])")
		allPassed = false
	}
	if Sum([]int{-1, 1}) != 0 {
		fmt.Println("FAIL: Sum([-1, 1])")
		allPassed = false
	}

	// Contains tests
	if !Contains([]int{1, 2, 3}, 2) {
		fmt.Println("FAIL: Contains([1,2,3], 2)")
		allPassed = false
	}
	if Contains([]int{1, 2, 3}, 5) {
		fmt.Println("FAIL: Contains([1,2,3], 5)")
		allPassed = false
	}

	// Filter tests
	isPositive := func(n int) bool { return n > 0 }
	if fmt.Sprint(Filter([]int{-1, 0, 1, 2}, isPositive)) != "[1 2]" {
		fmt.Println("FAIL: Filter positive")
		allPassed = false
	}

	// Map tests
	square := func(n int) int { return n * n }
	if fmt.Sprint(Map([]int{1, 2, 3}, square)) != "[1 4 9]" {
		fmt.Println("FAIL: Map square")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
