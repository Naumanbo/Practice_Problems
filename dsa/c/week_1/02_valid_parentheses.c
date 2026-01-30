/*
DSA Problem 2: Valid Parentheses

Tests: Stack data structure (manual implementation), matching pairs, string traversal

Difficulty: Easy
Source: LeetCode #20

Problem:
Given a string s containing just '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

Valid if:
1. Open brackets closed by same type
2. Open brackets closed in correct order
3. Every close bracket has corresponding open bracket

Compile: gcc -o 02_valid_parentheses 02_valid_parentheses.c
Run: ./02_valid_parentheses
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

// Time: O(?)
// Space: O(?)
// Hint: Implement a stack using an array
bool isValid(const char* s) {
    // Your implementation
    return false;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

typedef struct {
    const char* s;
    bool expected;
    const char* desc;
} TestCase;

char* repeat(const char* str, int n) {
    int len = strlen(str);
    char* result = malloc(len * n + 1);
    result[0] = '\0';
    for (int i = 0; i < n; i++) {
        strcat(result, str);
    }
    return result;
}

char* repeatChar(char c, int n) {
    char* result = malloc(n + 1);
    for (int i = 0; i < n; i++) {
        result[i] = c;
    }
    result[n] = '\0';
    return result;
}

int main() {
    // Generate long test strings
    char* longValid = repeat("()", 100);
    char* manyNested = malloc(101);
    char* offByOne = malloc(100);
    strcpy(manyNested, repeatChar('(', 50));
    strcat(manyNested, repeatChar(')', 50));
    strcpy(offByOne, repeatChar('(', 50));
    strcat(offByOne, repeatChar(')', 49));

    TestCase testCases[] = {
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
        {longValid, true, "Long valid string"},
        {manyNested, true, "Many nested"},
        {offByOne, false, "Off by one"},

        // Edge cases - complex patterns
        {"{[()()]}", true, "Complex valid"},
        {"[({})]", true, "All types nested"},
        {"[(])", false, "Complex invalid"},

        // Edge case - close before open
        {")(", false, "Close before open"},
        {"}{", false, "Close before open v2"},
    };

    int total = sizeof(testCases) / sizeof(TestCase);

    printf("======================================================================\n");
    printf("VALID PARENTHESES - Test Results\n");
    printf("======================================================================\n");

    int passed = 0;

    for (int i = 0; i < total; i++) {
        TestCase tc = testCases[i];
        bool result = isValid(tc.s);
        bool ok = (result == tc.expected);
        if (ok) passed++;

        const char* status = ok ? "PASS" : "FAIL";

        // Truncate display string if too long
        char displayS[25];
        if (strlen(tc.s) <= 20) {
            strcpy(displayS, tc.s);
        } else {
            strncpy(displayS, tc.s, 17);
            strcpy(displayS + 17, "...");
        }

        printf("%2d. [%s] %s\n", i + 1, status, tc.desc);
        printf("    Input: '%s' | Got: %s | Expected: %s\n",
               displayS,
               result ? "true" : "false",
               tc.expected ? "true" : "false");
    }

    printf("\n======================================================================\n");
    printf("Summary: %d/%d passed\n", passed, total);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. Why is a stack the right data structure for this problem?\n");
    printf("2. How did you implement the stack in C?\n");
    printf("3. What's the time/space complexity?\n");

    // Free allocated memory
    free(longValid);
    free(manyNested);
    free(offByOne);

    return 0;
}
