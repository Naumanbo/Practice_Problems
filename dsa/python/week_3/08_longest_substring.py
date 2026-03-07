# Key Takeaways:
# 1. Sliding window invariant: the data structure (set/list) always mirrors EXACTLY
#    the characters in s[left:right+1]. Every time right expands, you add s[right].
#    Every time left contracts, you remove s[left]. Never rebuild from scratch —
#    maintaining the invariant incrementally is what makes the algorithm O(n).
#
# 2. The while loop on collision, not an if: when s[right] is already in the window,
#    you may need to remove multiple characters from the left before the duplicate
#    is gone. An `if` only removes one character, leaving the window invalid if the
#    duplicate was not at s[left]. The `while` keeps shrinking until the window is
#    clean again. This is the core correctness requirement of the algorithm.
#
# 3. Use a set, not a list: my implementation uses a list with .remove(), which is
#    O(n) per removal — it scans the list to find the element. A set gives O(1)
#    membership check (the `in` test) and O(1) discard. The total pointer moves are
#    still O(n), but each operation is O(1) instead of O(n):
#      window = set()
#      window.discard(s[left])   # O(1), no error if key missing
#      window.add(s[right])      # O(1)
#
# 4. Off-by-one on initialization: seeding the window with s[0] and starting right
#    at index 1 is correct, but subtle. The alternative is starting right at 0
#    with an empty window and no special-case guard. Either works as long as s[0]
#    is in the window exactly once before right reaches index 1.
#
# 5. Edge case — len(s) < 2 returns len(s): this handles both "" → 0 and "a" → 1.
#    Without this guard, a single-character string would return 0 because the for
#    loop never executes (range(1,1) is empty) and maxLength stays 0, even though
#    s[0] is correctly seeded in the window. Always trace your initialization through
#    the smallest inputs.
#
# 6. Update maxLength AFTER appending s[right]: at that moment the window is fully
#    valid (no duplicates) and at its maximum size for this value of right. Updating
#    before the append misses counting the current character.
#
# 7. Identifying sliding window problems: the signal is "contiguous subarray/substring"
#    + a validity condition that can be maintained incrementally (uniqueness, sum ≤ k,
#    at most k distinct chars, no repeats). If the brute-force approach is O(n²) nested
#    loops scanning every possible window, a sliding window eliminates the inner loop
#    by never re-examining characters that are still valid. Common variants:
#      - Fixed window size: slide right and left together (both move +1 each step)
#      - Variable window: expand right greedily, shrink left on violation (this problem)
#
# 8. Space is O(min(n, |alphabet|)): the window can never hold more characters than
#    exist in the full alphabet. For ASCII that cap is 128; for lowercase letters it
#    is 26. This means for very long strings over a small alphabet, the set stays tiny
#    regardless of input size — the space bound is alphabet-size-limited, not n-limited.
#
# Related EECS 281 Lectures: Lec 13 (Strings & Sequences — substring problems,
#   sliding window over character streams), Lec 15 (Hash Tables — O(1) set
#   membership for tracking seen characters in the current window)
# Complexity: Time O(n), Space O(min(n, alphabet_size))

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

    # Sliding window?
    if len(s) < 2:
        return len(s)
    left = 0
    substring = []
    substring.append(s[left])
    maxLength = 0

    for right in range(1, len(s)):
        # print(substring, 's[right]: ', s[right], "s[left]: ", s[left])
        while s[right] in substring:
            substring.remove(s[left])
            left += 1
        substring.append(s[right])
        if len(substring) > maxLength:
            maxLength = len(substring)
    
    
    return maxLength





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
