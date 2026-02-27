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
    pass


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
