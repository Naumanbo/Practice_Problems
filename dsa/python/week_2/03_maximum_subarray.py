"""
Maximum Subarray (LeetCode 53)

Tests: Dynamic programming, Kadane's algorithm, contiguous subarray

Given an integer array nums, find the contiguous subarray with the largest sum,
and return its sum.

Example 1:
    Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
    Output: 6
    Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:
    Input: nums = [1]
    Output: 1

Example 3:
    Input: nums = [5,4,-1,7,8]
    Output: 23
    Explanation: The entire array is the subarray with maximum sum.

Constraints:
    - 1 <= nums.length <= 10^5
    - -10^4 <= nums[i] <= 10^4

Follow-up: Can you solve it in O(n) time with O(1) space? (Kadane's algorithm)
"""

from typing import List


def max_subarray(nums: List[int]) -> int:
    # TODO: Implement Kadane's algorithm
    # Hint: At each position, decide: extend current subarray or start new one?
    pass


# Tests
if __name__ == "__main__":
    test_cases = [
        {"nums": [-2, 1, -3, 4, -1, 2, 1, -5, 4], "expected": 6, "desc": "mixed values"},
        {"nums": [1], "expected": 1, "desc": "single element"},
        {"nums": [5, 4, -1, 7, 8], "expected": 23, "desc": "mostly positive"},
        {"nums": [-1], "expected": -1, "desc": "single negative"},
        {"nums": [-2, -1], "expected": -1, "desc": "all negative"},
        {"nums": [-2, -3, -1, -5], "expected": -1, "desc": "all negative longer"},
        {"nums": [1, 2, 3, 4], "expected": 10, "desc": "all positive"},
        {"nums": [-1, 0, -2], "expected": 0, "desc": "zero is max"},
        {"nums": [8, -19, 5, -4, 20], "expected": 21, "desc": "recovery after negative"},
    ]

    all_passed = True
    for tc in test_cases:
        result = max_subarray(tc["nums"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: max_subarray({tc['nums']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
