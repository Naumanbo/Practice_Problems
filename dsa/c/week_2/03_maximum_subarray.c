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
        {{1, -1, 1, -1, 1}, 5, 1, "alternating"},
        {{5, 6, -100, 7, 8}, 5, 15, "large negative breaks"},
        {{10, -1, -1, -1, -1}, 5, 10, "subarray at start"},
        {{-1, -1, -1, -1, 10}, 5, 10, "subarray at end"},
        {{5, -10, 5}, 3, 5, "two equal subarrays"},
        {{100}, 1, 100, "single large positive"},
        {{10, -5, 10, -5, 10}, 5, 20, "worth keeping negatives"},
        {{1, -1, 1, -1, 2}, 5, 2, "max at end"},
        {{3, -1, 2, -1, 4}, 5, 7, "entire array is max"},
        {{0}, 1, 0, "single zero"},
        {{1, 2, -1, 3, -2, 4}, 6, 7, "scattered negatives"},
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
