"""
DSA Problem 1: Two Sum

Tests: Hash map usage, time/space complexity analysis, array traversal

Difficulty: Easy
Source: LeetCode #1

Problem:
Given an array of integers `nums` and an integer `target`, return indices
of the two numbers such that they add up to target.

Constraints:
    - 2 <= nums.length <= 10^4
    - -10^9 <= nums[i] <= 10^9
    - Only one valid answer exists.
    - You may not use the same element twice.
"""

from typing import List


def two_sum_brute(nums: List[int], target: int) -> List[int]:
    """
    Brute force approach.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


def two_sum_optimal(nums: List[int], target: int) -> List[int]:
    """
    Optimal approach using hash map.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        # Basic cases
        {"nums": [2, 7, 11, 15], "target": 9, "desc": "Basic - first two elements"},
        {"nums": [3, 2, 4], "target": 6, "desc": "Basic - middle elements"},
        {"nums": [3, 3], "target": 6, "desc": "Duplicate values"},

        # Edge cases - array size
        {"nums": [1, 2], "target": 3, "desc": "Minimum size (2 elements)"},
        {"nums": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10], "target": 19, "desc": "Larger array"},

        # Edge cases - negative numbers
        {"nums": [-1, -2, -3, -4], "target": -6, "desc": "All negative"},
        {"nums": [-3, 4, 3, 90], "target": 0, "desc": "Negative + positive = 0"},
        {"nums": [5, -5, 10], "target": 0, "desc": "Sum to zero"},

        # Edge cases - zeros
        {"nums": [0, 4, 3, 0], "target": 0, "desc": "Two zeros"},
        {"nums": [0, 1, 2], "target": 2, "desc": "Zero + positive"},

        # Edge cases - large numbers
        {"nums": [1000000000, 2, 1000000000], "target": 2000000000, "desc": "Large numbers"},

        # Edge cases - position variations
        {"nums": [1, 5, 8, 3], "target": 4, "desc": "Answer at start and end"},
        {"nums": [4, 5, 1, 2], "target": 6, "desc": "Answer in middle"},
        {"nums": [5, 5, 5, 5], "target": 10, "desc": "All same values"},
    ]

    def validate(result, nums, target):
        if result is None:
            return False
        if len(result) != 2:
            return False
        if result[0] == result[1]:
            return False
        return nums[result[0]] + nums[result[1]] == target

    print("=" * 70)
    print("TWO SUM - Test Results")
    print("=" * 70)

    brute_passed = 0
    optimal_passed = 0
    total = len(test_cases)

    for i, tc in enumerate(test_cases, 1):
        nums, target, desc = tc["nums"], tc["target"], tc["desc"]

        result_brute = two_sum_brute(nums.copy(), target)
        result_optimal = two_sum_optimal(nums.copy(), target)

        brute_ok = validate(result_brute, nums, target)
        optimal_ok = validate(result_optimal, nums, target)

        brute_passed += brute_ok
        optimal_passed += optimal_ok

        b_status = "PASS" if brute_ok else "FAIL"
        o_status = "PASS" if optimal_ok else "FAIL"

        print(f"\n{i}. {desc}")
        print(f"   Input: nums={nums}, target={target}")
        print(f"   Brute:   [{b_status}] {result_brute}")
        print(f"   Optimal: [{o_status}] {result_optimal}")

    print("\n" + "=" * 70)
    print(f"Summary: Brute {brute_passed}/{total} | Optimal {optimal_passed}/{total}")
    print("=" * 70)
    print("\nQuestions:")
    print("1. What is the time complexity of brute force? Why?")
    print("2. What is the time complexity of the hash map approach? Why?")
    print("3. Can this be solved with two pointers? What's the prerequisite?")
