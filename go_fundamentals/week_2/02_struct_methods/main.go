package main

import (
	"fmt"
	"math"
)

// Tests: Structs, methods, pointer receivers vs value receivers
//
// Implement methods on the Rectangle struct:
// - Area() returns the area (width * height)
// - Perimeter() returns the perimeter (2*width + 2*height)
// - Scale(factor) scales both dimensions by factor (use pointer receiver)
// - Diagonal() returns the diagonal length using Pythagorean theorem

type Rectangle struct {
	Width  float64
	Height float64
}

// TODO: Implement Area() - value receiver
func (r Rectangle) Area() float64 {
	return 0
}

// TODO: Implement Perimeter() - value receiver
func (r Rectangle) Perimeter() float64 {
	return 0
}

// TODO: Implement Scale() - pointer receiver (modifies the rectangle)
func (r *Rectangle) Scale(factor float64) {
}

// TODO: Implement Diagonal() - value receiver
func (r Rectangle) Diagonal() float64 {
	return 0
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}

	fmt.Printf("Rectangle: %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area())           // Expected: 12.00
	fmt.Printf("Perimeter: %.2f\n", rect.Perimeter()) // Expected: 14.00
	fmt.Printf("Diagonal: %.2f\n", rect.Diagonal())   // Expected: 5.00

	rect.Scale(2)
	fmt.Printf("\nAfter Scale(2): %+v\n", rect)
	fmt.Printf("Area: %.2f\n", rect.Area()) // Expected: 48.00

	// Test cases
	tests := []struct {
		w, h          float64
		wantArea      float64
		wantPerimeter float64
		wantDiagonal  float64
	}{
		{3, 4, 12, 14, 5},
		{5, 5, 25, 20, math.Sqrt(50)},
		{1, 1, 1, 4, math.Sqrt(2)},
		{10, 0, 0, 20, 10},
	}

	allPassed := true
	for _, tc := range tests {
		r := Rectangle{Width: tc.w, Height: tc.h}
		if r.Area() != tc.wantArea {
			fmt.Printf("FAIL: Area(%v, %v) = %v, want %v\n", tc.w, tc.h, r.Area(), tc.wantArea)
			allPassed = false
		}
		if r.Perimeter() != tc.wantPerimeter {
			fmt.Printf("FAIL: Perimeter(%v, %v) = %v, want %v\n", tc.w, tc.h, r.Perimeter(), tc.wantPerimeter)
			allPassed = false
		}
		if math.Abs(r.Diagonal()-tc.wantDiagonal) > 0.0001 {
			fmt.Printf("FAIL: Diagonal(%v, %v) = %v, want %v\n", tc.w, tc.h, r.Diagonal(), tc.wantDiagonal)
			allPassed = false
		}
	}

	// Test Scale with pointer receiver
	r := Rectangle{Width: 2, Height: 3}
	r.Scale(3)
	if r.Width != 6 || r.Height != 9 {
		fmt.Printf("FAIL: Scale(3) on {2,3} = {%v,%v}, want {6,9}\n", r.Width, r.Height)
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
