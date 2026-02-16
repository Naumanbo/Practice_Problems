package main

import (
	"fmt"
)

// Tests: Generic types, methods on generic structs, pointer receivers
//
// Implement a generic Stack data structure.
//
// 1. Define Stack[T any] struct
//    - Use a slice internally to store elements
//
// 2. Implement these methods (all pointer receivers):
//    - Push(val T)          - add element to top
//    - Pop() (T, bool)      - remove and return top element; bool = false if empty
//    - Peek() (T, bool)     - return top element without removing; bool = false if empty
//    - Size() int           - return number of elements
//    - IsEmpty() bool       - return true if stack has no elements
//
// 3. Implement a standalone function:
//    - Reverse[T any](slice []T) []T
//    - Uses a Stack internally to reverse a slice

// TODO: Add a field to hold elements (hint: use a slice of T)
type Stack[T any] struct {
	elements []T
}

// TODO: Implement Push
func (s *Stack[T]) Push(val T) {
}

// TODO: Implement Pop
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	return zero, false
}

// TODO: Implement Peek
func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	return zero, false
}

// TODO: Implement Size
func (s *Stack[T]) Size() int {
	return 0
}

// TODO: Implement IsEmpty
func (s *Stack[T]) IsEmpty() bool {
	return true
}

// TODO: Implement Reverse using a Stack
func Reverse[T any](slice []T) []T {
	return nil
}

func main() {
	// Test int stack
	fmt.Println("=== Int Stack ===")
	s := &Stack[int]{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Size:", s.Size())         // 3
	fmt.Println("IsEmpty:", s.IsEmpty())   // false

	if val, ok := s.Peek(); ok {
		fmt.Println("Peek:", val) // 30
	}
	fmt.Println("Size after peek:", s.Size()) // 3 (peek doesn't remove)

	if val, ok := s.Pop(); ok {
		fmt.Println("Pop:", val) // 30
	}
	if val, ok := s.Pop(); ok {
		fmt.Println("Pop:", val) // 20
	}
	if val, ok := s.Pop(); ok {
		fmt.Println("Pop:", val) // 10
	}
	fmt.Println("IsEmpty:", s.IsEmpty()) // true

	// Pop from empty stack
	if _, ok := s.Pop(); !ok {
		fmt.Println("Pop from empty: ok=false (correct)")
	}

	// Test string stack
	fmt.Println("\n=== String Stack ===")
	ss := &Stack[string]{}
	ss.Push("hello")
	ss.Push("world")
	if val, ok := ss.Peek(); ok {
		fmt.Println("Peek:", val) // world
	}

	// Test Reverse
	fmt.Println("\n=== Reverse ===")
	fmt.Println(Reverse([]int{1, 2, 3, 4, 5}))       // [5 4 3 2 1]
	fmt.Println(Reverse([]string{"a", "b", "c"}))     // [c b a]
	fmt.Println(Reverse([]int{}))                       // []

	// Run test cases
	allPassed := true

	// === Stack LIFO order ===
	ts := &Stack[int]{}
	ts.Push(1)
	ts.Push(2)
	ts.Push(3)
	if val, _ := ts.Pop(); val != 3 {
		fmt.Println("FAIL: LIFO order")
		allPassed = false
	}
	if val, _ := ts.Pop(); val != 2 {
		fmt.Println("FAIL: LIFO order second pop")
		allPassed = false
	}
	if ts.Size() != 1 {
		fmt.Println("FAIL: Size after 2 pops")
		allPassed = false
	}

	// === Empty stack behavior ===
	empty := &Stack[int]{}
	if !empty.IsEmpty() {
		fmt.Println("FAIL: new stack should be empty")
		allPassed = false
	}
	if empty.Size() != 0 {
		fmt.Println("FAIL: new stack size should be 0")
		allPassed = false
	}
	if _, ok := empty.Pop(); ok {
		fmt.Println("FAIL: pop empty should return false")
		allPassed = false
	}
	if _, ok := empty.Peek(); ok {
		fmt.Println("FAIL: peek empty should return false")
		allPassed = false
	}

	// === Push then Pop all ===
	allStack := &Stack[int]{}
	allStack.Push(10)
	allStack.Push(20)
	allStack.Push(30)
	v1, ok1 := allStack.Pop()
	v2, ok2 := allStack.Pop()
	v3, ok3 := allStack.Pop()
	v4, ok4 := allStack.Pop()
	if v1 != 30 || !ok1 || v2 != 20 || !ok2 || v3 != 10 || !ok3 {
		fmt.Println("FAIL: Pop all three")
		allPassed = false
	}
	if ok4 {
		fmt.Println("FAIL: Pop beyond empty should return false")
		allPassed = false
	}
	_ = v4

	// === Peek doesn't modify ===
	peekStack := &Stack[int]{}
	peekStack.Push(99)
	pv1, _ := peekStack.Peek()
	pv2, _ := peekStack.Peek()
	if pv1 != 99 || pv2 != 99 || peekStack.Size() != 1 {
		fmt.Println("FAIL: Peek should not modify stack")
		allPassed = false
	}

	// === IsEmpty after push and pop ===
	toggleStack := &Stack[int]{}
	toggleStack.Push(1)
	if toggleStack.IsEmpty() {
		fmt.Println("FAIL: stack with element should not be empty")
		allPassed = false
	}
	toggleStack.Pop()
	if !toggleStack.IsEmpty() {
		fmt.Println("FAIL: stack after popping all should be empty")
		allPassed = false
	}

	// === String stack ===
	strStack := &Stack[string]{}
	strStack.Push("a")
	strStack.Push("b")
	if sv, _ := strStack.Pop(); sv != "b" {
		fmt.Println("FAIL: String stack LIFO")
		allPassed = false
	}
	if sv, _ := strStack.Pop(); sv != "a" {
		fmt.Println("FAIL: String stack second pop")
		allPassed = false
	}

	// === Reverse tests ===
	// preserves original
	original := []int{1, 2, 3}
	reversed := Reverse(original)
	if original[0] != 1 || original[2] != 3 {
		fmt.Println("FAIL: Reverse should not modify original")
		allPassed = false
	}
	if reversed[0] != 3 || reversed[2] != 1 {
		fmt.Println("FAIL: Reverse result incorrect")
		allPassed = false
	}
	// single element
	single := Reverse([]int{42})
	if len(single) != 1 || single[0] != 42 {
		fmt.Println("FAIL: Reverse single element")
		allPassed = false
	}
	// empty
	emptyRev := Reverse([]int{})
	if len(emptyRev) != 0 {
		fmt.Println("FAIL: Reverse empty")
		allPassed = false
	}
	// two elements
	two := Reverse([]int{1, 2})
	if len(two) != 2 || two[0] != 2 || two[1] != 1 {
		fmt.Println("FAIL: Reverse two elements")
		allPassed = false
	}
	// reverse strings
	revStr := Reverse([]string{"a", "b", "c"})
	if revStr[0] != "c" || revStr[1] != "b" || revStr[2] != "a" {
		fmt.Println("FAIL: Reverse strings")
		allPassed = false
	}
	// palindrome (reverse should equal original values)
	palindrome := Reverse([]int{1, 2, 1})
	if palindrome[0] != 1 || palindrome[1] != 2 || palindrome[2] != 1 {
		fmt.Println("FAIL: Reverse palindrome")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
