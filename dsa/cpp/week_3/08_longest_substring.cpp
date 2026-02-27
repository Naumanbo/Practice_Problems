// Tests: Sliding window, hash map/set, two pointers
//
// Longest Substring Without Repeating Characters (LeetCode #3)
// Find the length of the longest substring without repeating characters.

#include <iostream>
#include <vector>
#include <string>
using namespace std;

int lengthOfLongestSubstring(string s) {
    return 0;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    string s;
    int expected;
    string desc;
};

int main() {
    vector<TestCase> tests = {
        {"abcabcbb",   3, "basic case"},
        {"bbbbb",      1, "all same char"},
        {"",           0, "empty string"},
        {"pwwkew",     3, "repeat in middle"},
        {"a",          1, "single char"},
        {"au",         2, "two unique"},
        {"dvdf",       3, "overlap window"},
        {"abcdefg",    7, "all unique"},
        {"aab",        2, "repeat at start"},
        {"tmmzuxt",    5, "repeat at both ends"},
        {" ",          1, "single space"},
        {"abba",       2, "palindrome"},
        {"abcbda",     4, "complex window"},
        {"aababcabcd", 4, "increasing unique suffix"},
        {"abcdeabcde", 5, "repeated block"},
        {"ohvhjdml",   6, "long unique suffix"},
    };

    cout << "======================================================================" << endl;
    cout << "LONGEST SUBSTRING - Test Results" << endl;
    cout << "======================================================================" << endl;

    int passed = 0;
    for (size_t i = 0; i < tests.size(); i++) {
        int result = lengthOfLongestSubstring(tests[i].s);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        cout << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tests[i].desc
             << ": \"" << tests[i].s << "\" -> " << result << endl;
    }

    cout << "======================================================================" << endl;
    cout << "Summary: " << passed << "/" << tests.size() << " passed" << endl;
    cout << "======================================================================" << endl;
    return 0;
}
