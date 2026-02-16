package main

import "fmt"

// Tests: Pointers, pointer receivers, linked data structures
//
// KEY TAKEAWAYS:
// - append() returns a NEW slice — you must assign it back: s = append(s, val).
//   Forgetting this is a silent bug; the original slice stays unchanged.
// - Linked list Delete: "last node" and "middle node" cases merge into one:
//   predecessor.Next = node.Next works for both. Don't overengineer with 3 cases.
// - Always decrement size counters (ll.size--) in Delete — easy to forget.
// - Reverse uses the three-pointer technique (prev, current, next). Order matters:
//   1. next = current.Next (save reference before overwriting)
//   2. current.Next = prev (reverse the pointer)
//   3. prev = current (advance prev)
//   4. current = next (advance current)
//   Swapping steps 3 and 4 loses your reference. Set ll.Head = prev at the end.
// - Always handle the empty list case (ll.Head == nil) to avoid nil pointer panics.
// - In-place pointer reversal is O(n) time, O(1) space — don't convert to slice and rebuild.
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
	// 0 -> 1 -> 2, pushFront(3): 3 -> 0 -> 1 -> 2
	var pushNode Node = Node{val, ll.Head}
	ll.Head = &pushNode
	ll.size += 1
}

// TODO: Implement PushBack
func (ll *LinkedList) PushBack(val int) {
	var pushNode Node = Node{val, nil}
	if ll.Head == nil {
		ll.Head = &pushNode
	} else {
		// 1. Set the Next pointer for second to last element to pushNode
		// 2. Push pushNode and give it a nil next pointer
		var tmp *Node = ll.Head
		for ; tmp.Next != nil; tmp = tmp.Next {
			continue
		}
		tmp.Next = &pushNode

	}

	ll.size += 1

}

// TODO: Implement PopFront
func (ll *LinkedList) PopFront() (int, bool) {
	if ll.Head != nil {
		var newHead *Node = ll.Head.Next
		var ret int = ll.Head.Value

		ll.Head = newHead
		ll.size--
		return ret, true

	} else {
		return 0, false
	}
}

// TODO: Implement Find
func (ll *LinkedList) Find(val int) *Node {

	for tmp := ll.Head; tmp != nil; tmp = tmp.Next {
		if tmp.Value == val {
			return tmp
		}
	}

	return nil

}

// TODO: Implement Delete
func (ll *LinkedList) Delete(val int) bool {
	/* Test case 1: 0 -> 1 -> 2, delete (1)
		(1) store 1.Next
		(2) Find element before 1
		(3) Replace 0.Next with 1.Next (2)
	   Test Case 2: 0 -> 1, delete (1)
	   	(1) get element before last element in list
		(2) set last element next node to nil
	   Test Case 3: 0 -> 1 -> 2, delete (0)
	   	(1) if find(0) == ll.Head,
		(2) store tmp = ll.Head.Next
		(3) set ll.Head.Next = nil
		(4) set ll.Head = tmp
	   Test Case 4: 0, delete (0)
	   	(1) This is both first and last element. What will happen?
		(2) will default to first element case since ll.Head == node and it will perform correctly because tmp will be nil if only one element in list so new head will become nil as expected
	*/
	node := ll.Find(val)
	if node != nil {
		// 3 different cases:
		// 1. It is first element in the list
		if ll.Head == node {
			tmp := ll.Head.Next
			ll.Head = tmp
		} else if node.Next == nil { // 2. It is the last element of the list
			for tmp := ll.Head; tmp.Next != nil; tmp = tmp.Next {
				if tmp.Next == node {
					tmp.Next = nil
					break
				}
			}
		} else { // 3. It is in the middle somewhere of the list
			elt := node.Next
			for tmp := ll.Head; tmp.Next != nil; tmp = tmp.Next {
				if tmp.Next == node {
					tmp.Next = elt
					break
				}
			}
		}

		ll.size--
		return true
	}
	return false
}

// TODO: Implement Reverse
func (ll *LinkedList) Reverse() {
	// Reverse in place, use prev, curr and next pointers
	var prev *Node = nil
	current := ll.Head
	var next *Node = nil

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	ll.Head = prev

}

// TODO: Implement Size
func (ll *LinkedList) Size() int {
	return ll.size
}

// TODO: Implement ToSlice
func (ll *LinkedList) ToSlice() []int {
	if ll.Head != nil {
		var s []int
		for tmp := ll.Head; tmp != nil; tmp = tmp.Next {
			s = append(s, tmp.Value)
		}
		return s
	}
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
	fmt.Println(ll.ToSlice())       // [0 1 2 3]
	fmt.Println("Size:", ll.Size()) // 4

	fmt.Println("\n=== Find ===")
	node := ll.Find(2)
	if node != nil {
		fmt.Println("Found:", node.Value) // 2
	}
	fmt.Println("Find 99:", ll.Find(99)) // <nil>

	fmt.Println("\n=== Delete ===")
	ll.Delete(0)                                 // delete head
	fmt.Println("After delete 0:", ll.ToSlice()) // [1 2 3]

	ll.Delete(2)                                 // delete middle
	fmt.Println("After delete 2:", ll.ToSlice()) // [1 3]

	ll.Delete(3)                                 // delete tail
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

	// === Additional edge case tests ===

	// PopFront on empty list
	emptyLL := &LinkedList{}
	if _, ok := emptyLL.PopFront(); ok {
		fmt.Println("FAIL: PopFront on empty should return false")
		allPassed = false
	}

	// Find on empty list
	if emptyLL.Find(1) != nil {
		fmt.Println("FAIL: Find on empty should return nil")
		allPassed = false
	}

	// Delete on empty list
	if emptyLL.Delete(1) {
		fmt.Println("FAIL: Delete on empty should return false")
		allPassed = false
	}

	// ToSlice on empty list
	if len(emptyLL.ToSlice()) != 0 {
		fmt.Println("FAIL: ToSlice on empty should return empty slice")
		allPassed = false
	}

	// PushFront multiple then verify order
	t10 := &LinkedList{}
	t10.PushFront(3)
	t10.PushFront(2)
	t10.PushFront(1)
	if fmt.Sprint(t10.ToSlice()) != "[1 2 3]" {
		fmt.Println("FAIL: PushFront order, got", t10.ToSlice())
		allPassed = false
	}

	// PushBack then PushFront interleaved
	t11 := &LinkedList{}
	t11.PushBack(2)
	t11.PushFront(1)
	t11.PushBack(3)
	if fmt.Sprint(t11.ToSlice()) != "[1 2 3]" {
		fmt.Println("FAIL: Mixed Push order, got", t11.ToSlice())
		allPassed = false
	}

	// Delete head when multiple elements
	t12 := &LinkedList{}
	t12.PushBack(10)
	t12.PushBack(20)
	t12.PushBack(30)
	t12.Delete(10)
	if fmt.Sprint(t12.ToSlice()) != "[20 30]" || t12.Size() != 2 {
		fmt.Println("FAIL: Delete head of 3-element list")
		allPassed = false
	}

	// Delete tail when multiple elements
	t13 := &LinkedList{}
	t13.PushBack(10)
	t13.PushBack(20)
	t13.PushBack(30)
	t13.Delete(30)
	if fmt.Sprint(t13.ToSlice()) != "[10 20]" || t13.Size() != 2 {
		fmt.Println("FAIL: Delete tail of 3-element list")
		allPassed = false
	}

	// Reverse preserves size
	t14 := &LinkedList{}
	t14.PushBack(1)
	t14.PushBack(2)
	t14.PushBack(3)
	t14.Reverse()
	if t14.Size() != 3 {
		fmt.Println("FAIL: Reverse should preserve size")
		allPassed = false
	}
	if fmt.Sprint(t14.ToSlice()) != "[3 2 1]" {
		fmt.Println("FAIL: Reverse 3 elements, got", t14.ToSlice())
		allPassed = false
	}

	// Double reverse returns original order
	t15 := &LinkedList{}
	t15.PushBack(1)
	t15.PushBack(2)
	t15.PushBack(3)
	t15.Reverse()
	t15.Reverse()
	if fmt.Sprint(t15.ToSlice()) != "[1 2 3]" {
		fmt.Println("FAIL: Double reverse should restore original order")
		allPassed = false
	}

	// Find returns first occurrence
	t16 := &LinkedList{}
	t16.PushBack(5)
	t16.PushBack(10)
	t16.PushBack(5)
	found := t16.Find(5)
	if found == nil || found.Value != 5 {
		fmt.Println("FAIL: Find should return node")
		allPassed = false
	}

	// PopFront returns correct value and decrements size
	t17 := &LinkedList{}
	t17.PushBack(100)
	t17.PushBack(200)
	val2, ok2 := t17.PopFront()
	if val2 != 100 || !ok2 || t17.Size() != 1 {
		fmt.Println("FAIL: PopFront value and size")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
