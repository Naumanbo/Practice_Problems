// Compile: gcc -o 02_valid_anagram 02_valid_anagram.c
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

bool isAnagram(const char* s, const char* t) {
    // Your implementation
    return false;
}

struct TestCase {
    const char* s;
    const char* t;
    bool expected;
    const char* desc;
};

int main() {
    struct TestCase tests[] = {
        {"anagram", "nagaram", true, "classic anagram"},
        {"rat", "car", false, "not anagram"},
        {"listen", "silent", true, "listen/silent"},
        {"hello", "world", false, "different letters"},
        {"a", "a", true, "single char same"},
        {"a", "b", false, "single char different"},
        {"ab", "ba", true, "two chars swapped"},
        {"ab", "cd", false, "two chars different"},
        {"abc", "ab", false, "different lengths"},
        {"a", "ab", false, "subset string"},
        {"aaa", "aaa", true, "all same chars"},
        {"aab", "baa", true, "repeated chars anagram"},
        {"aacc", "ccac", false, "same chars wrong count"},
        {"aabb", "abab", true, "interleaved duplicates"},
        {"abcde", "abcdf", false, "one char different"},
        {"abcd", "abce", false, "last char differs"},
        {"aaab", "aaba", true, "rearranged with repeats"},
        {"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true, "full alphabet reversed"},
        {"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyy", false, "full alphabet one off"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);
    int passed = 0;

    printf("======================================================================\n");
    printf("VALID ANAGRAM - Test Results\n");
    printf("======================================================================\n\n");

    for (int i = 0; i < numTests; i++) {
        bool result = isAnagram(tests[i].s, tests[i].t);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        printf("%2d. [%s] %s: '%s','%s' -> %s\n",
               i + 1, ok ? "PASS" : "FAIL", tests[i].desc,
               tests[i].s, tests[i].t, result ? "true" : "false");
    }

    printf("\n======================================================================\n");
    printf("Summary: %d/%d passed\n", passed, numTests);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. Why is int count[26] sufficient for lowercase English?\n");
    printf("2. How would you handle Unicode in C?\n");
    printf("3. Could you solve by sorting? Trade-off?\n");
    return 0;
}
