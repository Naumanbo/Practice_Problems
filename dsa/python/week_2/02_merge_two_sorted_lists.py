"""
Merge Two Sorted Lists (LeetCode 21)

Tests: Linked lists, two pointers, recursion vs iteration

Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the input lists.

Example 1:
    Input: list1 = [1,2,4], list2 = [1,3,4]
    Output: [1,1,2,3,4,4]

Example 2:
    Input: list1 = [], list2 = []
    Output: []

Example 3:
    Input: list1 = [], list2 = [0]
    Output: [0]

Constraints:
    - The number of nodes in both lists is in range [0, 50]
    - -100 <= Node.val <= 100
    - Both lists are sorted in non-decreasing order
"""

from typing import Optional, List


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode" = None):
        self.val = val
        self.next = next


def merge_two_lists(list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
    # TODO: Implement solution
    # Hint: Use a dummy head node to simplify edge cases
    pass


# Helper functions for testing
def list_to_linked(arr: List[int]) -> Optional[ListNode]:
    if not arr:
        return None
    head = ListNode(arr[0])
    current = head
    for val in arr[1:]:
        current.next = ListNode(val)
        current = current.next
    return head


def linked_to_list(head: Optional[ListNode]) -> List[int]:
    result = []
    while head:
        result.append(head.val)
        head = head.next
    return result


# Tests
if __name__ == "__main__":
    test_cases = [
        {"list1": [1, 2, 4], "list2": [1, 3, 4], "expected": [1, 1, 2, 3, 4, 4], "desc": "basic merge"},
        {"list1": [], "list2": [], "expected": [], "desc": "both empty"},
        {"list1": [], "list2": [0], "expected": [0], "desc": "first empty"},
        {"list1": [1], "list2": [], "expected": [1], "desc": "second empty"},
        {"list1": [1, 3, 5], "list2": [2, 4, 6], "expected": [1, 2, 3, 4, 5, 6], "desc": "interleaved"},
        {"list1": [1, 2, 3], "list2": [4, 5, 6], "expected": [1, 2, 3, 4, 5, 6], "desc": "no overlap"},
        {"list1": [5], "list2": [1, 2, 3], "expected": [1, 2, 3, 5], "desc": "single vs multiple"},
        {"list1": [1, 1, 1], "list2": [1, 1], "expected": [1, 1, 1, 1, 1], "desc": "all duplicates"},
        {"list1": [-3, -1, 0], "list2": [-2, 5], "expected": [-3, -2, -1, 0, 5], "desc": "negative values"},
        {"list1": [1], "list2": [2], "expected": [1, 2], "desc": "single element each"},
        {"list1": [2], "list2": [1], "expected": [1, 2], "desc": "single elements reversed"},
        {"list1": [1], "list2": [2, 3, 4, 5, 6, 7, 8, 9, 10], "expected": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], "desc": "one short one long"},
        {"list1": [5, 5, 5], "list2": [5, 5, 5], "expected": [5, 5, 5, 5, 5, 5], "desc": "all same values"},
        {"list1": [1, 2, 3], "list2": [1, 2, 3], "expected": [1, 1, 2, 2, 3, 3], "desc": "identical lists"},
        {"list1": [1000000], "list2": [-1000000], "expected": [-1000000, 1000000], "desc": "large values"},
        {"list1": [-100, -50, 0, 50, 100], "list2": [-75, -25, 25, 75], "expected": [-100, -75, -50, -25, 0, 25, 50, 75, 100], "desc": "interleaved with negatives"},
        {"list1": [1, 2, 3, 4, 5], "list2": [6, 7, 8, 9, 10], "expected": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], "desc": "no overlap longer"},
    ]

    all_passed = True
    for tc in test_cases:
        l1 = list_to_linked(tc["list1"])
        l2 = list_to_linked(tc["list2"])
        result = linked_to_list(merge_two_lists(l1, l2))
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: merge({tc['list1']}, {tc['list2']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
