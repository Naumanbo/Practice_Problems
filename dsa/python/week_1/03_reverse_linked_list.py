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

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


def reverse_recursive(head: Optional[ListNode]) -> Optional[ListNode]:
    """
    Recursive approach.

    Time: O(?)
    Space: O(?) - consider call stack

    Your implementation:
    """
    pass


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
    print("2. What's the space complexity difference between approaches?")
    print("3. Why prefer iterative in production code?")
