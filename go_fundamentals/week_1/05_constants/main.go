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
const AppName string = "GoLearner"
const AppNameUntyped = "GoLearner2"
const (
	MaxRetries int     = 3
	Timeout    float64 = 30.0
)

const (
	OK         = 200
	Created    = 201
	BadRequest = 400
)

func main() {
	// Print your constants
	// NOTE: Cannot reassign or change constants here, they stay the same as when they are initialized, immutable type
	fmt.Println(AppName)
	fmt.Println(AppNameUntyped)
	fmt.Println(MaxRetries)
	fmt.Println(Timeout)
	fmt.Println(OK)
	fmt.Println(Created)
	fmt.Println(BadRequest)

	fmt.Println("Done")
}
