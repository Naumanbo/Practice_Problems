package main

import "fmt"

// Tests: Pointers, pointer receivers, linked data structures
//
// Build a singly linked list using pointers.
//
// 1. Define a Node struct: Value int, Next *Node
//
// 2. Define a LinkedList struct: Head *Node, size int
//
// 3. Implement these methods (pointer receivers):
//    - PushFront(val int)        - insert at head, O(1)
//    - PushBack(val int)         - insert at tail, O(n)
//    - PopFront() (int, bool)    - remove and return head value; false if empty
//    - Find(val int) *Node       - return pointer to first node with value, nil if not found
//    - Delete(val int) bool      - remove first node with value, return true if found
//    - Reverse()                 - reverse the list in place (no new allocations)
//    - Size() int                - return number of elements
//    - ToSlice() []int           - return values as a slice (for easy testing)
//
// Key concepts tested:
// - Pointer manipulation (next pointers, nil checks)
// - Why pointer receivers are required here (modifying Head)
// - The "previous pointer" pattern for deletion

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	size int
}

// TODO: Implement PushFront
func (ll *LinkedList) PushFront(val int) {
}

// TODO: Implement PushBack
func (ll *LinkedList) PushBack(val int) {
}

// TODO: Implement PopFront
func (ll *LinkedList) PopFront() (int, bool) {
	return 0, false
}

// TODO: Implement Find
func (ll *LinkedList) Find(val int) *Node {
	return nil
}

// TODO: Implement Delete
func (ll *LinkedList) Delete(val int) bool {
	return false
}

// TODO: Implement Reverse
func (ll *LinkedList) Reverse() {
}

// TODO: Implement Size
func (ll *LinkedList) Size() int {
	return 0
}

// TODO: Implement ToSlice
func (ll *LinkedList) ToSlice() []int {
	return nil
}

func main() {
	ll := &LinkedList{}

	fmt.Println("=== PushFront / PushBack ===")
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	fmt.Println(ll.ToSlice()) // [1 2 3]

	ll.PushFront(0)
	fmt.Println(ll.ToSlice()) // [0 1 2 3]
	fmt.Println("Size:", ll.Size()) // 4

	fmt.Println("\n=== Find ===")
	node := ll.Find(2)
	if node != nil {
		fmt.Println("Found:", node.Value) // 2
	}
	fmt.Println("Find 99:", ll.Find(99)) // <nil>

	fmt.Println("\n=== Delete ===")
	ll.Delete(0) // delete head
	fmt.Println("After delete 0:", ll.ToSlice()) // [1 2 3]

	ll.Delete(2) // delete middle
	fmt.Println("After delete 2:", ll.ToSlice()) // [1 3]

	ll.Delete(3) // delete tail
	fmt.Println("After delete 3:", ll.ToSlice()) // [1]

	fmt.Println("\n=== PopFront ===")
	val, ok := ll.PopFront()
	fmt.Printf("PopFront: %d, ok=%v\n", val, ok) // 1, true
	_, ok = ll.PopFront()
	fmt.Printf("PopFront empty: ok=%v\n", ok) // false

	fmt.Println("\n=== Reverse ===")
	ll2 := &LinkedList{}
	ll2.PushBack(1)
	ll2.PushBack(2)
	ll2.PushBack(3)
	ll2.PushBack(4)
	ll2.PushBack(5)
	ll2.Reverse()
	fmt.Println("Reversed:", ll2.ToSlice()) // [5 4 3 2 1]

	// Run test cases
	allPassed := true

	// PushFront on empty
	t1 := &LinkedList{}
	t1.PushFront(42)
	if t1.Size() != 1 || t1.ToSlice()[0] != 42 {
		fmt.Println("FAIL: PushFront on empty")
		allPassed = false
	}

	// PushBack on empty
	t2 := &LinkedList{}
	t2.PushBack(42)
	if t2.Size() != 1 || t2.ToSlice()[0] != 42 {
		fmt.Println("FAIL: PushBack on empty")
		allPassed = false
	}

	// Delete only element
	t3 := &LinkedList{}
	t3.PushBack(1)
	t3.Delete(1)
	if t3.Size() != 0 || len(t3.ToSlice()) != 0 {
		fmt.Println("FAIL: Delete only element")
		allPassed = false
	}

	// Delete nonexistent
	t4 := &LinkedList{}
	t4.PushBack(1)
	if t4.Delete(99) {
		fmt.Println("FAIL: Delete nonexistent should return false")
		allPassed = false
	}

	// Delete first of duplicates
	t5 := &LinkedList{}
	t5.PushBack(1)
	t5.PushBack(2)
	t5.PushBack(1)
	t5.PushBack(3)
	t5.Delete(1)
	if fmt.Sprint(t5.ToSlice()) != "[2 1 3]" {
		fmt.Println("FAIL: Delete should remove only first occurrence, got", t5.ToSlice())
		allPassed = false
	}

	// Reverse empty list
	t6 := &LinkedList{}
	t6.Reverse()
	if t6.Size() != 0 {
		fmt.Println("FAIL: Reverse empty")
		allPassed = false
	}

	// Reverse single element
	t7 := &LinkedList{}
	t7.PushBack(42)
	t7.Reverse()
	if fmt.Sprint(t7.ToSlice()) != "[42]" {
		fmt.Println("FAIL: Reverse single element")
		allPassed = false
	}

	// Reverse two elements
	t8 := &LinkedList{}
	t8.PushBack(1)
	t8.PushBack(2)
	t8.Reverse()
	if fmt.Sprint(t8.ToSlice()) != "[2 1]" {
		fmt.Println("FAIL: Reverse two elements")
		allPassed = false
	}

	// Size tracks correctly through operations
	t9 := &LinkedList{}
	t9.PushBack(1)
	t9.PushBack(2)
	t9.PushFront(0)
	t9.Delete(1)
	t9.PopFront()
	if t9.Size() != 1 {
		fmt.Printf("FAIL: Size tracking, expected 1 got %d\n", t9.Size())
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
