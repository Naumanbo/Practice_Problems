# Key Takeaways:
# 1. Array-as-linked-list translation: treat index i as a node, and nums[i] as
#    the "next pointer". You jump to index nums[i], not index i+1. This is the
#    critical mental shift — neither slow nor fast ever increments by +1.
#      slow = nums[slow]           # follow value once
#      fast = nums[nums[fast]]     # follow value twice
#
# 2. The duplicate creates a cycle because two indices hold the same value,
#    meaning two different nodes both point to the same next index. That shared
#    destination is the cycle entrance — and it equals the duplicate value.
#
# 3. Floyd's has two phases:
#    Phase 1 — find intersection: move slow 1 step, fast 2 steps until they meet.
#    Phase 2 — find cycle entrance: reset slow to nums[0], move both 1 step until
#    they meet again. Meeting point = cycle entrance = duplicate.
#
# 4. Use `while True` with a break for phase 1 — NOT `while slow != fast`.
#    Both pointers start at nums[0] so they're already equal before any movement.
#    A pre-condition check exits immediately. `while True` is Python's do-while:
#    move first, check after.
#
# 5. Why phase 2 works: the distance from start to cycle entrance (F) equals
#    the remaining distance from the intersection back to the entrance.
#    So one pointer at nums[0] and one at the intersection both travel F steps
#    and arrive at the entrance simultaneously.
#
# 6. Pigeonhole Principle: if you have n+1 values all in range [1,n], at least
#    one value must repeat — you have more pigeons (values) than holes (slots).
#    This is the mathematical guarantee that a duplicate always exists and the
#    algorithm will always terminate. Without this guarantee, Floyd's could loop
#    indefinitely or go out of bounds.
#
# 7. The constraint `1 <= nums[i] <= n` is not just a boundary — it's the
#    structural requirement that makes the array representable as a linked list.
#    Every value is a valid index. Without it, nums[slow] could go out of bounds.
#    Always check constraints for hidden structural guarantees like this.
#
# 8. Same Floyd's algorithm as linked list cycle detection (week_3/03) — the only
#    difference is how "next" is defined. Linked list: node.next. Array: nums[i].
#    Recognizing this isomorphism (two problems with identical structure) is a
#    core skill for reducing new problems to ones you already know.
#
# Related EECS 281 Lectures: Lec 19 (Graph Introduction — cycle detection, directed
#   graphs), Lec 15 (Hash Tables — hash set as O(n) space alternative)
# Complexity: Time O(n), Space O(1)

"""
DSA Problem: Find the Duplicate Number

Tests: Floyd's cycle detection on arrays, two pointers, O(1) space constraint

Difficulty: Medium
Source: LeetCode #287

Problem:
Given an array nums of n+1 integers where each is in range [1, n],
there is exactly one repeated number. Find it without modifying the array
and using only O(1) extra space.

Constraints:
    - 1 <= n <= 10^5
    - nums.length == n + 1
    - 1 <= nums[i] <= n
    - Exactly one value is repeated
"""

from typing import List


def find_duplicate(nums: List[int]) -> int:
    slow = nums[0]
    fast = nums[0]

    while True:
        slow = nums[slow]
        fast = nums[nums[fast]]

        if slow == fast:
            break
    
    slow = nums[0]
    
    while fast != slow:
        slow = nums[slow]
        fast = nums[fast]
    return slow



# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"nums": [1, 3, 4, 2, 2],                   "expected": 2, "desc": "basic case"},
        {"nums": [3, 1, 3, 4, 2],                   "expected": 3, "desc": "duplicate not adjacent"},
        {"nums": [1, 1],                             "expected": 1, "desc": "two elements"},
        {"nums": [1, 1, 2],                          "expected": 1, "desc": "three elements"},
        {"nums": [1, 2, 3, 4, 4],                   "expected": 4, "desc": "duplicate at end"},
        {"nums": [1, 2, 3, 1],                      "expected": 1, "desc": "duplicate wraps"},
        {"nums": [2, 1, 2, 3],                      "expected": 2, "desc": "duplicate early"},
        {"nums": [6, 2, 4, 1, 3, 5, 6],             "expected": 6, "desc": "six elements"},
        {"nums": [9, 7, 4, 6, 3, 2, 8, 5, 1, 1],   "expected": 1, "desc": "duplicate is 1"},
        {"nums": [2, 5, 9, 6, 3, 8, 7, 1, 4, 9],   "expected": 9, "desc": "duplicate is 9"},
        {"nums": [3, 4, 8, 5, 9, 1, 6, 8, 7, 2],   "expected": 8, "desc": "duplicate in middle"},
        {"nums": [5, 1, 2, 3, 4, 5],                "expected": 5, "desc": "duplicate at boundaries"},
        {"nums": [1, 2, 3, 2, 4],                   "expected": 2, "desc": "five elements"},
        {"nums": [4, 3, 2, 1, 4],                   "expected": 4, "desc": "duplicate at front and back"},
        {"nums": [1, 2, 1, 3],                      "expected": 1, "desc": "four elements"},
        {"nums": [10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 9], "expected": 9, "desc": "reverse sorted with dup"},
    ]

    all_passed = True
    for tc in test_cases:
        result = find_duplicate(tc["nums"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: find_duplicate({tc['nums']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
