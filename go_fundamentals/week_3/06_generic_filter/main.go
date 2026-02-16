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

	// === Filter tests ===
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
	// single element passes
	if filtered := Filter([]int{5}, func(n int) bool { return n == 5 }); len(filtered) != 1 || filtered[0] != 5 {
		fmt.Println("FAIL: Filter single match")
		allPassed = false
	}
	// single element fails
	if filtered := Filter([]int{5}, func(n int) bool { return n == 3 }); len(filtered) != 0 {
		fmt.Println("FAIL: Filter single no match")
		allPassed = false
	}
	// filter strings by length
	shortWords := Filter(words, func(s string) bool { return len(s) <= 2 })
	if len(shortWords) != 2 || shortWords[0] != "go" || shortWords[1] != "is" {
		fmt.Println("FAIL: Filter strings by length")
		allPassed = false
	}
	// filter preserves order
	ordered := Filter([]int{5, 3, 1, 4, 2}, func(n int) bool { return n > 2 })
	if len(ordered) != 3 || ordered[0] != 5 || ordered[1] != 3 || ordered[2] != 4 {
		fmt.Println("FAIL: Filter preserves order")
		allPassed = false
	}

	// === Contains tests ===
	if Contains([]int{}, 1) {
		fmt.Println("FAIL: Contains empty slice")
		allPassed = false
	}
	if !Contains([]string{"a", "b", "c"}, "b") {
		fmt.Println("FAIL: Contains string")
		allPassed = false
	}
	// first element
	if !Contains([]int{10, 20, 30}, 10) {
		fmt.Println("FAIL: Contains first element")
		allPassed = false
	}
	// last element
	if !Contains([]int{10, 20, 30}, 30) {
		fmt.Println("FAIL: Contains last element")
		allPassed = false
	}
	// not present
	if Contains([]int{10, 20, 30}, 99) {
		fmt.Println("FAIL: Contains missing element")
		allPassed = false
	}
	// single element present
	if !Contains([]int{42}, 42) {
		fmt.Println("FAIL: Contains single present")
		allPassed = false
	}
	// single element absent
	if Contains([]int{42}, 7) {
		fmt.Println("FAIL: Contains single absent")
		allPassed = false
	}

	// === Map tests ===
	if mapped := Map([]int{}, func(n int) int { return n }); len(mapped) != 0 {
		fmt.Println("FAIL: Map empty slice")
		allPassed = false
	}
	// identity map
	identity := Map([]int{1, 2, 3}, func(n int) int { return n })
	if len(identity) != 3 || identity[0] != 1 || identity[2] != 3 {
		fmt.Println("FAIL: Map identity")
		allPassed = false
	}
	// type-changing map: int -> string
	strs := Map([]int{1, 2, 3}, func(n int) string { return fmt.Sprintf("%d", n) })
	if len(strs) != 3 || strs[0] != "1" || strs[2] != "3" {
		fmt.Println("FAIL: Map int to string")
		allPassed = false
	}
	// single element
	single := Map([]int{5}, func(n int) int { return n * 10 })
	if len(single) != 1 || single[0] != 50 {
		fmt.Println("FAIL: Map single element")
		allPassed = false
	}
	// map preserves order
	neg := Map([]int{1, 2, 3}, func(n int) int { return -n })
	if neg[0] != -1 || neg[1] != -2 || neg[2] != -3 {
		fmt.Println("FAIL: Map preserves order")
		allPassed = false
	}

	// === Reduce tests ===
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
	// single element
	if Reduce([]int{5}, 0, func(acc, n int) int { return acc + n }) != 5 {
		fmt.Println("FAIL: Reduce single element")
		allPassed = false
	}
	// reduce to count
	count := Reduce([]string{"a", "bb", "ccc"}, 0, func(acc int, s string) int { return acc + len(s) })
	if count != 6 {
		fmt.Println("FAIL: Reduce count chars")
		allPassed = false
	}
	// reduce multiply
	product2 := Reduce([]int{2, 3, 4}, 1, func(acc, n int) int { return acc * n })
	if product2 != 24 {
		fmt.Println("FAIL: Reduce multiply")
		allPassed = false
	}
	// reduce with initial 0 and multiply (everything becomes 0)
	zeroProduct := Reduce([]int{1, 2, 3}, 0, func(acc, n int) int { return acc * n })
	if zeroProduct != 0 {
		fmt.Println("FAIL: Reduce multiply with zero initial")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
