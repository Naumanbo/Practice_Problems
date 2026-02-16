// Compile: gcc -o 01_binary_search 01_binary_search.c
#include <stdio.h>

int binarySearchIterative(int* nums, int numsSize, int target) {
    // Your implementation
    return -1;
}

int bsHelper(int* nums, int left, int right, int target) {
    // Your implementation
    return -1;
}

int binarySearchRecursive(int* nums, int numsSize, int target) {
    return bsHelper(nums, 0, numsSize - 1, target);
}

struct TestCase {
    int nums[110];
    int numsSize;
    int target;
    int expected;
    char* desc;
};

int main() {
    struct TestCase tests[] = {
        {{-1, 0, 3, 5, 9, 12}, 6, 9, 4, "target in middle"},
        {{-1, 0, 3, 5, 9, 12}, 6, 2, -1, "target not found"},
        {{1, 2, 3, 4, 5}, 5, 1, 0, "target at start"},
        {{1, 2, 3, 4, 5}, 5, 5, 4, "target at end"},
        {{1, 2, 3, 4, 5}, 5, 3, 2, "target in center"},
        {{5}, 1, 5, 0, "single element found"},
        {{5}, 1, 3, -1, "single element not found"},
        {{1, 2}, 2, 1, 0, "two elements first"},
        {{1, 2}, 2, 2, 1, "two elements second"},
        {{1, 2}, 2, 3, -1, "two elements not found"},
        {{-10, -5, 0, 5, 10}, 5, -5, 1, "negative target"},
        {{-10, -5, 0, 5, 10}, 5, 0, 2, "zero target"},
        {{-100, -50, -10, -1}, 4, -100, 0, "all negative first"},
        {{-100, -50, -10, -1}, 4, -1, 3, "all negative last"},
        {{-100, -50, -10, -1}, 4, 5, -1, "all negative not found"},
        {{1, 2, 3, 4, 5}, 5, 0, -1, "target below range"},
        {{1, 2, 3, 4, 5}, 5, 6, -1, "target above range"},
        {{-9999, 0, 9999}, 3, 9999, 2, "near constraint bounds"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);

    printf("======================================================================\n");
    printf("BINARY SEARCH - Test Results\n");
    printf("======================================================================\n");

    printf("\nIterative:\n");
    int ip = 0;
    for (int i = 0; i < numTests; i++) {
        int r = binarySearchIterative(tests[i].nums, tests[i].numsSize, tests[i].target);
        int ok = (r == tests[i].expected);
        if (ok) ip++;
        printf("  %2d. [%s] %s: target=%d -> %d (expected %d)\n",
               i + 1, ok ? "PASS" : "FAIL", tests[i].desc, tests[i].target, r, tests[i].expected);
    }

    printf("\nRecursive:\n");
    int rp = 0;
    for (int i = 0; i < numTests; i++) {
        int r = binarySearchRecursive(tests[i].nums, tests[i].numsSize, tests[i].target);
        int ok = (r == tests[i].expected);
        if (ok) rp++;
        printf("  %2d. [%s] %s: target=%d -> %d (expected %d)\n",
               i + 1, ok ? "PASS" : "FAIL", tests[i].desc, tests[i].target, r, tests[i].expected);
    }

    printf("\n======================================================================\n");
    printf("Summary: Iterative %d/%d | Recursive %d/%d\n", ip, numTests, rp, numTests);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. Why must the array be sorted?\n");
    printf("2. Why is left+(right-left)/2 safer than (left+right)/2?\n");
    printf("3. How does bsearch() from stdlib.h work?\n");
    return 0;
}
