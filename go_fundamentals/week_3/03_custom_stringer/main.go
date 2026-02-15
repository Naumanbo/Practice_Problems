package main

import (
	"fmt"
	"strings"
)

// Tests: fmt.Stringer interface, String() method, formatting
//
// Implement the fmt.Stringer interface for each type below.
// The Stringer interface is: type Stringer interface { String() string }
//
// 1. Temperature - can be Celsius or Fahrenheit
//    - Format: "72.5°F" or "22.5°C"
//
// 2. Duration - represents time in seconds
//    - Format as "Xh Ym Zs", omitting zero components
//    - Examples: "1h 30m 0s" -> "1h 30m", "90s" -> "1m 30s", "3600s" -> "1h"
//
// 3. RGB - represents a color
//    - Format: "#RRGGBB" (hex, uppercase)
//    - Example: RGB{255, 128, 0} -> "#FF8000"
//
// 4. Person - has name, age, and occupation
//    - Format: "Name (Age) - Occupation"
//    - Example: "Alice (30) - Engineer"
//
// 5. Matrix - 2D slice of ints
//    - Format as rows with brackets
//    - Example: [[1,2],[3,4]] -> "[1, 2]\n[3, 4]"

// Temperature represents a temperature value with unit
type Temperature struct {
	Value float64
	Unit  rune // 'C' for Celsius, 'F' for Fahrenheit
}

// TODO: Implement String() for Temperature
func (t Temperature) String() string { return "" } // TODO: implement

// Duration represents a time duration in seconds
type Duration int // seconds

// TODO: Implement String() for Duration
func (d Duration) String() string { return "" } // TODO: implement

// RGB represents a color
type RGB struct {
	R, G, B uint8
}

// TODO: Implement String() for RGB
func (c RGB) String() string { return "" } // TODO: implement

// Person represents a person
type Person struct {
	Name       string
	Age        int
	Occupation string
}

// TODO: Implement String() for Person
func (p Person) String() string { return "" } // TODO: implement

// Matrix represents a 2D grid of integers
type Matrix [][]int

// TODO: Implement String() for Matrix
func (m Matrix) String() string { return "" } // TODO: implement

func main() {
	// Test Temperature
	fmt.Println("=== Temperature ===")
	t1 := Temperature{72.5, 'F'}
	t2 := Temperature{22.5, 'C'}
	t3 := Temperature{0, 'C'}
	fmt.Println(t1) // 72.5°F
	fmt.Println(t2) // 22.5°C
	fmt.Println(t3) // 0°C

	// Test Duration
	fmt.Println("\n=== Duration ===")
	fmt.Println(Duration(3661)) // 1h 1m 1s
	fmt.Println(Duration(3600)) // 1h
	fmt.Println(Duration(90))   // 1m 30s
	fmt.Println(Duration(45))   // 45s
	fmt.Println(Duration(0))    // 0s

	// Test RGB
	fmt.Println("\n=== RGB ===")
	fmt.Println(RGB{255, 255, 255}) // #FFFFFF
	fmt.Println(RGB{0, 0, 0})       // #000000
	fmt.Println(RGB{255, 128, 0})   // #FF8000
	fmt.Println(RGB{18, 52, 86})    // #123456

	// Test Person
	fmt.Println("\n=== Person ===")
	p := Person{"Alice", 30, "Engineer"}
	fmt.Println(p) // Alice (30) - Engineer

	// Test Matrix
	fmt.Println("\n=== Matrix ===")
	m := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(m)
	// [1, 2, 3]
	// [4, 5, 6]
	// [7, 8, 9]

	// Run test cases
	allPassed := true

	// Temperature tests
	tc1 := Temperature{100, 'C'}
	if tc1.String() != "100°C" {
		fmt.Println("FAIL: Temperature 100C")
		allPassed = false
	}
	tc2 := Temperature{32, 'F'}
	if tc2.String() != "32°F" {
		fmt.Println("FAIL: Temperature 32F")
		allPassed = false
	}

	// Duration tests
	if Duration(0).String() != "0s" {
		fmt.Println("FAIL: Duration 0")
		allPassed = false
	}
	if Duration(59).String() != "59s" {
		fmt.Println("FAIL: Duration 59s")
		allPassed = false
	}
	if Duration(60).String() != "1m" {
		fmt.Println("FAIL: Duration 60s = 1m")
		allPassed = false
	}
	if Duration(3600).String() != "1h" {
		fmt.Println("FAIL: Duration 3600s = 1h")
		allPassed = false
	}
	if Duration(3661).String() != "1h 1m 1s" {
		fmt.Println("FAIL: Duration 3661s")
		allPassed = false
	}

	// RGB tests
	black := RGB{0, 0, 0}
	if black.String() != "#000000" {
		fmt.Println("FAIL: RGB black")
		allPassed = false
	}
	white := RGB{255, 255, 255}
	if white.String() != "#FFFFFF" {
		fmt.Println("FAIL: RGB white")
		allPassed = false
	}
	abc := RGB{171, 205, 239}
	if abc.String() != "#ABCDEF" {
		fmt.Println("FAIL: RGB ABCDEF")
		allPassed = false
	}

	// Person tests
	bob := Person{"Bob", 25, "Developer"}
	if bob.String() != "Bob (25) - Developer" {
		fmt.Println("FAIL: Person Bob")
		allPassed = false
	}

	// Matrix tests
	m2 := Matrix{{1, 2}, {3, 4}}
	expected := "[1, 2]\n[3, 4]"
	if m2.String() != expected {
		fmt.Printf("FAIL: Matrix 2x2, got %q, expected %q\n", m2.String(), expected)
		allPassed = false
	}

	// Empty matrix
	emptyMatrix := Matrix{}
	if emptyMatrix.String() != "" {
		fmt.Println("FAIL: Empty matrix")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: strings package imported for your convenience
var _ = strings.Join
