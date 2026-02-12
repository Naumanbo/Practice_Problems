package main

import "fmt"

// Problem 3: Functions
//
// KEY TAKEAWAYS:
// - Named return values (e.g., func f() (min int, max int)) allow naked returns.
// - DON'T name return values after built-in functions: naming a return "min" or "max"
//   shadows Go's built-in min()/max(), causing compile errors when you try to call them.
// - Renamed to minNumber/maxNumber to avoid shadowing.
// - Consider edge cases like division by zero even when not explicitly asked.
//
// Tests: Function syntax, multiple return values, naked returns with named return values
//
// Tasks:
// 1. Write a function "multiply" that takes two ints and returns their product
// 2. Write a function "divmod" that takes two ints and returns both quotient AND remainder
// 3. Write a function "minMax" that takes three ints and returns (minNumber, maxNumber) using naked return
//
// Run: go run 03_functions.go

// Your functions here:
func multiply(x, y int) int {
	return (x * y)
}

func divmod(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}

func minMax(x, y, z int) (minNumber int, maxNumber int) {
	minNumber = min(x, y, z)
	maxNumber = max(x, y, z)
	return
}

func main() {
	// Test multiply
	fmt.Println("3 * 4 =", multiply(3, 4))

	// Test divmod
	q, r := divmod(17, 5)
	fmt.Println("17 / 5 = quotient:", q, "remainder:", r)

	// Test minMax
	min, max := minMax(5, 2, 8)
	fmt.Println("min:", min, "max:", max)
}
