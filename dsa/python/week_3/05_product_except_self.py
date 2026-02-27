"""
DSA Problem: Product of Array Except Self

Tests: Arrays, prefix/suffix products, no-division constraint

Difficulty: Medium
Source: LeetCode #238

Problem:
Given an integer array nums, return an array answer such that answer[i] is
equal to the product of all elements of nums except nums[i].
You must solve it in O(n) time without using division.

Constraints:
    - 2 <= nums.length <= 10^5
    - -30 <= nums[i] <= 30
    - The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer
"""

from typing import List


def product_except_self(nums: List[int]) -> List[int]:
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"nums": [1, 2, 3, 4],        "expected": [24, 12, 8, 6],           "desc": "basic case"},
        {"nums": [2, 3, 4, 5],        "expected": [60, 40, 30, 24],         "desc": "all positive"},
        {"nums": [-1, 1, 0, -3, 3],   "expected": [0, 0, 9, 0, 0],         "desc": "contains zero"},
        {"nums": [1, 2],              "expected": [2, 1],                   "desc": "two elements"},
        {"nums": [3, 3, 3],           "expected": [9, 9, 9],                "desc": "all same"},
        {"nums": [1, 1, 1, 1],        "expected": [1, 1, 1, 1],             "desc": "all ones"},
        {"nums": [0, 0],              "expected": [0, 0],                   "desc": "all zeros"},
        {"nums": [1, 0],              "expected": [0, 1],                   "desc": "one zero"},
        {"nums": [-1, 2, 3],          "expected": [6, -3, -2],              "desc": "negative value"},
        {"nums": [2, 2, 2, 2],        "expected": [8, 8, 8, 8],             "desc": "all twos"},
        {"nums": [1, 2, 3, 4, 5],     "expected": [120, 60, 40, 30, 24],    "desc": "five elements"},
        {"nums": [-1, -2, -3, -4],    "expected": [-24, -12, -8, -6],       "desc": "all negative"},
        {"nums": [100, 1, 2],         "expected": [2, 200, 100],            "desc": "large first element"},
        {"nums": [0, 1, 2, 3],        "expected": [6, 0, 0, 0],             "desc": "zero at start"},
        {"nums": [1, 2, 0, 4],        "expected": [0, 0, 8, 0],             "desc": "zero in middle"},
        {"nums": [2, 3],              "expected": [3, 2],                   "desc": "two elements v2"},
    ]

    all_passed = True
    for tc in test_cases:
        result = product_except_self(tc["nums"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: product_except_self({tc['nums']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
