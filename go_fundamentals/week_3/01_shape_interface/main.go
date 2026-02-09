package main

import (
	"fmt"
	"math"
)

// Tests: Interface definition, implicit implementation, polymorphism
//
// 1. Define a Shape interface with two methods:
//    - Area() float64
//    - Perimeter() float64
//
// 2. Implement Shape for these types:
//    - Rectangle (width, height float64)
//    - Circle (radius float64)
//    - Triangle (a, b, c float64) - three side lengths
//
// 3. Implement TotalArea(shapes []Shape) float64
//    - Returns sum of all shape areas
//
// 4. Implement LargestShape(shapes []Shape) Shape
//    - Returns the shape with the largest area
//    - Return nil if slice is empty

// TODO: Define the Shape interface
type Shape interface {
}

// TODO: Implement Shape methods for Rectangle
type Rectangle struct {
	width, height float64
}

// TODO: Implement Shape methods for Circle
type Circle struct {
	radius float64
}

// TODO: Implement Shape methods for Triangle
// Hint: Use Heron's formula for area: sqrt(s*(s-a)*(s-b)*(s-c)) where s = (a+b+c)/2
type Triangle struct {
	a, b, c float64
}

// TODO: Implement TotalArea
func TotalArea(shapes []Shape) float64 {
	return 0
}

// TODO: Implement LargestShape
func LargestShape(shapes []Shape) Shape {
	return nil
}

// Helper for comparing floats
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.0001
}

func main() {
	// Test Rectangle
	fmt.Println("=== Rectangle ===")
	r := Rectangle{width: 4, height: 5}
	fmt.Printf("Area: %.2f (expected 20.00)\n", r.Area())
	fmt.Printf("Perimeter: %.2f (expected 18.00)\n", r.Perimeter())

	// Test Circle
	fmt.Println("\n=== Circle ===")
	c := Circle{radius: 3}
	fmt.Printf("Area: %.4f (expected 28.2743)\n", c.Area())
	fmt.Printf("Perimeter: %.4f (expected 18.8496)\n", c.Perimeter())

	// Test Triangle (3-4-5 right triangle)
	fmt.Println("\n=== Triangle ===")
	t := Triangle{a: 3, b: 4, c: 5}
	fmt.Printf("Area: %.2f (expected 6.00)\n", t.Area())
	fmt.Printf("Perimeter: %.2f (expected 12.00)\n", t.Perimeter())

	// Test TotalArea
	fmt.Println("\n=== TotalArea ===")
	shapes := []Shape{r, c, t}
	total := TotalArea(shapes)
	fmt.Printf("Total: %.4f (expected 54.2743)\n", total)

	// Test LargestShape
	fmt.Println("\n=== LargestShape ===")
	largest := LargestShape(shapes)
	fmt.Printf("Largest shape area: %.4f (expected Circle: 28.2743)\n", largest.Area())

	// Run test cases
	allPassed := true

	// Rectangle tests
	r1 := Rectangle{width: 10, height: 5}
	if !almostEqual(r1.Area(), 50) || !almostEqual(r1.Perimeter(), 30) {
		fmt.Println("FAIL: Rectangle 10x5")
		allPassed = false
	}

	// Circle tests
	c1 := Circle{radius: 1}
	if !almostEqual(c1.Area(), math.Pi) || !almostEqual(c1.Perimeter(), 2*math.Pi) {
		fmt.Println("FAIL: Circle radius 1")
		allPassed = false
	}

	// Triangle tests (equilateral triangle with side 2)
	t1 := Triangle{a: 2, b: 2, c: 2}
	expectedArea := math.Sqrt(3) // Equilateral triangle area = (sqrt(3)/4) * side^2
	if !almostEqual(t1.Area(), expectedArea) || !almostEqual(t1.Perimeter(), 6) {
		fmt.Println("FAIL: Equilateral triangle side 2")
		allPassed = false
	}

	// TotalArea with empty slice
	if TotalArea([]Shape{}) != 0 {
		fmt.Println("FAIL: TotalArea empty slice")
		allPassed = false
	}

	// LargestShape with empty slice
	if LargestShape([]Shape{}) != nil {
		fmt.Println("FAIL: LargestShape empty slice")
		allPassed = false
	}

	// LargestShape finds correct shape
	shapes2 := []Shape{Rectangle{2, 2}, Circle{10}, Triangle{3, 4, 5}}
	if largest := LargestShape(shapes2); !almostEqual(largest.Area(), Circle{10}.Area()) {
		fmt.Println("FAIL: LargestShape didn't find circle")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
