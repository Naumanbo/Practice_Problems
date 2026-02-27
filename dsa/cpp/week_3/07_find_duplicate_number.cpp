// Tests: Floyd's cycle detection on arrays, two pointers, O(1) space
//
// Find the Duplicate Number (LeetCode #287)
// Array of n+1 integers where each is in [1,n]. Find the one duplicate.
// Must not modify the array and use only O(1) extra space.

#include <iostream>
#include <vector>
using namespace std;

int findDuplicate(vector<int>& nums) {
    return 0;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    vector<int> nums;
    int expected;
    string desc;
};

int main() {
    vector<TestCase> tests = {
        {{1, 3, 4, 2, 2},                    2, "basic case"},
        {{3, 1, 3, 4, 2},                    3, "duplicate not adjacent"},
        {{1, 1},                             1, "two elements"},
        {{1, 1, 2},                          1, "three elements"},
        {{1, 2, 3, 4, 4},                   4, "duplicate at end"},
        {{1, 2, 3, 1},                      1, "duplicate wraps"},
        {{2, 1, 2, 3},                      2, "duplicate early"},
        {{6, 2, 4, 1, 3, 5, 6},            6, "six elements"},
        {{9, 7, 4, 6, 3, 2, 8, 5, 1, 1},  1, "duplicate is 1"},
        {{2, 5, 9, 6, 3, 8, 7, 1, 4, 9},  9, "duplicate is 9"},
        {{3, 4, 8, 5, 9, 1, 6, 8, 7, 2},  8, "duplicate in middle"},
        {{5, 1, 2, 3, 4, 5},               5, "duplicate at boundaries"},
        {{1, 2, 3, 2, 4},                  2, "five elements"},
        {{4, 3, 2, 1, 4},                  4, "duplicate at front and back"},
        {{1, 2, 1, 3},                     1, "four elements"},
        {{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 9}, 9, "reverse sorted with dup"},
    };

    cout << "======================================================================" << endl;
    cout << "FIND DUPLICATE NUMBER - Test Results" << endl;
    cout << "======================================================================" << endl;

    int passed = 0;
    for (size_t i = 0; i < tests.size(); i++) {
        int result = findDuplicate(tests[i].nums);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        cout << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tests[i].desc << endl;
    }

    cout << "======================================================================" << endl;
    cout << "Summary: " << passed << "/" << tests.size() << " passed" << endl;
    cout << "======================================================================" << endl;
    return 0;
}
