"""
DSA Problem: Climbing Stairs

Tests: Dynamic programming, memoization, bottom-up DP, Fibonacci pattern

Difficulty: Easy
Source: LeetCode #70

Problem:
You are climbing a staircase with n steps. Each time you can climb 1 or 2 steps.
In how many distinct ways can you climb to the top?

Constraints:
    - 1 <= n <= 45
"""


def climb_stairs(n: int) -> int:
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"n": 1,  "expected": 1,          "desc": "one step"},
        {"n": 2,  "expected": 2,          "desc": "two steps"},
        {"n": 3,  "expected": 3,          "desc": "three steps"},
        {"n": 4,  "expected": 5,          "desc": "four steps"},
        {"n": 5,  "expected": 8,          "desc": "five steps"},
        {"n": 6,  "expected": 13,         "desc": "six steps"},
        {"n": 7,  "expected": 21,         "desc": "seven steps"},
        {"n": 8,  "expected": 34,         "desc": "eight steps"},
        {"n": 9,  "expected": 55,         "desc": "nine steps"},
        {"n": 10, "expected": 89,         "desc": "ten steps"},
        {"n": 15, "expected": 987,        "desc": "fifteen steps"},
        {"n": 20, "expected": 10946,      "desc": "twenty steps"},
        {"n": 25, "expected": 196418,     "desc": "twenty-five steps"},
        {"n": 30, "expected": 1346269,    "desc": "thirty steps"},
        {"n": 35, "expected": 9227465,    "desc": "thirty-five steps"},
        {"n": 45, "expected": 1836311903, "desc": "max constraint"},
    ]

    all_passed = True
    for tc in test_cases:
        result = climb_stairs(tc["n"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: climb_stairs({tc['n']}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
