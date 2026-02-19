/*
DSA Problem 2: Valid Parentheses

Tests: Stack data structure, matching pairs, string traversal

Difficulty: Easy
Source: LeetCode #20

Problem:
Given a string s containing just '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

Valid if:
1. Open brackets closed by same type
2. Open brackets closed in correct order
3. Every close bracket has corresponding open bracket

Compile: g++ -std=c++17 -o 02_valid_parentheses 02_valid_parentheses.cpp
Run: ./02_valid_parentheses
*/

#include <iostream>
#include <string>
#include <stack>
#include <vector>
#include <map>

using namespace std;

// Time: O(n)
// Space: O(n)
// Hint: Use std::stack
bool isValid(const string& s) {
    // Your implementation
    map<char, char> closers;
    closers['{'] = '}';
    closers['('] = ')';
    closers['['] = ']';

    stack<char> openers;
    for (char c : s) {
        if (c == '{' || c == '[' || c == '(') {
            openers.push(c);
        }
        else {
            if (!openers.empty() && c == closers[openers.top()]) {
                openers.pop();
            }
            else {
                return false;
            }
        }
    }

    if (openers.empty()) {
        return true;
    }
    return false;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    string s;
    bool expected;
    string desc;
};

string repeat(const string& str, int n) {
    string result;
    for (int i = 0; i < n; i++) result += str;
    return result;
}

int main() {
    vector<TestCase> testCases = {
        // Basic valid cases
        {"()", true, "Single pair - parentheses"},
        {"[]", true, "Single pair - brackets"},
        {"{}", true, "Single pair - braces"},
        {"()[]{}", true, "Multiple pairs sequential"},
        {"{[]}", true, "Nested brackets"},
        {"([{}])", true, "Deeply nested"},

        // Basic invalid cases
        {"(]", false, "Mismatched types"},
        {"([)]", false, "Wrong order - interleaved"},
        {"{[}]", false, "Wrong order - interleaved v2"},

        // Edge cases - empty
        {"", true, "Empty string"},

        // Edge cases - single bracket
        {"(", false, "Single open paren"},
        {")", false, "Single close paren"},
        {"[", false, "Single open bracket"},
        {"}", false, "Single close brace"},

        // Edge cases - unbalanced
        {"(()", false, "Extra open at start"},
        {"())", false, "Extra close at end"},
        {"(())", true, "Balanced nested"},
        {"((()))", true, "Triple nested"},

        // Edge cases - long strings
        {repeat("()", 100), true, "Long valid string"},
        {repeat("(", 50) + repeat(")", 50), true, "Many nested"},
        {repeat("(", 50) + repeat(")", 49), false, "Off by one"},

        // Edge cases - complex patterns
        {"{[()()]}", true, "Complex valid"},
        {"[({})]", true, "All types nested"},
        {"[(])", false, "Complex invalid"},

        // Edge case - close before open
        {")(", false, "Close before open"},
        {"}{", false, "Close before open v2"},
    };

    cout << "======================================================================" << endl;
    cout << "VALID PARENTHESES - Test Results" << endl;
    cout << "======================================================================" << endl;

    int passed = 0;
    int total = testCases.size();

    for (size_t i = 0; i < testCases.size(); i++) {
        auto& tc = testCases[i];
        bool result = isValid(tc.s);
        bool ok = (result == tc.expected);
        if (ok) passed++;

        string status = ok ? "PASS" : "FAIL";
        string displayS = tc.s.length() <= 20 ? tc.s : tc.s.substr(0, 17) + "...";

        cout << (i + 1) << ". [" << status << "] " << tc.desc << endl;
        cout << "    Input: '" << displayS << "' | Got: " << (result ? "true" : "false")
             << " | Expected: " << (tc.expected ? "true" : "false") << endl;
    }

    cout << "\n======================================================================" << endl;
    cout << "Summary: " << passed << "/" << total << " passed" << endl;
    cout << "======================================================================" << endl;
    cout << "\nQuestions:" << endl;
    cout << "1. Why is a stack the right data structure for this problem?" << endl;
    cout << "2. What's the time/space complexity?" << endl;
    cout << "3. Could you solve this without a stack? What's the tradeoff?" << endl;

    return 0;
}
