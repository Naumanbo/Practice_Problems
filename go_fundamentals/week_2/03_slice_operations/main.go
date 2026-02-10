package main

import "fmt"

// Tests: Slice internals - capacity, aliasing, append reallocation
//
// Go slices have TRAPS that don't exist in Python lists or C++ vectors.
// This problem forces you to understand how slices work under the hood.
//
// A slice is a struct: { pointer to array, length, capacity }
// - append() may or may not allocate a new backing array
// - Multiple slices can share the same backing array (aliasing)
// - Reslicing doesn't copy data
//
// 1. Implement RemoveIndex(s []int, i int) []int
//    - Remove element at index i, preserving order
//    - Must NOT modify the original slice's visible elements
//    - Hint: think about whether append modifies the underlying array
//
// 2. Implement InsertAt(s []int, i int, val int) []int
//    - Insert val at index i, shifting elements right
//    - Return the new slice
//
// 3. Implement Compact(s []int) []int
//    - Remove consecutive duplicates: [1,1,2,2,2,3,1,1] -> [1,2,3,1]
//    - Return a new slice
//
// 4. Implement Chunk(s []int, size int) [][]int
//    - Split slice into chunks of given size
//    - Last chunk may be smaller
//    - [1,2,3,4,5] with size 2 -> [[1,2], [3,4], [5]]
//
// 5. Implement RotateLeft(s []int, k int) []int
//    - Rotate slice left by k positions
//    - [1,2,3,4,5] rotated by 2 -> [3,4,5,1,2]
//    - Must handle k > len(s) and k < 0
//    - Must NOT modify the original slice

// TODO: Implement RemoveIndex
func RemoveIndex(s []int, i int) []int {
	// avoid using append since this writes into the original array
	len := len(s[:i]) + len(s[i+1:])
	combined := make([]int, 0, len)
	combined = append(combined, s[i:]...)
	combined = append(combined, s[i+1:]...)
	return combined
}

// TODO: Implement InsertAt
func InsertAt(s []int, i int, val int) []int {
	combined := make([]int, 0, len(s[:i])+len(s[i:])+1)
	combined = append(combined, s[:i]...)
	combined = append(combined, val)
	combined = append(combined, s[i:]...)
	return combined
}

// TODO: Implement Compact
//   - Remove consecutive duplicates: [1,1,2,2,2,3,1,1] -> [1,2,3,1]
//   - Return a new slice
func Compact(s []int) []int {
	newSlice := []int{}
	// for loop iterating over s to look for consecutive duplicates to build new_slice without them
	if len(s) > 0 {
		newSlice = append(newSlice, s[0])
		for _, v := range s {
			if v != newSlice[len(newSlice)-1] {
				newSlice = append(newSlice, v)
			}
		}
	}
	return newSlice
}

// TODO: Implement Chunk
//   - Split slice into chunks of given size
//   - Last chunk may be smaller
//   - [1,2,3,4,5] with size 2 -> [[1,2], [3,4], [5]]
func Chunk(s []int, size int) [][]int {
	newSlice := [][]int{}
	if len(s) > 0 {
		for i := 0; i < len(s); i += size {
			if i+size > len(s) {
				newSlice = append(newSlice, s[i:])
			} else {
				newSlice = append(newSlice, s[i:i+size])
			}
		}

	}
	return newSlice

}

// TODO: Implement RotateLeft
//   - Rotate slice left by k positions
//   - [1,2,3,4,5] rotated by 2 -> [3,4,5,1,2]
//   - Must handle k > len(s) and k < 0
//   - Must NOT modify the original slice
func RotateLeft(s []int, k int) []int {
	// k % size = rotation value if k > len(s)

	return nil
}

func main() {
	// Test RemoveIndex
	fmt.Println("=== RemoveIndex ===")
	original := []int{1, 2, 3, 4, 5}
	removed := RemoveIndex(original, 2)
	fmt.Println("Removed index 2:", removed) // [1 2 4 5]
	fmt.Println("Original:", original[:5])   // Should still be [1 2 3 4 5]

	fmt.Println("Remove first:", RemoveIndex([]int{1, 2, 3}, 0)) // [2 3]
	fmt.Println("Remove last:", RemoveIndex([]int{1, 2, 3}, 2))  // [1 2]

	// Test InsertAt
	fmt.Println("\n=== InsertAt ===")
	fmt.Println(InsertAt([]int{1, 3, 4}, 1, 2)) // [1 2 3 4]
	fmt.Println(InsertAt([]int{2, 3}, 0, 1))    // [1 2 3]
	fmt.Println(InsertAt([]int{1, 2}, 2, 3))    // [1 2 3]

	// Test Compact
	fmt.Println("\n=== Compact ===")
	fmt.Println(Compact([]int{1, 1, 2, 2, 2, 3, 1, 1})) // [1 2 3 1]
	fmt.Println(Compact([]int{1, 1, 1, 1}))             // [1]
	fmt.Println(Compact([]int{1, 2, 3}))                // [1 2 3]
	fmt.Println(Compact([]int{}))                       // []

	// Test Chunk
	fmt.Println("\n=== Chunk ===")
	fmt.Println(Chunk([]int{1, 2, 3, 4, 5}, 2)) // [[1 2] [3 4] [5]]
	fmt.Println(Chunk([]int{1, 2, 3, 4}, 2))    // [[1 2] [3 4]]
	fmt.Println(Chunk([]int{1, 2, 3}, 5))       // [[1 2 3]]

	// Test RotateLeft
	fmt.Println("\n=== RotateLeft ===")
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, 2))  // [3 4 5 1 2]
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, 0))  // [1 2 3 4 5]
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, 7))  // [3 4 5 1 2] (7 % 5 = 2)
	fmt.Println(RotateLeft([]int{1, 2, 3, 4, 5}, -1)) // [5 1 2 3 4] (rotate right by 1)

	// Run test cases
	allPassed := true

	// RemoveIndex - doesn't corrupt original
	orig := []int{10, 20, 30, 40, 50}
	origCopy := make([]int, len(orig))
	copy(origCopy, orig)
	RemoveIndex(orig, 1)
	for i := range orig {
		if orig[i] != origCopy[i] {
			fmt.Println("FAIL: RemoveIndex modified original slice")
			allPassed = false
			break
		}
	}

	// RemoveIndex single element
	if fmt.Sprint(RemoveIndex([]int{42}, 0)) != "[]" {
		fmt.Println("FAIL: RemoveIndex single element")
		allPassed = false
	}

	// InsertAt at beginning
	if fmt.Sprint(InsertAt([]int{2, 3}, 0, 1)) != "[1 2 3]" {
		fmt.Println("FAIL: InsertAt beginning")
		allPassed = false
	}

	// InsertAt at end
	if fmt.Sprint(InsertAt([]int{1, 2}, 2, 3)) != "[1 2 3]" {
		fmt.Println("FAIL: InsertAt end")
		allPassed = false
	}

	// InsertAt into empty slice
	if fmt.Sprint(InsertAt([]int{}, 0, 1)) != "[1]" {
		fmt.Println("FAIL: InsertAt empty")
		allPassed = false
	}

	// Compact single element
	if fmt.Sprint(Compact([]int{5})) != "[5]" {
		fmt.Println("FAIL: Compact single")
		allPassed = false
	}

	// Compact all same
	if fmt.Sprint(Compact([]int{7, 7, 7, 7, 7})) != "[7]" {
		fmt.Println("FAIL: Compact all same")
		allPassed = false
	}

	// Compact no duplicates
	if fmt.Sprint(Compact([]int{1, 2, 3, 4})) != "[1 2 3 4]" {
		fmt.Println("FAIL: Compact no dups")
		allPassed = false
	}

	// Chunk empty
	if len(Chunk([]int{}, 3)) != 0 {
		fmt.Println("FAIL: Chunk empty")
		allPassed = false
	}

	// Chunk size 1
	chunks := Chunk([]int{1, 2, 3}, 1)
	if len(chunks) != 3 {
		fmt.Println("FAIL: Chunk size 1")
		allPassed = false
	}

	// Chunk exact fit
	chunks = Chunk([]int{1, 2, 3, 4}, 2)
	if len(chunks) != 2 || fmt.Sprint(chunks[0]) != "[1 2]" || fmt.Sprint(chunks[1]) != "[3 4]" {
		fmt.Println("FAIL: Chunk exact fit")
		allPassed = false
	}

	// RotateLeft by length (no change)
	if fmt.Sprint(RotateLeft([]int{1, 2, 3}, 3)) != "[1 2 3]" {
		fmt.Println("FAIL: RotateLeft by length")
		allPassed = false
	}

	// RotateLeft negative (rotate right)
	if fmt.Sprint(RotateLeft([]int{1, 2, 3, 4}, -1)) != "[4 1 2 3]" {
		fmt.Println("FAIL: RotateLeft negative")
		allPassed = false
	}

	// RotateLeft empty
	if fmt.Sprint(RotateLeft([]int{}, 5)) != "[]" {
		fmt.Println("FAIL: RotateLeft empty")
		allPassed = false
	}

	// RotateLeft doesn't modify original
	rotOrig := []int{1, 2, 3}
	RotateLeft(rotOrig, 1)
	if fmt.Sprint(rotOrig) != "[1 2 3]" {
		fmt.Println("FAIL: RotateLeft modified original")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
