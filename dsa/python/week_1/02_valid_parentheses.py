"""
DSA Problem 2: Valid Parentheses

Tests: Stack data structure, matching pairs, string traversal

Difficulty: Easy
Source: LeetCode #20

Problem:
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

Valid if:
1. Open brackets closed by same type
2. Open brackets closed in correct order
3. Every close bracket has corresponding open bracket
"""


def is_valid(s: str) -> bool:
    """
    Time: O(?)
    Space: O(?)

    Hint: Use a stack. Push opening brackets, pop and match for closing.

    Your implementation:
    """
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        # Basic valid cases
        {"s": "()", "expected": True, "desc": "Single pair - parentheses"},
        {"s": "[]", "expected": True, "desc": "Single pair - brackets"},
        {"s": "{}", "expected": True, "desc": "Single pair - braces"},
        {"s": "()[]{}", "expected": True, "desc": "Multiple pairs sequential"},
        {"s": "{[]}", "expected": True, "desc": "Nested brackets"},
        {"s": "([{}])", "expected": True, "desc": "Deeply nested"},

        # Basic invalid cases
        {"s": "(]", "expected": False, "desc": "Mismatched types"},
        {"s": "([)]", "expected": False, "desc": "Wrong order - interleaved"},
        {"s": "{[}]", "expected": False, "desc": "Wrong order - interleaved v2"},

        # Edge cases - empty
        {"s": "", "expected": True, "desc": "Empty string"},

        # Edge cases - single bracket
        {"s": "(", "expected": False, "desc": "Single open paren"},
        {"s": ")", "expected": False, "desc": "Single close paren"},
        {"s": "[", "expected": False, "desc": "Single open bracket"},
        {"s": "}", "expected": False, "desc": "Single close brace"},

        # Edge cases - unbalanced
        {"s": "(()", "expected": False, "desc": "Extra open at start"},
        {"s": "())", "expected": False, "desc": "Extra close at end"},
        {"s": "(())", "expected": True, "desc": "Balanced nested"},
        {"s": "((()))", "expected": True, "desc": "Triple nested"},

        # Edge cases - long strings
        {"s": "()" * 100, "expected": True, "desc": "Long valid string"},
        {"s": "(" * 50 + ")" * 50, "expected": True, "desc": "Many nested"},
        {"s": "(" * 50 + ")" * 49, "expected": False, "desc": "Off by one"},

        # Edge cases - complex patterns
        {"s": "{[()()]}", "expected": True, "desc": "Complex valid"},
        {"s": "[({})]", "expected": True, "desc": "All types nested"},
        {"s": "[(])", "expected": False, "desc": "Complex invalid"},

        # Edge case - close before open
        {"s": ")(", "expected": False, "desc": "Close before open"},
        {"s": "}{", "expected": False, "desc": "Close before open v2"},
    ]

    print("=" * 70)
    print("VALID PARENTHESES - Test Results")
    print("=" * 70)

    passed = 0
    total = len(test_cases)

    for i, tc in enumerate(test_cases, 1):
        s, expected, desc = tc["s"], tc["expected"], tc["desc"]
        result = is_valid(s)
        ok = result == expected
        passed += ok
        status = "PASS" if ok else "FAIL"

        display_s = s if len(s) <= 20 else s[:17] + "..."
        print(f"{i:2}. [{status}] {desc}")
        print(f"    Input: '{display_s}' | Got: {result} | Expected: {expected}")

    print("\n" + "=" * 70)
    print(f"Summary: {passed}/{total} passed")
    print("=" * 70)
    print("\nQuestions:")
    print("1. Why is a stack the right data structure for this problem?")
    print("2. What's the time/space complexity?")
    print("3. Could you solve this without a stack? What's the tradeoff?")
