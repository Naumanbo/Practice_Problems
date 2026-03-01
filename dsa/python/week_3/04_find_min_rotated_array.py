# Key Takeaways:
# 1. Compare nums[mid] to nums[right] (not nums[left]) to determine which half
#    contains the minimum. nums[right] always reliably indicates which side the
#    break point is on — nums[left] can mislead in the unrotated case.
#
# 2. Use `elif` not two separate `if` statements — running both branches in the
#    same iteration causes mid to bounce unpredictably and the loop never converges.
#    Two `if`s is a silent logic bug that produces wrong answers.
#
# 3. The clean template: `while left < right` with `right = mid` (not mid-1) on
#    the else branch. When left == right, both pointers are at the minimum.
#    Return nums[left] — no extra min-tracking variable needed.
#
# 4. Builds directly on standard binary search (week_3/01) — same skeleton,
#    different comparison logic. Recognizing the variant pattern is the key skill.
#
# Complexity: Time O(log n), Space O(1)

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
import sys
from typing import List


def find_min(nums: List[int]) -> int:
    # O(log n); cannot use built in sort functions, must do one binary search no iterating to sort either
    left = 0
    right = len(nums) - 1
    mid = 0
    min = nums[0]
    # My approach, embeds a bug by not using an if/else statement in checking whether min resides on left or right of mid.
    # Bug is resolved by always checking if nums[mid] < min every iteration since the loop currently ensures that nums[mid] may not be the min element because of how the if statements are setup.
    while left <= right:
        mid = left + (right - left) // 2
        if nums[mid] < min:
            min = nums[mid]
        if nums[mid] > nums[right]: # minimum item lies on right
            left = mid + 1
        if nums[mid] <= nums[right]: # min lies on left
            right = mid - 1
    
    return min

def find_min_opt(nums: List[int]) -> int:
    left = 0
    right = len(nums) - 1
    mid = 0

    while left <= right:
        mid = left + (right - left) // 2

        if nums[mid] > nums[right]:
            left = mid + 1
        else:
            right = mid
    return nums[left]

# [3 1 2]
'''
left = 0
right = 2

mid = 0 + (2 - 0) // 2 = 1
1 > 2 false
1 <= 2 true
right = 1 - 1 = 0


'''


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
