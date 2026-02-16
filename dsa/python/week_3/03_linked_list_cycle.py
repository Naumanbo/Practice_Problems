"""
DSA Problem: Linked List Cycle

Tests: Floyd's cycle detection, two pointers (fast/slow), space optimization

Difficulty: Easy
Source: LeetCode #141

Problem:
Given head of a linked list, determine if it has a cycle.

Constraints:
    - Number of nodes in range [0, 10^4]
    - -10^5 <= Node.val <= 10^5
"""

from typing import Optional


class ListNode:
    def __init__(self, val: int = 0, next: "ListNode" = None):
        self.val = val
        self.next = next


def has_cycle_set(head: Optional[ListNode]) -> bool:
    """
    Hash set approach - track visited nodes.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


def has_cycle_floyd(head: Optional[ListNode]) -> bool:
    """
    Floyd's tortoise and hare - O(1) space.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


# =============================================================================
# Helper functions
# =============================================================================
def make_list_with_cycle(values: list, cycle_pos: int) -> Optional[ListNode]:
    """Create a linked list. cycle_pos = index where tail connects back (-1 = no cycle)."""
    if not values:
        return None
    nodes = [ListNode(v) for v in values]
    for i in range(len(nodes) - 1):
        nodes[i].next = nodes[i + 1]
    if cycle_pos >= 0:
        nodes[-1].next = nodes[cycle_pos]
    return nodes[0]


# =============================================================================
# Test Cases - LeetCode Level
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"values": [3, 2, 0, -4], "cycle_pos": 1, "expected": True, "desc": "cycle at pos 1"},
        {"values": [1, 2], "cycle_pos": 0, "expected": True, "desc": "cycle at head"},
        {"values": [1, 2, 3, 4], "cycle_pos": 0, "expected": True, "desc": "cycle back to head"},
        {"values": [1, 2, 3, 4], "cycle_pos": 2, "expected": True, "desc": "cycle in middle"},
        {"values": [1], "cycle_pos": -1, "expected": False, "desc": "single node no cycle"},
        {"values": [1, 2], "cycle_pos": -1, "expected": False, "desc": "two nodes no cycle"},
        {"values": [1, 2, 3, 4, 5], "cycle_pos": -1, "expected": False, "desc": "five nodes no cycle"},
        {"values": [], "cycle_pos": -1, "expected": False, "desc": "empty list"},
        {"values": [1], "cycle_pos": 0, "expected": True, "desc": "self-loop"},
        {"values": [1, 2, 3], "cycle_pos": -1, "expected": False, "desc": "three nodes no cycle"},
        {"values": list(range(100)), "cycle_pos": -1, "expected": False, "desc": "large list no cycle"},
        {"values": list(range(100)), "cycle_pos": 50, "expected": True, "desc": "large list cycle at 50"},
        {"values": list(range(100)), "cycle_pos": 0, "expected": True, "desc": "large list cycle at head"},
        {"values": list(range(100)), "cycle_pos": 99, "expected": True, "desc": "large list self-loop tail"},
        {"values": [1, 2, 3, 4, 5], "cycle_pos": 4, "expected": True, "desc": "tail self-loop"},
        {"values": [1, 2, 3, 4, 5], "cycle_pos": 3, "expected": True, "desc": "cycle at second to last"},
        {"values": [1, 2, 3, 4, 5], "cycle_pos": 1, "expected": True, "desc": "cycle at pos 1"},
        {"values": [-1, -2, -3], "cycle_pos": -1, "expected": False, "desc": "negative values no cycle"},
        {"values": [-1, -2, -3], "cycle_pos": 0, "expected": True, "desc": "negative values with cycle"},
        {"values": [1, 2], "cycle_pos": 1, "expected": True, "desc": "two nodes tail self-loop"},
    ]

    print("=" * 70)
    print("LINKED LIST CYCLE - Test Results")
    print("=" * 70)

    for approach_name, func in [("HashSet", has_cycle_set), ("Floyd's", has_cycle_floyd)]:
        print(f"\n{approach_name}:")
        passed = 0
        total = len(test_cases)
        for i, tc in enumerate(test_cases, 1):
            head = make_list_with_cycle(tc["values"], tc["cycle_pos"])
            result = func(head)
            ok = result == tc["expected"]
            passed += ok
            status = "PASS" if ok else "FAIL"
            vals = tc["values"] if len(tc["values"]) <= 8 else tc["values"][:5] + ["..."]
            print(f"  {i:2}. [{status}] {tc['desc']}: vals={vals}, pos={tc['cycle_pos']} -> {result}")
        print(f"  Result: {passed}/{total}")

    print("\n" + "=" * 70)
    print("Questions:")
    print("1. Why does Floyd's algorithm guarantee the fast pointer meets the slow pointer?")
    print("2. What's the trade-off between the hash set and Floyd's approach?")
    print("3. Follow-up: How would you find WHERE the cycle begins? (LeetCode #142)")
