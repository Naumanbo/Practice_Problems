package main

import "fmt"

// Tests: Pointers, dereferencing, pass-by-reference
//
// Implement swap() to swap two integers using pointers.
// This reinforces pointer syntax (*T, &, dereferencing).

func swap(a, b *int) {
	// TODO: Swap the values that a and b point to
}

func main() {
	x, y := 10, 20
	fmt.Printf("Before: x=%d, y=%d\n", x, y)

	swap(&x, &y)

	fmt.Printf("After:  x=%d, y=%d\n", x, y)

	// Expected output:
	// Before: x=10, y=20
	// After:  x=20, y=10

	// Test cases
	tests := []struct {
		a, b         int
		wantA, wantB int
	}{
		{1, 2, 2, 1},
		{-5, 5, 5, -5},
		{0, 100, 100, 0},
		{42, 42, 42, 42}, // same values
	}

	allPassed := true
	for _, tc := range tests {
		a, b := tc.a, tc.b
		swap(&a, &b)
		if a != tc.wantA || b != tc.wantB {
			fmt.Printf("FAIL: swap(%d, %d) = (%d, %d), want (%d, %d)\n",
				tc.a, tc.b, a, b, tc.wantA, tc.wantB)
			allPassed = false
		}
	}

	if allPassed {
		fmt.Println("All tests passed!")
	}
}
