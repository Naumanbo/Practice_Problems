// Tests: Arrays, single pass, tracking minimum
//
// Best Time to Buy and Sell Stock (LeetCode 121)
// Given prices where prices[i] is stock price on day i,
// return maximum profit from one buy-sell transaction.

#include <stdio.h>
#include <limits.h>

int maxProfit(int* prices, int pricesSize) {
    // TODO: Implement solution
    // Hint: Track minimum price seen so far and maximum profit
    return 0;
}

int main() {
    struct TestCase {
        int prices[10];
        int size;
        int expected;
        char* desc;
    };

    struct TestCase tests[] = {
        {{7, 1, 5, 3, 6, 4}, 6, 5, "basic case"},
        {{7, 6, 4, 3, 1}, 5, 0, "decreasing prices"},
        {{1, 2, 3, 4, 5}, 5, 4, "increasing prices"},
        {{2, 4, 1}, 3, 2, "buy early sell middle"},
        {{3, 3, 3}, 3, 0, "flat prices"},
        {{1}, 1, 0, "single element"},
        {{2, 1}, 2, 0, "two elements decreasing"},
        {{1, 2}, 2, 1, "two elements increasing"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);
    int allPassed = 1;

    for (int i = 0; i < numTests; i++) {
        int result = maxProfit(tests[i].prices, tests[i].size);
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
