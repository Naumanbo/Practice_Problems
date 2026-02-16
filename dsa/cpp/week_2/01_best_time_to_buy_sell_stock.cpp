// Tests: Arrays, single pass, tracking minimum
//
// Best Time to Buy and Sell Stock (LeetCode 121)
// Given prices where prices[i] is stock price on day i,
// return maximum profit from one buy-sell transaction.

#include <iostream>
#include <vector>
#include <climits>
using namespace std;

int maxProfit(vector<int>& prices) {
    // TODO: Implement solution
    // Hint: Track minimum price seen so far and maximum profit
    return 0;
}

int main() {
    struct TestCase {
        vector<int> prices;
        int expected;
        string desc;
    };

    vector<TestCase> tests = {
        {{7, 1, 5, 3, 6, 4}, 5, "basic case"},
        {{7, 6, 4, 3, 1}, 0, "decreasing prices"},
        {{1, 2, 3, 4, 5}, 4, "increasing prices"},
        {{2, 4, 1}, 2, "buy early sell middle"},
        {{3, 3, 3}, 0, "flat prices"},
        {{1}, 0, "single element"},
        {{2, 1}, 0, "two elements decreasing"},
        {{1, 2}, 1, "two elements increasing"},
        {{10, 1, 10}, 9, "V-shape recovery"},
        {{3, 8, 2, 5}, 5, "peak then valley"},
        {{1, 5, 2, 8}, 7, "multiple peaks best at end"},
        {{5, 4, 3, 2, 1, 10}, 9, "min at end then spike"},
        {{0, 0, 0}, 0, "all zeros"},
        {{5, 5}, 0, "two same"},
        {{3, 1, 100, 2, 4}, 99, "spike in middle"},
        {{1, 10, 1, 10, 1}, 9, "oscillating"},
        {{0, 1}, 1, "zero then one"},
        {{10000, 1, 10000}, 9999, "near constraint max"},
        {{2, 1, 2, 0, 1}, 1, "multiple dips"},
        {{1, 1, 1, 1, 2}, 1, "flat then rise"},
    };

    bool allPassed = true;
    for (auto& tc : tests) {
        int result = maxProfit(tc.prices);
        if (result != tc.expected) {
            cout << "FAIL [" << tc.desc << "]: got " << result
                 << ", expected " << tc.expected << endl;
            allPassed = false;
        } else {
            cout << "PASS [" << tc.desc << "]" << endl;
        }
    }

    if (allPassed) {
        cout << "\nAll tests passed!" << endl;
    }

    return 0;
}
