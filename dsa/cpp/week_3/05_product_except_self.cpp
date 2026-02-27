// Tests: Arrays, prefix/suffix products, no-division constraint
//
// Product of Array Except Self (LeetCode #238)
// Return array where output[i] = product of all nums except nums[i].
// Must run in O(n) time without division.

#include <iostream>
#include <vector>
using namespace std;

vector<int> productExceptSelf(vector<int>& nums) {
    return {};
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    vector<int> nums;
    vector<int> expected;
    string desc;
};

int main() {
    vector<TestCase> tests = {
        {{1, 2, 3, 4},       {24, 12, 8, 6},        "basic case"},
        {{2, 3, 4, 5},       {60, 40, 30, 24},       "all positive"},
        {{-1, 1, 0, -3, 3},  {0, 0, 9, 0, 0},        "contains zero"},
        {{1, 2},             {2, 1},                  "two elements"},
        {{3, 3, 3},          {9, 9, 9},               "all same"},
        {{1, 1, 1, 1},       {1, 1, 1, 1},            "all ones"},
        {{0, 0},             {0, 0},                  "all zeros"},
        {{1, 0},             {0, 1},                  "one zero"},
        {{-1, 2, 3},         {6, -3, -2},             "negative value"},
        {{2, 2, 2, 2},       {8, 8, 8, 8},            "all twos"},
        {{1, 2, 3, 4, 5},    {120, 60, 40, 30, 24},   "five elements"},
        {{-1, -2, -3, -4},   {-24, -12, -8, -6},      "all negative"},
        {{100, 1, 2},        {2, 200, 100},           "large first element"},
        {{0, 1, 2, 3},       {6, 0, 0, 0},            "zero at start"},
        {{1, 2, 0, 4},       {0, 0, 8, 0},            "zero in middle"},
        {{2, 3},             {3, 2},                  "two elements v2"},
    };

    cout << "======================================================================" << endl;
    cout << "PRODUCT EXCEPT SELF - Test Results" << endl;
    cout << "======================================================================" << endl;

    int passed = 0;
    for (size_t i = 0; i < tests.size(); i++) {
        vector<int> result = productExceptSelf(tests[i].nums);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        cout << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tests[i].desc << endl;
    }

    cout << "======================================================================" << endl;
    cout << "Summary: " << passed << "/" << tests.size() << " passed" << endl;
    cout << "======================================================================" << endl;
    return 0;
}
