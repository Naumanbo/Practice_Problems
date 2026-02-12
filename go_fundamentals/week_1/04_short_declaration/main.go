package main

import "fmt"

// Problem 4: Short Declaration (:=) vs var
//
// KEY TAKEAWAYS:
// - := declares AND initializes. Only works inside functions, not at package level.
// - var is required at package level or when you want to declare without initializing.
// - := with an existing variable on the left side only works if at least one variable is new.
//
// Tests: Short declaration operator, var vs :=, scope rules (package vs function level)
//
// NEW CONCEPT: := is shorthand for declare + initialize. Only works inside functions.
//
// var x int = 10    // explicit type
// var x = 10        // type inferred
// x := 10           // short declaration (type inferred, inside function only)
//
// Tasks:
// 1. Use var to declare a package-level variable "version" = "1.0.0"
// 2. Inside main, use := to create "count" = 42
// 3. Try using := at package level - observe the compiler error, then fix it
// 4. Demonstrate re-assignment: you can't use := twice for the same variable
//
// Run: go run 04_short_declaration.go

// Package level - must use var
// version := "1.0.0"  // This won't work - fix it
var version string = "1.0.0"

func main() {
	// Your code here:
	count := 42
	// Experiment: What happens if you do this?
	// x := 5
	// x := 10  // Try this, observe error, then fix
	fmt.Println(count)
	count = 56 // cannot use ":=" again if variable is already initialized
	fmt.Println(count)
	fmt.Println(version)
	fmt.Println("Done")
}
