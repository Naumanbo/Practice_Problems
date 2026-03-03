# Key Takeaways:
# 1. Climbing stairs is the Fibonacci pattern: ways(n) = ways(n-1) + ways(n-2).
#    At each step you either came from 1 below or 2 below — those are the only
#    two choices, so the total ways is the sum of both sub-problems.
#
# 2. Bottom-up DP (tabulation) with a memo array: initialize base cases
#    memo[1]=1, memo[2]=2, then fill forward. Size the array n+1 so index n
#    is always valid — a common off-by-one mistake is sizing it to n.
#
# 3. The nested loop approach (summing all previous values) is wrong for this
#    problem — it overcounts. Only the previous two values matter, not all of them.
#
# 4. Space can be optimized to O(1) by keeping just two variables instead of
#    the full memo array — you only ever read memo[i-1] and memo[i-2].
#
# 5. Identifying DP problems: if the answer to size n depends on answers to
#    smaller sizes (overlapping subproblems), and recalculating from scratch
#    is wasteful (optimal substructure) — it's DP. Store results, don't recompute.
#
# Complexity: Time O(n), Space O(n) with memo array / O(1) with two variables

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
    # create memo
    if n == -1:
        return 0
    if n <=3:
        return n
    
    memo = [0] * (n+1)
    
    memo[0] = 0
    memo[1] = 1
    memo[2] = 2

    for i in range(3,n+1):
        # print("creating ", i)
        memo[i] = memo[i - 1] + memo[i-2]
        # print("curr steps: ", memo[i])

    return memo[n]



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
        {"n": 25, "expected": 121393,     "desc": "twenty-five steps"},
        {"n": 30, "expected": 1346269,    "desc": "thirty steps"},
        {"n": 35, "expected": 14930352,   "desc": "thirty-five steps"},
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
