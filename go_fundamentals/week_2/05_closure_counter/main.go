package main

import "fmt"

// Tests: Closures, function values, state encapsulation
//
// Implement the following closure factories:
// - Counter() func() int - returns a function that increments and returns count (starting at 0)
// - Accumulator(initial int) func(int) int - returns a function that adds to running total
// - Limiter(max int) func() (int, bool) - returns count and whether limit reached
// - Memoize(fn func(int) int) func(int) int - returns memoized version of fn

// TODO: Implement Counter
// Each call to the returned function increments and returns the count
// First call returns 1, second returns 2, etc.
func Counter() func() int {
	return nil
}

// TODO: Implement Accumulator
// Returns a function that adds its argument to a running total
func Accumulator(initial int) func(int) int {
	return nil
}

// TODO: Implement Limiter
// Returns count (starting at 1) and true while count <= max
// After reaching max, returns max and false
func Limiter(max int) func() (int, bool) {
	return nil
}

// TODO: Implement Memoize
// Returns a cached version of fn - if called with same input, returns cached result
func Memoize(fn func(int) int) func(int) int {
	return nil
}

func main() {
	// Test Counter
	fmt.Println("=== Counter ===")
	count := Counter()
	fmt.Println(count()) // 1
	fmt.Println(count()) // 2
	fmt.Println(count()) // 3

	count2 := Counter() // separate counter
	fmt.Println(count2()) // 1 (independent)

	// Test Accumulator
	fmt.Println("\n=== Accumulator ===")
	acc := Accumulator(10)
	fmt.Println(acc(5))  // 15
	fmt.Println(acc(3))  // 18
	fmt.Println(acc(-8)) // 10

	// Test Limiter
	fmt.Println("\n=== Limiter ===")
	limit := Limiter(3)
	for i := 0; i < 5; i++ {
		n, ok := limit()
		fmt.Printf("n=%d, ok=%v\n", n, ok)
	}
	// Expected: n=1,ok=true; n=2,ok=true; n=3,ok=true; n=3,ok=false; n=3,ok=false

	// Test Memoize
	fmt.Println("\n=== Memoize ===")
	callCount := 0
	expensiveFn := func(n int) int {
		callCount++
		fmt.Printf("  Computing for %d...\n", n)
		return n * n
	}

	memoized := Memoize(expensiveFn)
	fmt.Println("Result:", memoized(5)) // Computes
	fmt.Println("Result:", memoized(5)) // Uses cache
	fmt.Println("Result:", memoized(3)) // Computes
	fmt.Println("Result:", memoized(5)) // Uses cache
	fmt.Printf("Total calls to expensive function: %d\n", callCount) // Should be 2

	// Run test cases
	allPassed := true

	// Counter tests
	c := Counter()
	if c() != 1 || c() != 2 || c() != 3 {
		fmt.Println("FAIL: Counter sequence")
		allPassed = false
	}

	// Independent counters
	c1, c2 := Counter(), Counter()
	c1()
	c1()
	if c2() != 1 {
		fmt.Println("FAIL: Counters should be independent")
		allPassed = false
	}

	// Accumulator tests
	a := Accumulator(0)
	if a(5) != 5 || a(5) != 10 {
		fmt.Println("FAIL: Accumulator")
		allPassed = false
	}

	// Limiter tests
	l := Limiter(2)
	n1, ok1 := l()
	n2, ok2 := l()
	n3, ok3 := l()
	if n1 != 1 || !ok1 || n2 != 2 || !ok2 || n3 != 2 || ok3 {
		fmt.Println("FAIL: Limiter")
		allPassed = false
	}

	// Memoize tests
	calls := 0
	fn := func(x int) int {
		calls++
		return x * 2
	}
	memo := Memoize(fn)
	memo(1)
	memo(2)
	memo(1) // should use cache
	memo(2) // should use cache
	if calls != 2 {
		fmt.Printf("FAIL: Memoize - expected 2 calls, got %d\n", calls)
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
