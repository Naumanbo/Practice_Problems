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

from typing import List


def max_profit(prices: List[int]) -> int:
    # TODO: Implement solution
    # Hint: Track the minimum price seen so far and the maximum profit
    pass


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
