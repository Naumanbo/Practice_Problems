// Tests: Dynamic programming, Kadane's algorithm
//
// Maximum Subarray (LeetCode 53)
// Find the contiguous subarray with the largest sum.

#include <stdio.h>
#include <limits.h>

int maxSubArray(int* nums, int numsSize) {
    // TODO: Implement Kadane's algorithm
    // Hint: At each position, decide: extend current subarray or start new?
    return 0;
}

int main() {
    struct TestCase {
        int nums[10];
        int size;
        int expected;
        char* desc;
    };

    struct TestCase tests[] = {
        {{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 9, 6, "mixed values"},
        {{1}, 1, 1, "single element"},
        {{5, 4, -1, 7, 8}, 5, 23, "mostly positive"},
        {{-1}, 1, -1, "single negative"},
        {{-2, -1}, 2, -1, "all negative"},
        {{-2, -3, -1, -5}, 4, -1, "all negative longer"},
        {{1, 2, 3, 4}, 4, 10, "all positive"},
        {{-1, 0, -2}, 3, 0, "zero is max"},
        {{8, -19, 5, -4, 20}, 5, 21, "recovery after negative"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);
    int allPassed = 1;

    for (int i = 0; i < numTests; i++) {
        int result = maxSubArray(tests[i].nums, tests[i].size);
        if (result != tests[i].expected) {
            printf("FAIL [%s]: got %d, expected %d\n",
                   tests[i].desc, result, tests[i].expected);
            allPassed = 0;
        } else {
            printf("PASS [%s]\n", tests[i].desc);
        }
    }

    if (allPassed) {
        printf("\nAll tests passed!\n");
    }

    return 0;
}
