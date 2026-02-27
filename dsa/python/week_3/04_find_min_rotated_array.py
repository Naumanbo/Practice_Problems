"""
DSA Problem: Find Minimum in Rotated Sorted Array

Tests: Binary search variant, rotated arrays, boundary conditions

Difficulty: Medium
Source: LeetCode #153

Problem:
Given a sorted array that has been rotated between 1 and n times, find the minimum element.
You must write an algorithm that runs in O(log n) time.

Constraints:
    - 1 <= nums.length <= 5000
    - -5000 <= nums[i] <= 5000
    - All integers in nums are unique
    - nums is sorted and rotated between 1 and n times
"""

from typing import List


def find_min(nums: List[int]) -> int:
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"nums": [3, 4, 5, 1, 2],                       "expected": 1,   "desc": "basic rotated"},
        {"nums": [4, 5, 6, 7, 0, 1, 2],                 "expected": 0,   "desc": "min is zero"},
        {"nums": [11, 13, 15, 17],                       "expected": 11,  "desc": "not rotated"},
        {"nums": [1],                                     "expected": 1,   "desc": "single element"},
        {"nums": [2, 1],                                  "expected": 1,   "desc": "two elements rotated"},
        {"nums": [1, 2],                                  "expected": 1,   "desc": "two elements not rotated"},
        {"nums": [3, 1, 2],                               "expected": 1,   "desc": "three elements"},
        {"nums": [5, 6, 7, 8, 1, 2, 3, 4],              "expected": 1,   "desc": "rotated midpoint"},
        {"nums": [10, 1, 2, 3, 4, 5, 6, 7, 8, 9],      "expected": 1,   "desc": "rotated once"},
        {"nums": [2, 3, 4, 5, 6, 7, 8, 1],              "expected": 1,   "desc": "min at end"},
        {"nums": [6, 7, 1, 2, 3, 4, 5],                 "expected": 1,   "desc": "rotated two thirds"},
        {"nums": [-5, -3, -1, -10, -8],                  "expected": -10, "desc": "all negative"},
        {"nums": [0, 1, 2, 3, -1],                       "expected": -1,  "desc": "negative min at end"},
        {"nums": [100, 200, 300, 10, 50],                "expected": 10,  "desc": "large values rotated"},
        {"nums": [1, 2, 3, 4, 5],                        "expected": 1,   "desc": "fully sorted"},
        {"nums": [5, 1, 2, 3, 4],                        "expected": 1,   "desc": "rotated once from front"},
    ]

    all_passed = True
    for tc in test_cases:
        result = find_min(tc["nums"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: find_min({tc['nums']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
