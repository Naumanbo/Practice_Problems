package main

// Problem 1: Imports and Exported Names
//
// Tests: Batch import syntax, accessing exported names from packages
//
// Tasks:
// 1. Import "fmt" and "math" using batch import syntax
// 2. Print the value of Pi from the math package
// 3. Print the square root of 144 using math.Sqrt
//
// Run: go run 01_imports_and_exports.go

// Your code here:
import (
	"fmt"
	"math"
)

func main() {
	// Print Pi and Sqrt(144)
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(144))
}
