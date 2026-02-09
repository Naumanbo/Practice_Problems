package main

import "fmt"

// Tests: Generic functions, type parameters, comparable constraint
//
// 1. Implement Filter[T any](slice []T, predicate func(T) bool) []T
//    - Returns a new slice containing only elements where predicate returns true
//    - Should work with any type
//
// 2. Implement Contains[T comparable](slice []T, target T) bool
//    - Returns true if target exists in the slice
//    - T must be comparable (supports ==)
//
// 3. Implement Map[T any, U any](slice []T, transform func(T) U) []U
//    - Applies transform to each element and returns the results
//    - Input and output types can differ
//
// 4. Implement Reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U
//    - Reduces a slice to a single value using fn
//    - Starts with initial, applies fn(accumulator, element) for each element

// TODO: Implement Filter
func Filter[T any](slice []T, predicate func(T) bool) []T {
	return nil
}

// TODO: Implement Contains
func Contains[T comparable](slice []T, target T) bool {
	return false
}

// TODO: Implement Map
func Map[T any, U any](slice []T, transform func(T) U) []U {
	return nil
}

// TODO: Implement Reduce
func Reduce[T any, U any](slice []T, initial U, fn func(U, T) U) U {
	return initial
}

func main() {
	// Test Filter
	fmt.Println("=== Filter ===")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println("Evens:", evens) // [2 4 6 8 10]

	words := []string{"go", "is", "awesome", "and", "fast"}
	long := Filter(words, func(s string) bool { return len(s) > 2 })
	fmt.Println("Long words:", long) // [awesome and fast]

	// Test Contains
	fmt.Println("\n=== Contains ===")
	fmt.Println("Contains 5:", Contains(nums, 5))   // true
	fmt.Println("Contains 11:", Contains(nums, 11)) // false
	fmt.Println("Contains 'go':", Contains(words, "go"))     // true
	fmt.Println("Contains 'rust':", Contains(words, "rust")) // false

	// Test Map
	fmt.Println("\n=== Map ===")
	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println("Doubled:", doubled) // [2 4 6 8 10 12 14 16 18 20]

	lengths := Map(words, func(s string) int { return len(s) })
	fmt.Println("Lengths:", lengths) // [2 2 7 3 4]

	// Map with different input/output types
	asStrings := Map([]int{1, 2, 3}, func(n int) string {
		return fmt.Sprintf("#%d", n)
	})
	fmt.Println("As strings:", asStrings) // [#1 #2 #3]

	// Test Reduce
	fmt.Println("\n=== Reduce ===")
	sum := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println("Sum:", sum) // 55

	product := Reduce([]int{1, 2, 3, 4, 5}, 1, func(acc, n int) int { return acc * n })
	fmt.Println("Product:", product) // 120

	// Reduce to build a string
	sentence := Reduce(words, "", func(acc string, s string) string {
		if acc == "" {
			return s
		}
		return acc + " " + s
	})
	fmt.Println("Sentence:", sentence) // go is awesome and fast

	// Run test cases
	allPassed := true

	// Filter tests
	if filtered := Filter([]int{}, func(n int) bool { return true }); len(filtered) != 0 {
		fmt.Println("FAIL: Filter empty slice")
		allPassed = false
	}
	if filtered := Filter(nums, func(n int) bool { return n > 100 }); len(filtered) != 0 {
		fmt.Println("FAIL: Filter no matches")
		allPassed = false
	}
	if filtered := Filter(nums, func(n int) bool { return true }); len(filtered) != 10 {
		fmt.Println("FAIL: Filter all match")
		allPassed = false
	}

	// Contains tests
	if Contains([]int{}, 1) {
		fmt.Println("FAIL: Contains empty slice")
		allPassed = false
	}
	if !Contains([]string{"a", "b", "c"}, "b") {
		fmt.Println("FAIL: Contains string")
		allPassed = false
	}

	// Map tests
	if mapped := Map([]int{}, func(n int) int { return n }); len(mapped) != 0 {
		fmt.Println("FAIL: Map empty slice")
		allPassed = false
	}

	// Reduce tests
	if result := Reduce([]int{}, 42, func(acc, n int) int { return acc + n }); result != 42 {
		fmt.Println("FAIL: Reduce empty returns initial")
		allPassed = false
	}
	maxVal := Reduce([]int{3, 7, 2, 9, 1}, nums[0], func(acc, n int) int {
		if n > acc {
			return n
		}
		return acc
	})
	if maxVal != 9 {
		fmt.Println("FAIL: Reduce find max")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
