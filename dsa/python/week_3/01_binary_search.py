# Key Takeaways:
# 1. Binary search template — three rules to never break:
#    - `while left <= right` (not <) — otherwise you skip the single-element case
#    - `mid = left + (right - left) // 2` (not `(right - left) // 2`) — the latter
#      gives an offset, not an index. The `left +` form is also overflow-safe in C++/Go.
#    - `left = mid + 1` / `right = mid - 1` — always skip mid after checking it,
#      never `left = mid` or `right = mid` which causes infinite loops.
#
# 2. For recursive binary search with a fixed public API (nums, target), use the
#    helper function pattern: define an inner function that carries left/right as
#    parameters, and call it with (0, len(nums) - 1) from the outer function.
#
# 3. Recursive base case is `if left > right: return -1` — the inverse of the
#    iterative while condition. Same logic, different framing.
#
# Complexity: Time O(log n), Space O(1) iterative / O(log n) recursive (call stack)

"""
DSA Problem: Binary Search

Tests: Binary search, sorted array, O(log n) search, iterative vs recursive

Difficulty: Easy
Source: LeetCode #704

Problem:
Given a sorted array of integers nums and a target value, return the index
if target is found. If not, return -1.

Constraints:
    - 1 <= nums.length <= 10^4
    - -10^4 < nums[i], target < 10^4
    - All integers in nums are unique
    - nums is sorted in ascending order
"""

from typing import List


def binary_search_iterative(nums: List[int], target: int) -> int:
    
    """
    Iterative binary search.

    Time: O(log n)
    Space: O(1)

    Your implementation:
    """
    left = 0
    right = len(nums) - 1

    while left <= right:
        mid = left + (right - left) // 2
        if target > nums[mid]: # 
            left = mid + 1
        elif target < nums[mid]:
            right = mid - 1
        else:
            return mid
    
    return -1


def binary_search_recursive(nums: List[int], target: int) -> int:
    """
    Recursive binary search.

    Time: O(log n)
    Space: O(log n) - consider call stack

    Your implementation:
    """
    def helper(left, right):
        mid = left + (right - left) // 2
        if target == nums[mid]: # base case
            return mid
        
        if left <= right:
            if target > nums[mid]: # recursive case: target resides on right of mid
                return helper(mid+1, right)
            elif target < nums[mid]: # recursive case 2: target resides on left of mid
                return helper(left, mid-1)
        return -1 # base case 2: target not found
    
    return helper(0, len(nums) - 1) # start recursive calls
            


# =============================================================================
# Test Cases - LeetCode Level
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"nums": [-1, 0, 3, 5, 9, 12], "target": 9, "expected": 4, "desc": "target in middle"},
        {"nums": [-1, 0, 3, 5, 9, 12], "target": 2, "expected": -1, "desc": "target not found"},
        {"nums": [1, 2, 3, 4, 5], "target": 1, "expected": 0, "desc": "target at start"},
        {"nums": [1, 2, 3, 4, 5], "target": 5, "expected": 4, "desc": "target at end"},
        {"nums": [1, 2, 3, 4, 5], "target": 3, "expected": 2, "desc": "target in center"},
        {"nums": [5], "target": 5, "expected": 0, "desc": "single element found"},
        {"nums": [5], "target": 3, "expected": -1, "desc": "single element not found"},
        {"nums": [1, 2], "target": 1, "expected": 0, "desc": "two elements first"},
        {"nums": [1, 2], "target": 2, "expected": 1, "desc": "two elements second"},
        {"nums": [1, 2], "target": 3, "expected": -1, "desc": "two elements not found"},
        {"nums": [-10, -5, 0, 5, 10], "target": -5, "expected": 1, "desc": "negative target"},
        {"nums": [-10, -5, 0, 5, 10], "target": 0, "expected": 2, "desc": "zero target"},
        {"nums": [-100, -50, -10, -1], "target": -100, "expected": 0, "desc": "all negative first"},
        {"nums": [-100, -50, -10, -1], "target": -1, "expected": 3, "desc": "all negative last"},
        {"nums": [-100, -50, -10, -1], "target": 5, "expected": -1, "desc": "all negative not found"},
        {"nums": [1, 2, 3, 4, 5], "target": 0, "expected": -1, "desc": "target below range"},
        {"nums": [1, 2, 3, 4, 5], "target": 6, "expected": -1, "desc": "target above range"},
        {"nums": list(range(1, 101)), "target": 50, "expected": 49, "desc": "larger array middle"},
        {"nums": list(range(1, 101)), "target": 1, "expected": 0, "desc": "larger array first"},
        {"nums": list(range(1, 101)), "target": 100, "expected": 99, "desc": "larger array last"},
        {"nums": list(range(1, 101)), "target": 101, "expected": -1, "desc": "larger array not found"},
        {"nums": [-9999, 0, 9999], "target": 9999, "expected": 2, "desc": "near constraint bounds"},
    ]

    def run_tests(func, name):
        print(f"\n{name}:")
        passed = 0
        for i, tc in enumerate(test_cases, 1):
            result = func(tc["nums"][:], tc["target"])
            ok = result == tc["expected"]
            passed += ok
            status = "PASS" if ok else "FAIL"
            print(f"  {i:2}. [{status}] {tc['desc']}: target={tc['target']} -> {result} (expected {tc['expected']})")
        return passed

    print("=" * 70)
    print("BINARY SEARCH - Test Results")
    print("=" * 70)

    total = len(test_cases)
    iter_passed = run_tests(binary_search_iterative, "Iterative")
    rec_passed = run_tests(binary_search_recursive, "Recursive")

    print("\n" + "=" * 70)
    print(f"Summary: Iterative {iter_passed}/{total} | Recursive {rec_passed}/{total}")
    print("=" * 70)
    print("\nQuestions:")
    print("1. Why must the array be sorted for binary search?")
    print("2. What happens if you use (left + right) / 2 vs left + (right - left) // 2?")
    print("3. What's the relationship between binary search and bisect module?")
