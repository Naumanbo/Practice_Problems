package main

import "fmt"

// Exercise: defer, panic, and recover
//
// Practice Go's error flow control mechanisms.
//
// What to build:
// - SafeDivide: divides two floats, uses recover to catch any panic and return an error instead
// - RunWithCleanup: runs a function, uses defer to always print "cleanup done" after, even on panic
// - MustPositive: returns the number if positive, panics with a message if <= 0
// - SafeRun: runs a function and returns any panic message as an error (recover-based wrapper)
//
// Read the tests in main() to understand exact function signatures and expected behavior.

// === WRITE YOUR CODE BELOW ===

// === END YOUR CODE ===

func main() {
	allPassed := true

	// === SafeDivide tests ===
	result, err := SafeDivide(10, 2)
	if err != nil || result != 5.0 {
		fmt.Println("FAIL: SafeDivide(10, 2) should return 5.0, nil")
		allPassed = false
	}

	result, err = SafeDivide(7, 0)
	if err == nil {
		fmt.Println("FAIL: SafeDivide(7, 0) should return an error")
		allPassed = false
	}
	_ = result

	result, err = SafeDivide(-10, 2)
	if err != nil || result != -5.0 {
		fmt.Println("FAIL: SafeDivide(-10, 2) should return -5.0, nil")
		allPassed = false
	}

	result, err = SafeDivide(0, 5)
	if err != nil || result != 0.0 {
		fmt.Println("FAIL: SafeDivide(0, 5) should return 0.0, nil")
		allPassed = false
	}

	// === MustPositive tests ===
	if MustPositive(5) != 5 {
		fmt.Println("FAIL: MustPositive(5) should return 5")
		allPassed = false
	}

	if MustPositive(1) != 1 {
		fmt.Println("FAIL: MustPositive(1) should return 1")
		allPassed = false
	}

	// MustPositive(-1) should panic — verify via SafeRun
	err = SafeRun(func() { MustPositive(-1) })
	if err == nil {
		fmt.Println("FAIL: MustPositive(-1) should panic")
		allPassed = false
	}

	err = SafeRun(func() { MustPositive(0) })
	if err == nil {
		fmt.Println("FAIL: MustPositive(0) should panic")
		allPassed = false
	}

	// === SafeRun tests ===
	err = SafeRun(func() {})
	if err != nil {
		fmt.Println("FAIL: SafeRun with no panic should return nil")
		allPassed = false
	}

	err = SafeRun(func() { panic("something went wrong") })
	if err == nil {
		fmt.Println("FAIL: SafeRun should catch panic and return error")
		allPassed = false
	}

	// === RunWithCleanup test ===
	// Should print "cleanup done" even when the function panics
	fmt.Println("--- RunWithCleanup (expect 'cleanup done' to always print) ---")
	SafeRun(func() {
		RunWithCleanup(func() {
			fmt.Println("  doing work...")
			panic("mid-work panic")
		})
	})
	RunWithCleanup(func() {
		fmt.Println("  doing work without panic...")
	})

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
