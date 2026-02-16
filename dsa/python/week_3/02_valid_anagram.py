"""
DSA Problem: Valid Anagram

Tests: Hash maps, character counting, sorting approach

Difficulty: Easy
Source: LeetCode #242

Problem:
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Constraints:
    - 1 <= s.length, t.length <= 5 * 10^4
    - s and t consist of lowercase English letters
"""


def is_anagram_sort(s: str, t: str) -> bool:
    """
    Sorting approach.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


def is_anagram_hashmap(s: str, t: str) -> bool:
    """
    Hash map counting approach.

    Time: O(?)
    Space: O(?)

    Your implementation:
    """
    pass


# =============================================================================
# Test Cases - LeetCode Level
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {"s": "anagram", "t": "nagaram", "expected": True, "desc": "classic anagram"},
        {"s": "rat", "t": "car", "expected": False, "desc": "not anagram"},
        {"s": "listen", "t": "silent", "expected": True, "desc": "listen/silent"},
        {"s": "hello", "t": "world", "expected": False, "desc": "different letters"},
        {"s": "a", "t": "a", "expected": True, "desc": "single char same"},
        {"s": "a", "t": "b", "expected": False, "desc": "single char different"},
        {"s": "ab", "t": "ba", "expected": True, "desc": "two chars swapped"},
        {"s": "ab", "t": "cd", "expected": False, "desc": "two chars different"},
        {"s": "abc", "t": "ab", "expected": False, "desc": "different lengths"},
        {"s": "a", "t": "ab", "expected": False, "desc": "subset string"},
        {"s": "aaa", "t": "aaa", "expected": True, "desc": "all same chars"},
        {"s": "aab", "t": "baa", "expected": True, "desc": "repeated chars anagram"},
        {"s": "aacc", "t": "ccac", "expected": False, "desc": "same chars wrong count"},
        {"s": "aabb", "t": "abab", "expected": True, "desc": "interleaved duplicates"},
        {"s": "abcde", "t": "abcdf", "expected": False, "desc": "one char different"},
        {"s": "abcd", "t": "abce", "expected": False, "desc": "last char differs"},
        {"s": "aaab", "t": "aaba", "expected": True, "desc": "rearranged with repeats"},
        {"s": "a" * 100, "t": "a" * 100, "expected": True, "desc": "long same string"},
        {"s": "a" * 99 + "b", "t": "b" + "a" * 99, "expected": True, "desc": "long anagram"},
        {"s": "a" * 100, "t": "a" * 99 + "b", "expected": False, "desc": "long near miss"},
        {"s": "abcdefghijklmnopqrstuvwxyz", "t": "zyxwvutsrqponmlkjihgfedcba", "expected": True, "desc": "full alphabet reversed"},
    ]

    print("=" * 70)
    print("VALID ANAGRAM - Test Results")
    print("=" * 70)

    for approach_name, func in [("Sorting", is_anagram_sort), ("HashMap", is_anagram_hashmap)]:
        print(f"\n{approach_name}:")
        passed = 0
        total = len(test_cases)
        for i, tc in enumerate(test_cases, 1):
            result = func(tc["s"], tc["t"])
            ok = result == tc["expected"]
            passed += ok
            status = "PASS" if ok else "FAIL"
            s_display = tc["s"] if len(tc["s"]) <= 15 else tc["s"][:12] + "..."
            t_display = tc["t"] if len(tc["t"]) <= 15 else tc["t"][:12] + "..."
            print(f"  {i:2}. [{status}] {tc['desc']}: '{s_display}','{t_display}' -> {result}")
        print(f"  Result: {passed}/{total}")

    print("\n" + "=" * 70)
    print("Questions:")
    print("1. Which approach is better for follow-up: 'What if inputs contain Unicode?'")
    print("2. Can you solve this with a single counter (increment for s, decrement for t)?")
    print("3. What's the space complexity of sorting approach? (Hint: depends on sort algorithm)")
