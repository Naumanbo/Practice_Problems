package main

import (
	"fmt"
	"math"
)

// Tests: Interface definition, implicit implementation, polymorphism
//
// KEY TAKEAWAYS:
// - Interfaces in Go are IMPLICIT — no "implements" keyword. If a type has the right methods,
//   it satisfies the interface automatically.
// - for _, v := range slice — always use _ for index when you only need the value.
//   Single-variable range gives the INDEX (int), not the element. Recurring mistake.
// - %f is the format verb for float64. Use %.2f to control decimal precision.
// - When finding the max/min in a slice, initialize with the first element (shapes[0])
//   instead of a dummy zero value — avoids coupling to a specific type.
// - A []Shape can hold Rectangle, Circle, Triangle — this is polymorphism via interfaces.
//   You call shape.Area() without knowing the concrete type.
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
	Area() float64
	Perimeter() float64
}

// TODO: Implement Shape methods for Rectangle
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
} // TODO: implement
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
} // TODO: implement

// TODO: Implement Shape methods for Circle
type Circle struct {
	radius float64
}

func (c Circle) Area() float64      { return math.Pi * c.radius * c.radius } // TODO: implement
func (c Circle) Perimeter() float64 { return math.Pi * 2 * c.radius }        // TODO: implement

// TODO: Implement Shape methods for Triangle
// Hint: Use Heron's formula for area: sqrt(s*(s-a)*(s-b)*(s-c)) where s = (a+b+c)/2
type Triangle struct {
	a, b, c float64
}

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2

	return math.Sqrt((s * (s - t.a) * (s - t.b) * (s - t.c))) // TODO: implement
}
func (t Triangle) Perimeter() float64 { return t.a + t.b + t.c } // TODO: implement

// TODO: Implement TotalArea
func TotalArea(shapes []Shape) float64 {
	var totalArea float64 = 0.0
	for _, shape := range shapes {
		totalArea += shape.Area()
	}
	return totalArea
}

// TODO: Implement LargestShape
func LargestShape(shapes []Shape) Shape {
	if len(shapes) == 0 {
		return nil
	}

	var largestShape Shape = shapes[0]

	for _, shape := range shapes {
		if shape.Area() > largestShape.Area() {
			largestShape = shape
		}
	}

	// fmt.Printf("Largest Area: %f\n", largestShape.Area())

	return largestShape
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
	bigCircle := Circle{10}
	if largest2 := LargestShape(shapes2); !almostEqual(largest2.Area(), bigCircle.Area()) {
		fmt.Println("FAIL: LargestShape didn't find circle")
		allPassed = false
	}

	// === Additional Rectangle tests ===
	// square
	sq := Rectangle{width: 5, height: 5}
	if !almostEqual(sq.Area(), 25) || !almostEqual(sq.Perimeter(), 20) {
		fmt.Println("FAIL: Rectangle square 5x5")
		allPassed = false
	}
	// very thin rectangle
	thin := Rectangle{width: 100, height: 0.1}
	if !almostEqual(thin.Area(), 10) || !almostEqual(thin.Perimeter(), 200.2) {
		fmt.Println("FAIL: Rectangle thin 100x0.1")
		allPassed = false
	}
	// unit rectangle
	unit := Rectangle{width: 1, height: 1}
	if !almostEqual(unit.Area(), 1) || !almostEqual(unit.Perimeter(), 4) {
		fmt.Println("FAIL: Rectangle unit 1x1")
		allPassed = false
	}

	// === Additional Circle tests ===
	// zero radius
	zeroCircle := Circle{radius: 0}
	if !almostEqual(zeroCircle.Area(), 0) || !almostEqual(zeroCircle.Perimeter(), 0) {
		fmt.Println("FAIL: Circle radius 0")
		allPassed = false
	}
	// large radius
	bigC := Circle{radius: 100}
	if !almostEqual(bigC.Area(), math.Pi*10000) || !almostEqual(bigC.Perimeter(), 200*math.Pi) {
		fmt.Println("FAIL: Circle radius 100")
		allPassed = false
	}

	// === Additional Triangle tests ===
	// isosceles triangle
	iso := Triangle{a: 5, b: 5, c: 6}
	isoS := (5.0 + 5.0 + 6.0) / 2
	isoExpected := math.Sqrt(isoS * (isoS - 5) * (isoS - 5) * (isoS - 6))
	if !almostEqual(iso.Area(), isoExpected) || !almostEqual(iso.Perimeter(), 16) {
		fmt.Println("FAIL: Isosceles triangle 5-5-6")
		allPassed = false
	}

	// === Additional TotalArea tests ===
	// single shape
	if !almostEqual(TotalArea([]Shape{Rectangle{3, 4}}), 12) {
		fmt.Println("FAIL: TotalArea single rectangle")
		allPassed = false
	}
	// all same type
	allCircles := []Shape{Circle{1}, Circle{2}, Circle{3}}
	expectedTotal := math.Pi*1 + math.Pi*4 + math.Pi*9
	if !almostEqual(TotalArea(allCircles), expectedTotal) {
		fmt.Println("FAIL: TotalArea all circles")
		allPassed = false
	}

	// === Additional LargestShape tests ===
	// single shape
	singleShape := []Shape{Rectangle{2, 3}}
	if ls := LargestShape(singleShape); !almostEqual(ls.Area(), 6) {
		fmt.Println("FAIL: LargestShape single shape")
		allPassed = false
	}
	// all same area
	sameArea := []Shape{Rectangle{2, 6}, Rectangle{3, 4}, Rectangle{1, 12}}
	if ls := LargestShape(sameArea); ls == nil {
		fmt.Println("FAIL: LargestShape all same area should return one")
		allPassed = false
	}
	// rectangle is largest
	shapes3 := []Shape{Circle{1}, Rectangle{100, 100}, Triangle{3, 4, 5}}
	if ls := LargestShape(shapes3); !almostEqual(ls.Area(), 10000) {
		fmt.Println("FAIL: LargestShape rectangle largest")
		allPassed = false
	}
	// triangle is largest
	shapes4 := []Shape{Rectangle{1, 1}, Circle{0.5}, Triangle{10, 10, 10}}
	triS := (10.0 + 10.0 + 10.0) / 2
	triArea := math.Sqrt(triS * (triS - 10) * (triS - 10) * (triS - 10))
	if ls := LargestShape(shapes4); !almostEqual(ls.Area(), triArea) {
		fmt.Println("FAIL: LargestShape triangle largest")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
