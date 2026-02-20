# Key Takeaways:
# 1. Classic single-pass greedy: track the minimum price seen so far, and at each
#    step compute the profit if you sold today. No need to look ahead.
#
# 2. Use `sys.maxsize` or `float('inf')` to initialize a minimum tracker.
#    `float('inf')` is idiomatic for DSA — works in comparisons with ints directly.
#
# 3. The `elif` is key — you can't sell on the same day you update the minimum,
#    so the two cases (new min vs. new profit) are mutually exclusive.
#
# 4. Return 0 if no profit is possible — initialize max_profit to 0 so unprofitable
#    cases naturally return 0 without extra checks.
#
# Complexity: Time O(n), Space O(1)

"""
Best Time to Buy and Sell Stock (LeetCode 121)

Tests: Arrays, single pass, tracking minimum

You are given an array prices where prices[i] is the price of a stock on the ith day.
You want to maximize profit by choosing a single day to buy and a single day to sell.
Return the maximum profit. If no profit is possible, return 0.

Note: You cannot sell before you buy.

Example 1:
    Input: prices = [7,1,5,3,6,4]
    Output: 5
    Explanation: Buy on day 2 (price=1), sell on day 5 (price=6), profit = 6-1 = 5

Example 2:
    Input: prices = [7,6,4,3,1]
    Output: 0
    Explanation: No profitable transaction possible

Constraints:
    - 1 <= prices.length <= 10^5
    - 0 <= prices[i] <= 10^4
"""
import sys
from typing import List


def max_profit(prices: List[int]) -> int:
    # TODO: Implement solution
    # Hint: Track the minimum price seen so far and the maximum profit
    maximum_profit = 0
    min_price = sys.maxsize
    for price in prices:
        if price < min_price:
            min_price = price
        elif price - min_price > maximum_profit:
            maximum_profit = price - min_price
    return maximum_profit


# Tests
if __name__ == "__main__":
    test_cases = [
        {"prices": [7, 1, 5, 3, 6, 4], "expected": 5, "desc": "basic case"},
        {"prices": [7, 6, 4, 3, 1], "expected": 0, "desc": "decreasing prices"},
        {"prices": [1, 2, 3, 4, 5], "expected": 4, "desc": "increasing prices"},
        {"prices": [2, 4, 1], "expected": 2, "desc": "buy early sell middle"},
        {"prices": [3, 3, 3], "expected": 0, "desc": "flat prices"},
        {"prices": [1], "expected": 0, "desc": "single element"},
        {"prices": [2, 1], "expected": 0, "desc": "two elements decreasing"},
        {"prices": [1, 2], "expected": 1, "desc": "two elements increasing"},
        {"prices": [10, 1, 10], "expected": 9, "desc": "V-shape recovery"},
        {"prices": [3, 8, 2, 5], "expected": 5, "desc": "peak then valley"},
        {"prices": [1, 5, 2, 8], "expected": 7, "desc": "multiple peaks, best at end"},
        {"prices": [5, 4, 3, 2, 1, 10], "expected": 9, "desc": "min at end then spike"},
        {"prices": [0, 0, 0], "expected": 0, "desc": "all zeros"},
        {"prices": [5, 5], "expected": 0, "desc": "two same"},
        {"prices": [3, 1, 100, 2, 4], "expected": 99, "desc": "spike in middle"},
        {"prices": [1, 10, 1, 10, 1], "expected": 9, "desc": "oscillating"},
        {"prices": [0, 1], "expected": 1, "desc": "zero then one"},
        {"prices": [10000, 1, 10000], "expected": 9999, "desc": "near constraint max"},
        {"prices": [2, 1, 2, 0, 1], "expected": 1, "desc": "multiple dips"},
        {"prices": [1, 1, 1, 1, 2], "expected": 1, "desc": "flat then rise"},
    ]

    all_passed = True
    for tc in test_cases:
        result = max_profit(tc["prices"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: max_profit({tc['prices']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
