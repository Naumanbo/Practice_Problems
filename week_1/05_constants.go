package main

import "fmt"

// Problem 5: Constants
//
// Tests: Constant declaration, typed vs untyped constants, const blocks, immutability
//
// NEW CONCEPT: const declares immutable values. Can be typed or untyped.
//
// const Pi = 3.14159           // untyped constant
// const MaxSize int = 100      // typed constant
//
// const (                      // batch declaration
//     StatusOK = 200
//     StatusNotFound = 404
// )
//
// Tasks:
// 1. Declare a constant "AppName" with value "GoLearner"
// 2. Declare typed constants for "MaxRetries" (int) = 3 and "Timeout" (float64) = 30.0
// 3. Create a const block for HTTP status codes: OK=200, Created=201, BadRequest=400
// 4. Try to reassign a constant - observe the error
//
// Run: go run 05_constants.go

// Your constants here:

func main() {
	// Print your constants
	fmt.Println("Done")
}
