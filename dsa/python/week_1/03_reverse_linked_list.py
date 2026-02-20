# Key Takeaways:
# 1. Recursive template: (1) base case, (2) recurse on smaller input,
#    (3) combine result. Never recurse on the same input — something must shrink.
#
# 2. For recursive linked list reversal: trust the recursion reversed the rest,
#    then use head.next.next = head to flip the back-pointer, and head.next = None
#    to cut the forward link and avoid a cycle.
#
# 3. Iterative reversal uses three pointers (prev, current, next) and is O(1) space.
#    Recursive reversal is O(n) space due to the call stack — one frame per node.
#
# 4. Common recursion bugs: no base case (stack overflow), recursing on the same
#    input (infinite loop), or ignoring the return value (result lost, returns None).
#
# Complexity: Time O(n), Space O(1) iterative / O(n) recursive

"""
DSA Problem 3: Reverse Linked List

Tests: Linked list traversal, pointer manipulation, iterative vs recursive 

Difficulty: Easy
Source: LeetCode #206

Problem:
Given the head of a singly linked list, reverse the list and return it.
"""

from typing import Optional


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode" = None):
        self.val = val
        self.next = next


def reverse_iterative(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Iterative approach using three pointers.

    Time: O(n)
    Space: O(1)

    Your implementation:
    use prev, current, and next pointers to reverse list
    """
    if not head:
        return None
    
    prev = None
    current = head
    next = None

    while current:
        next = current.next
        current.next = prev
        prev = current
        current = next

    return prev

    pass


def reverse_recursive(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Recursive approach.

    Time: O(n)
    Space: O(n) - consider call stac. 

    Your implementation:
    """
    # base case
    if head is None or head.next is None:
        return head
    
    new_head = reverse_recursive(head.next)
    head.next.next = head
    head.next = None
    return new_head


# =============================================================================
# Helper functions
# =============================================================================
def list_to_linked(arr: list) -> Optional[ListNode]:
    if not arr:
        return None
    head = ListNode(arr[0])
    current = head
    for val in arr[1:]:
        current.next = ListNode(val)
        current = current.next
    return head


def linked_to_list(head: Optional[ListNode]) -> list:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        # Base cases
        {"input": [], "expected": [], "desc": "Empty list"},
        {"input": [1], "expected": [1], "desc": "Single element"},
        {"input": [1, 2], "expected": [2, 1], "desc": "Two elements"},

        # Basic cases
        {"input": [1, 2, 3], "expected": [3, 2, 1], "desc": "Three elements"},
        {"input": [1, 2, 3, 4, 5], "expected": [5, 4, 3, 2, 1], "desc": "Five elements"},

        # Edge cases - duplicates
        {"input": [1, 1, 1], "expected": [1, 1, 1], "desc": "All same values"},
        {"input": [1, 2, 2, 1], "expected": [1, 2, 2, 1], "desc": "Palindrome"},
        {"input": [1, 1, 2, 2], "expected": [2, 2, 1, 1], "desc": "Pairs"},

        # Edge cases - negative numbers
        {"input": [-1, -2, -3], "expected": [-3, -2, -1], "desc": "Negative values"},
        {"input": [-1, 0, 1], "expected": [1, 0, -1], "desc": "Mixed signs"},

        # Edge cases - large values
        {"input": [1000000, 2000000], "expected": [2000000, 1000000], "desc": "Large values"},

        # Longer lists
        {"input": list(range(10)), "expected": list(range(9, -1, -1)), "desc": "0-9 sequence"},
        {"input": list(range(1, 101)), "expected": list(range(100, 0, -1)), "desc": "1-100 sequence"},
    ]

    def run_tests(func, name):
        print(f"\n{name}:")
        passed = 0
        for i, tc in enumerate(test_cases, 1):
            inp, expected, desc = tc["input"], tc["expected"], tc["desc"]
            head = list_to_linked(inp.copy())
            result = linked_to_list(func(head))
            ok = result == expected
            passed += ok
            status = "PASS" if ok else "FAIL"

            display_in = inp if len(inp) <= 8 else inp[:5] + ["..."]
            display_out = result if len(result) <= 8 else result[:5] + ["..."]
            print(f"  {i:2}. [{status}] {desc}: {display_in} -> {display_out}")

        return passed

    print("=" * 70)
    print("REVERSE LINKED LIST - Test Results")
    print("=" * 70)

    total = len(test_cases)
    iter_passed = run_tests(reverse_iterative, "Iterative")
    rec_passed = run_tests(reverse_recursive, "Recursive")

    print("\n" + "=" * 70)
    print(f"Summary: Iterative {iter_passed}/{total} | Recursive {rec_passed}/{total}")
    print("=" * 70)
    print("\nQuestions:")
    print("1. Draw pointer changes for [1,2,3] step by step (iterative).")
    print("A: ")
    print("2. What's the space complexity difference between approaches?")
    print("A: iterative is O(1) while recursive is O(n) since each new call of itself creates a function call in the call stack until the base case is reached.")
    print("3. Why prefer iterative in production code?")
    print("A: Initial:")
    print("   prev=None  curr=1  next=?")
    print("   1 → 2 → 3 → None")
    print("")
    print("   Step 1 (curr=1):")
    print("     next = 2")
    print("     1.next = None  (point back to prev)")
    print("     prev = 1, curr = 2")
    print("     None ← 1    2 → 3 → None")
    print("")
    print("   Step 2 (curr=2):")
    print("     next = 3")
    print("     2.next = 1  (point back to prev)")
    print("     prev = 2, curr = 3")
    print("     None ← 1 ← 2    3 → None")
    print("")
    print("   Step 3 (curr=3):")
    print("     next = None")
    print("     3.next = 2  (point back to prev)")
    print("     prev = 3, curr = None")
    print("     None ← 1 ← 2 ← 3")
    print("")
    print("   curr=None, exit loop. Return prev=3.")
    print("3. Follow-up: How would you find WHERE the cycle begins? (LeetCode #142)")
    print("A: Because the call stack balloons in size with a recursive approach using exponential memory")
