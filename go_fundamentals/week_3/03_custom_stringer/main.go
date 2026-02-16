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

	// === Temperature tests ===
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
	// zero temperature
	tc3 := Temperature{0, 'C'}
	if tc3.String() != "0°C" {
		fmt.Println("FAIL: Temperature 0C")
		allPassed = false
	}
	// negative temperature
	tc4 := Temperature{-40, 'F'}
	if tc4.String() != "-40°F" {
		fmt.Println("FAIL: Temperature -40F")
		allPassed = false
	}
	// decimal temperature
	tc5 := Temperature{72.5, 'F'}
	if tc5.String() != "72.5°F" {
		fmt.Println("FAIL: Temperature 72.5F")
		allPassed = false
	}

	// === Duration tests ===
	if Duration(0).String() != "0s" {
		fmt.Println("FAIL: Duration 0")
		allPassed = false
	}
	if Duration(1).String() != "1s" {
		fmt.Println("FAIL: Duration 1s")
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
	if Duration(61).String() != "1m 1s" {
		fmt.Println("FAIL: Duration 61s")
		allPassed = false
	}
	if Duration(90).String() != "1m 30s" {
		fmt.Println("FAIL: Duration 90s")
		allPassed = false
	}
	if Duration(3600).String() != "1h" {
		fmt.Println("FAIL: Duration 3600s = 1h")
		allPassed = false
	}
	if Duration(3601).String() != "1h 1s" {
		fmt.Println("FAIL: Duration 3601s = 1h 1s")
		allPassed = false
	}
	if Duration(3660).String() != "1h 1m" {
		fmt.Println("FAIL: Duration 3660s = 1h 1m")
		allPassed = false
	}
	if Duration(3661).String() != "1h 1m 1s" {
		fmt.Println("FAIL: Duration 3661s")
		allPassed = false
	}
	if Duration(7200).String() != "2h" {
		fmt.Println("FAIL: Duration 7200s = 2h")
		allPassed = false
	}

	// === RGB tests ===
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
	// single digit hex values need zero-padding
	lowVal := RGB{1, 2, 3}
	if lowVal.String() != "#010203" {
		fmt.Println("FAIL: RGB low values need zero-padding")
		allPassed = false
	}
	red := RGB{255, 0, 0}
	if red.String() != "#FF0000" {
		fmt.Println("FAIL: RGB red")
		allPassed = false
	}
	// mixed values
	mixed := RGB{16, 32, 128}
	if mixed.String() != "#102080" {
		fmt.Println("FAIL: RGB mixed")
		allPassed = false
	}

	// === Person tests ===
	bob := Person{"Bob", 25, "Developer"}
	if bob.String() != "Bob (25) - Developer" {
		fmt.Println("FAIL: Person Bob")
		allPassed = false
	}
	// age zero
	baby := Person{"Baby", 0, "None"}
	if baby.String() != "Baby (0) - None" {
		fmt.Println("FAIL: Person age 0")
		allPassed = false
	}

	// === Matrix tests ===
	m2 := Matrix{{1, 2}, {3, 4}}
	expected := "[1, 2]\n[3, 4]"
	if m2.String() != expected {
		fmt.Printf("FAIL: Matrix 2x2, got %q, expected %q\n", m2.String(), expected)
		allPassed = false
	}
	// single row
	m3 := Matrix{{5, 10, 15}}
	if m3.String() != "[5, 10, 15]" {
		fmt.Printf("FAIL: Matrix single row, got %q\n", m3.String())
		allPassed = false
	}
	// single element
	m4 := Matrix{{42}}
	if m4.String() != "[42]" {
		fmt.Printf("FAIL: Matrix single element, got %q\n", m4.String())
		allPassed = false
	}
	// 3x3
	m5 := Matrix{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected5 := "[1, 2, 3]\n[4, 5, 6]\n[7, 8, 9]"
	if m5.String() != expected5 {
		fmt.Printf("FAIL: Matrix 3x3, got %q\n", m5.String())
		allPassed = false
	}
	// empty matrix
	emptyMatrix := Matrix{}
	if emptyMatrix.String() != "" {
		fmt.Println("FAIL: Empty matrix")
		allPassed = false
	}
	// negative values
	m6 := Matrix{{-1, 0}, {0, -1}}
	expected6 := "[-1, 0]\n[0, -1]"
	if m6.String() != expected6 {
		fmt.Printf("FAIL: Matrix negatives, got %q\n", m6.String())
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: strings package imported for your convenience
var _ = strings.Join
