// Key Takeaways:
// 1. Kadane's algorithm: at each element, decide whether to extend the current
//    subarray or start fresh. `currSum = max(nums[i], currSum + nums[i])` —
//    if currSum is negative, starting fresh is always better.
//
// 2. Initialize both currSum and maxSum to nums[0], not 0. If all values are
//    negative, returning 0 would be wrong — the answer is the least negative element.
//
// 3. Use std::max from <algorithm> for clean comparisons. Loop starts at index 1
//    since index 0 is used for initialization.
//
// 4. This is O(n) time, O(1) space — a single pass with no extra storage.
//    The naive approach (all subarrays) is O(n²) or O(n³).
//
// Complexity: Time O(n), Space O(1)

// Tests: Dynamic programming, Kadane's algorithm
//
// Maximum Subarray (LeetCode 53)
// Find the contiguous subarray with the largest sum.

#include <iostream>
#include <vector>
#include <climits>
#include <algorithm>
using namespace std;

int maxSubArray(vector<int>& nums) {
    // TODO: Implement Kadane's algorithm
    // Hint: At each position, decide: extend current subarray or start new?
    int currSum = nums[0];
    int maxSum = nums[0];

    for (int i = 1; i < nums.size(); i++) {
        currSum = max(nums[i], currSum + nums[i]);
        maxSum = max(currSum, maxSum);
    }
    return maxSum;
}

int main() {
    struct TestCase {
        vector<int> nums;
        int expected;
        string desc;
    };

    vector<TestCase> tests = {
        {{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, "mixed values"},
        {{1}, 1, "single element"},
        {{5, 4, -1, 7, 8}, 23, "mostly positive"},
        {{-1}, -1, "single negative"},
        {{-2, -1}, -1, "all negative"},
        {{-2, -3, -1, -5}, -1, "all negative longer"},
        {{1, 2, 3, 4}, 10, "all positive"},
        {{-1, 0, -2}, 0, "zero is max"},
        {{8, -19, 5, -4, 20}, 21, "recovery after negative"},
        {{1, -1, 1, -1, 1}, 1, "alternating"},
        {{5, 6, -100, 7, 8}, 15, "large negative breaks subarray"},
        {{10, -1, -1, -1, -1}, 10, "subarray at start"},
        {{-1, -1, -1, -1, 10}, 10, "subarray at end"},
        {{5, -10, 5}, 5, "two equal subarrays"},
        {{100}, 100, "single large positive"},
        {{10, -5, 10, -5, 10}, 20, "worth keeping negatives"},
        {{1, -1, 1, -1, 2}, 2, "max at end"},
        {{3, -1, 2, -1, 4}, 7, "entire array is max"},
        {{0}, 0, "single zero"},
        {{-10000}, -10000, "single large negative"},
        {{1, 2, -1, 3, -2, 4}, 7, "scattered negatives"},
    };

    bool allPassed = true;
    for (auto& tc : tests) {
        int result = maxSubArray(tc.nums);
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
