/*
DSA Problem 1: Two Sum

Tests: Hash map usage, time/space complexity analysis, array traversal

Difficulty: Easy
Source: LeetCode #1

Problem:
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

Constraints:
  - 2 <= nums.size() <= 10^4
  - -10^9 <= nums[i] <= 10^9
  - Only one valid answer exists.
  - You may not use the same element twice.

Compile: g++ -std=c++17 -o 01_two_sum 01_two_sum.cpp
Run: ./01_two_sum
*/

#include <iostream>
#include <vector>
#include <unordered_map>
#include <string>

using namespace std;

// Brute force approach
// Time: O(?)
// Space: O(?)
vector<int> twoSumBrute(vector<int>& nums, int target) {
    for (int i = 0; i < nums.size(); i++) {
        for (int j = i + 1; j < nums.size(); j++ ) {
            if (target == nums[i] + nums[j]) {
                return vector<int>{i,j};
            }
        }
    }
    return {};
}

// Optimal approach using hash map
// Time: O(?)
// Space: O(?)
vector<int> twoSumOptimal(vector<int>& nums, int target) {
    // Your implementation
    unordered_map<int, int> seen;
    for (int i = 0; i < nums.size(); i++) {
        int requiredValue = target - nums[i];
        if (seen.find(requiredValue) != seen.end()) {
            return vector<int>{i, seen[requiredValue]};
        }
        seen[nums[i]] = i;
    }
    return {};
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    vector<int> nums;
    int target;
    string desc;
};

bool validate(const vector<int>& result, const vector<int>& nums, int target) {
    if (result.size() != 2) return false;
    if (result[0] == result[1]) return false;
    if (result[0] < 0 || result[0] >= (int)nums.size()) return false;
    if (result[1] < 0 || result[1] >= (int)nums.size()) return false;
    return nums[result[0]] + nums[result[1]] == target;
}

void printVector(const vector<int>& v) {
    cout << "[";
    for (size_t i = 0; i < v.size(); i++) {
        cout << v[i];
        if (i < v.size() - 1) cout << ", ";
    }
    cout << "]";
}

int main() {
    vector<TestCase> testCases = {
        // Basic cases
        {{2, 7, 11, 15}, 9, "Basic - first two elements"},
        {{3, 2, 4}, 6, "Basic - middle elements"},
        {{3, 3}, 6, "Duplicate values"},

        // Edge cases - array size
        {{1, 2}, 3, "Minimum size (2 elements)"},
        {{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 19, "Larger array"},

        // Edge cases - negative numbers
        {{-1, -2, -3, -4}, -6, "All negative"},
        {{-3, 4, 3, 90}, 0, "Negative + positive = 0"},
        {{5, -5, 10}, 0, "Sum to zero"},

        // Edge cases - zeros
        {{0, 4, 3, 0}, 0, "Two zeros"},
        {{0, 1, 2}, 2, "Zero + positive"},

        // Edge cases - large numbers
        {{1000000000, 2, 1000000000}, 2000000000, "Large numbers"},

        // Edge cases - position variations
        {{1, 5, 8, 3}, 4, "Answer at start and end"},
        {{4, 5, 1, 2}, 6, "Answer in middle"},
        {{5, 5, 5, 5}, 10, "All same values"},
    };

    cout << "======================================================================" << endl;
    cout << "TWO SUM - Test Results" << endl;
    cout << "======================================================================" << endl;

    int brutePassed = 0, optimalPassed = 0;
    int total = testCases.size();

    for (size_t i = 0; i < testCases.size(); i++) {
        auto& tc = testCases[i];
        vector<int> bruteCopy = tc.nums;
        vector<int> optimalCopy = tc.nums;

        vector<int> bruteResult = twoSumBrute(bruteCopy, tc.target);
        vector<int> optimalResult = twoSumOptimal(optimalCopy, tc.target);

        bool bruteOk = validate(bruteResult, tc.nums, tc.target);
        bool optimalOk = validate(optimalResult, tc.nums, tc.target);

        if (bruteOk) brutePassed++;
        if (optimalOk) optimalPassed++;

        string bruteStatus = bruteOk ? "PASS" : "FAIL";
        string optimalStatus = optimalOk ? "PASS" : "FAIL";

        cout << "\n" << (i + 1) << ". " << tc.desc << endl;
        cout << "   Input: nums=";
        printVector(tc.nums);
        cout << ", target=" << tc.target << endl;
        cout << "   Brute:   [" << bruteStatus << "] ";
        printVector(bruteResult);
        cout << endl;
        cout << "   Optimal: [" << optimalStatus << "] ";
        printVector(optimalResult);
        cout << endl;
    }

    cout << "\n======================================================================" << endl;
    cout << "Summary: Brute " << brutePassed << "/" << total
         << " | Optimal " << optimalPassed << "/" << total << endl;
    cout << "======================================================================" << endl;
    cout << "\nQuestions:" << endl;
    cout << "1. What is the time complexity of brute force? Why?" << endl;
    cout << "2. What is the time complexity of the hash map approach? Why?" << endl;
    cout << "3. Can this be solved with two pointers? What's the prerequisite?" << endl;

    return 0;
}
