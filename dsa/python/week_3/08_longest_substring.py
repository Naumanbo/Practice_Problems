"""
DSA Problem: Longest Substring Without Repeating Characters

Tests: Sliding window, hash map/set, two pointers

Difficulty: Medium
Source: LeetCode #3

Problem:
Given a string s, find the length of the longest substring without repeating characters.

Constraints:
    - 0 <= s.length <= 5 * 10^4
    - s consists of English letters, digits, symbols and spaces
"""


def length_of_longest_substring(s: str) -> int:
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"s": "abcabcbb", "expected": 3, "desc": "basic case"},
        {"s": "bbbbb",    "expected": 1, "desc": "all same char"},
        {"s": "",         "expected": 0, "desc": "empty string"},
        {"s": "pwwkew",   "expected": 3, "desc": "repeat in middle"},
        {"s": "a",        "expected": 1, "desc": "single char"},
        {"s": "au",       "expected": 2, "desc": "two unique"},
        {"s": "dvdf",     "expected": 3, "desc": "overlap window"},
        {"s": "abcdefg",  "expected": 7, "desc": "all unique"},
        {"s": "aab",      "expected": 2, "desc": "repeat at start"},
        {"s": "tmmzuxt",  "expected": 5, "desc": "repeat at both ends"},
        {"s": " ",        "expected": 1, "desc": "single space"},
        {"s": "abba",     "expected": 2, "desc": "palindrome"},
        {"s": "abcbda",   "expected": 4, "desc": "complex window"},
        {"s": "aababcabcd","expected": 4, "desc": "increasing unique suffix"},
        {"s": "abcdeabcde","expected": 5, "desc": "repeated block"},
        {"s": "ohvhjdml", "expected": 6, "desc": "long unique suffix"},
    ]

    all_passed = True
    for tc in test_cases:
        result = length_of_longest_substring(tc["s"])
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: length_of_longest_substring({tc['s']!r}) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
