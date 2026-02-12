package main

import "fmt"

// Problem 2: Variables and Type Declaration
//
// KEY TAKEAWAYS:
// - Go places the type AFTER the variable name: var x int (not int x).
// - Zero values: int=0, string="", bool=false, float64=0.0. Uninitialized vars are never garbage.
// - Use camelCase in Go, not snake_case (Python habit to break).
//
// Tests: Variable declaration syntax, type placement after name, zero values
//
// Tasks:
// 1. Declare two variables x and y of type int using the short form (one line)
// 2. Declare a string variable "name" using var syntax
// 3. Use := to declare and initialize "age" to 25
// 4. Declare a bool "isActive" without initialization (observe zero value)
//
// Run: go run 02_variables.go

func main() {
	// Your code here:
	var x, y int = 50, 43
	var name string = "Nauman"
	age := 25
	var isActive bool // zero value is false
	// Uncomment and fix:
	fmt.Println("x:", x, "y:", y)
	fmt.Println("name:", name)
	fmt.Println("age:", age)
	fmt.Println("isActive:", isActive)
}
