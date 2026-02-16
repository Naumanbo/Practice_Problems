/*
DSA Problem 1: Two Sum

Tests: Hash map usage (manual implementation), time/space complexity, array traversal

Difficulty: Easy
Source: LeetCode #1

Problem:
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

Constraints:
  - 2 <= numsSize <= 10^4
  - -10^9 <= nums[i] <= 10^9
  - Only one valid answer exists.
  - You may not use the same element twice.

Compile: gcc -o 01_two_sum 01_two_sum.c
Run: ./01_two_sum
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>


// Result structure (since C can't return arrays easily)
typedef struct {
    int indices[2];
    bool found;
} TwoSumResult;

// Brute force approach
// Time: O(?)
// Space: O(?)
TwoSumResult twoSumBrute(int* nums, int numsSize, int target) {
    TwoSumResult result = {{-1, -1}, false};
    // Your implementation
    for (int i = 0; i < numsSize; i++) {
        for (int j = i +1; j < numsSize; j++) {
            if (target == nums[j] + nums[i]) {
                result = (TwoSumResult){{i,j}, true};
                return result;
            }
        }
    }
    return result;
}

// Optimal approach - Note: In C, you'd implement a hash table manually
// For simplicity, you can use the brute force or implement a basic hash
// Time: O(?)
// Space: O(?)
TwoSumResult twoSumOptimal(int* nums, int numsSize, int target) {
    TwoSumResult result = {{-1, -1}, false};
    // Your implementation (hash table or sorted two-pointer)
    map
    for (int i = 0; i < numsSize; i++) {
        for (int j = i + 1; j < numsSize; j++) {

        }
    }
    return result;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

typedef struct {
    int* nums;
    int numsSize;
    int target;
    const char* desc;
} TestCase;

bool validate(TwoSumResult result, int* nums, int numsSize, int target) {
    if (!result.found) return false;
    int i = result.indices[0];
    int j = result.indices[1];
    if (i == j) return false;
    if (i < 0 || i >= numsSize || j < 0 || j >= numsSize) return false;
    return nums[i] + nums[j] == target;
}

void printArray(int* arr, int size) {
    printf("[");
    for (int i = 0; i < size && i < 8; i++) {
        printf("%d", arr[i]);
        if (i < size - 1 && i < 7) printf(", ");
    }
    if (size > 8) printf("...");
    printf("]");
}

void printResult(TwoSumResult r) {
    if (r.found) {
        printf("[%d, %d]", r.indices[0], r.indices[1]);
    } else {
        printf("[]");
    }
}

int main() {
    // Test data
    int t1[] = {2, 7, 11, 15};
    int t2[] = {3, 2, 4};
    int t3[] = {3, 3};
    int t4[] = {1, 2};
    int t5[] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    int t6[] = {-1, -2, -3, -4};
    int t7[] = {-3, 4, 3, 90};
    int t8[] = {5, -5, 10};
    int t9[] = {0, 4, 3, 0};
    int t10[] = {0, 1, 2};
    int t11[] = {1000000000, 2, 1000000000};
    int t12[] = {1, 5, 8, 3};
    int t13[] = {4, 5, 1, 2};
    int t14[] = {5, 5, 5, 5};

    TestCase testCases[] = {
        {t1, 4, 9, "Basic - first two elements"},
        {t2, 3, 6, "Basic - middle elements"},
        {t3, 2, 6, "Duplicate values"},
        {t4, 2, 3, "Minimum size (2 elements)"},
        {t5, 10, 19, "Larger array"},
        {t6, 4, -6, "All negative"},
        {t7, 4, 0, "Negative + positive = 0"},
        {t8, 3, 0, "Sum to zero"},
        {t9, 4, 0, "Two zeros"},
        {t10, 3, 2, "Zero + positive"},
        {t11, 3, 2000000000, "Large numbers"},
        {t12, 4, 4, "Answer at start and end"},
        {t13, 4, 6, "Answer in middle"},
        {t14, 4, 10, "All same values"},
    };

    int total = sizeof(testCases) / sizeof(TestCase);

    printf("======================================================================\n");
    printf("TWO SUM - Test Results\n");
    printf("======================================================================\n");

    int brutePassed = 0, optimalPassed = 0;

    for (int i = 0; i < total; i++) {
        TestCase tc = testCases[i];

        TwoSumResult bruteResult = twoSumBrute(tc.nums, tc.numsSize, tc.target);
        TwoSumResult optimalResult = twoSumOptimal(tc.nums, tc.numsSize, tc.target);

        bool bruteOk = validate(bruteResult, tc.nums, tc.numsSize, tc.target);
        bool optimalOk = validate(optimalResult, tc.nums, tc.numsSize, tc.target);

        if (bruteOk) brutePassed++;
        if (optimalOk) optimalPassed++;

        printf("\n%d. %s\n", i + 1, tc.desc);
        printf("   Input: nums=");
        printArray(tc.nums, tc.numsSize);
        printf(", target=%d\n", tc.target);
        printf("   Brute:   [%s] ", bruteOk ? "PASS" : "FAIL");
        printResult(bruteResult);
        printf("\n");
        printf("   Optimal: [%s] ", optimalOk ? "PASS" : "FAIL");
        printResult(optimalResult);
        printf("\n");
    }

    printf("\n======================================================================\n");
    printf("Summary: Brute %d/%d | Optimal %d/%d\n", brutePassed, total, optimalPassed, total);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. What is the time complexity of brute force? Why?\n");
    printf("2. How would you implement a hash table in C?\n");
    printf("3. Can this be solved with two pointers? What's the prerequisite?\n");

    return 0;
}
