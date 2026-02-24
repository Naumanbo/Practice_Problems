# Key Takeaways:
# 1. Kadane's invariant:
#    `currSum` = best subarray sum that MUST end at the current index.
#    `maxSum` = best subarray sum seen anywhere so far.
#
# 2. At each index, make one decision:
#    start a new subarray at `nums[i]` OR extend the previous one (`currSum + nums[i]`).
#    Formula: `currSum = max(nums[i], currSum + nums[i])`
#
# 3. Update the global answer every iteration AFTER updating `currSum`:
#    `maxSum = max(maxSum, currSum)`
#
# 4. Initialize with `nums[0]` (not 0) so all-negative arrays work correctly.
#    Example: [-3, -1, -2] should return -1, not 0.
#
# 5. Avoid index lookahead (`nums[i+1]`) in this problem.
#    Kadane only needs the current value and previous running sum.
#
# Complexity: Time O(n), Space O(1)
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
    # Hint: At each position, decide: extend current subarray or start new one? greedy
    currSum = nums[0]
    maxSum = nums[0]
    start_new = False
    for i in range(1, len(nums)):
        currSum = max(nums[i], currSum + nums[i])
        maxSum = max(currSum, maxSum)
        
    return maxSum


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
        {"nums": [1, -1, 1, -1, 1], "expected": 1, "desc": "alternating"},
        {"nums": [5, 6, -100, 7, 8], "expected": 15, "desc": "large negative breaks subarray"},
        {"nums": [10, -1, -1, -1, -1], "expected": 10, "desc": "subarray at start"},
        {"nums": [-1, -1, -1, -1, 10], "expected": 10, "desc": "subarray at end"},
        {"nums": [5, -10, 5], "expected": 5, "desc": "two equal subarrays"},
        {"nums": [100], "expected": 100, "desc": "single large positive"},
        {"nums": [10, -5, 10, -5, 10], "expected": 20, "desc": "worth keeping negatives"},
        {"nums": [1, -1, 1, -1, 2], "expected": 2, "desc": "max at end"},
        {"nums": [3, -1, 2, -1, 4], "expected": 7, "desc": "entire array is max"},
        {"nums": [0], "expected": 0, "desc": "single zero"},
        {"nums": [-10000], "expected": -10000, "desc": "single large negative"},
        {"nums": [1, 2, -1, 3, -2, 4], "expected": 7, "desc": "scattered negatives"},
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
