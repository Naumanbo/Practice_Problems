# Key Takeaways:
# 1. Core insight: result[i] = (product of all elements LEFT of i) *
#    (product of all elements RIGHT of i). Every "except self" problem
#    decomposes into a left context times a right context.
#
# 2. Two-pass prefix/suffix pattern:
#    Pass 1 (left → right): build prefix products into result array.
#      res[0] = 1 (no elements to the left)
#      res[i] = res[i-1] * nums[i-1]
#    Pass 2 (right → left): multiply in suffix products using a running variable.
#      suffix starts at 1, grows as you walk left.
#      res[i] *= suffix; suffix *= nums[i]
#    After both passes, res[i] = left_product * right_product = product except self.
#
# 3. The suffix running variable is the key optimization — instead of storing
#    a full suffix array (O(n) extra space), one variable accumulates the
#    right-side product as you scan backwards. Space drops from O(n) to O(1)
#    extra (the result array itself doesn't count as extra space).
#
# 4. Zeros require no special handling — the prefix/suffix approach handles
#    them naturally. If there's one zero, only that index gets a non-zero result.
#    If two or more zeros exist, every index returns 0.
#
# 5. Identifying prefix/suffix problems: look for "result at each index depends
#    on all OTHER elements", range queries (sum/product from i to j), or any
#    problem where brute force is O(n²) nested loops scanning left and right.
#    Common examples: Trapping Rain Water, Subarray Sum Equals K, this problem.
#
# 6. Division is the naive approach (total product / nums[i]) but fails on zeros
#    and is explicitly banned here. Prefix/suffix avoids division entirely.
#
# Complexity: Time O(n) two passes, Space O(1) extra (excluding output array)

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
    # [1 2 3 4] = [24 12 8 6]
    # pre: [1 1 2 6]
    # [1] = [0]
    # key insight: multiple everything from left of i with right of i
    # ret[i] = left * right
    res = [0] * len(nums)   
    # setup prefix
    for i in range(len(nums)):
        if i == 0:
            res[i] = 1
        else:
            res[i] = res[i-1] * nums[i-1]
    # increment and set suffix, start from right because results is built from right, we are now multiplying everything to the right of i so need to build suffix from last elt
    suffix = 1
    for i in range(len(nums)-1, -1, -1):
        res[i] *= suffix
        suffix *= nums[i]
        

    


    return res
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
