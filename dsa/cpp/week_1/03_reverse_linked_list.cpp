// Key Takeaways:
// 1. In C++, always check `head == nullptr` before `head->next == nullptr`.
//    The || operator short-circuits left to right — dereferencing a null pointer
//    before checking for null causes a segfault.
//
// 2. Iterative reversal uses three pointers (prev, curr, next), all declared as
//    ListNode*. C++ is case-sensitive — ListNode* != listNode* != lidtNode*.
//
// 3. Recursive template: base case first, recurse on smaller input, then combine.
//    For reversal: head->next->next = head flips the back-pointer,
//    head->next = nullptr cuts the forward link to prevent cycles.
//
// 4. Iterative is O(1) space. Recursive is O(n) space due to call stack.
//    Prefer iterative in production for large lists.
//
// Complexity: Time O(n), Space O(1) iterative / O(n) recursive

/*
DSA Problem 3: Reverse Linked List

Tests: Linked list traversal, pointer manipulation, iterative vs recursive

Difficulty: Easy
Source: LeetCode #206

Problem:
Given the head of a singly linked list, reverse the list and return it.

Compile: g++ -std=c++17 -o 03_reverse_linked_list 03_reverse_linked_list.cpp
Run: ./03_reverse_linked_list
*/

#include <iostream>
#include <vector>
#include <string>

using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

// Iterative approach using three pointers
// Time: O(?)
// Space: O(?)
ListNode* reverseIterative(ListNode* head) {
    // Your implementation
    ListNode* prev = nullptr;
    ListNode* curr = head;
    ListNode* next = nullptr;

    while (curr != nullptr) {
        next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }

    return prev;
}

// Recursive approach
// Time: O(?)
// Space: O(?) - consider call stack
ListNode* reverseRecursive(ListNode* head) {
    // Your implementation
    if (head == nullptr || head->next == nullptr) {
        return head;
    }

    ListNode* new_head = reverseRecursive(head->next);
    head->next->next = head;
    head->next = nullptr;
    return new_head;
}

// =============================================================================
// Helper functions
// =============================================================================

ListNode* vectorToList(const vector<int>& arr) {
    if (arr.empty()) return nullptr;
    ListNode* head = new ListNode(arr[0]);
    ListNode* current = head;
    for (size_t i = 1; i < arr.size(); i++) {
        current->next = new ListNode(arr[i]);
        current = current->next;
    }
    return head;
}

vector<int> listToVector(ListNode* head) {
    vector<int> result;
    while (head) {
        result.push_back(head->val);
        head = head->next;
    }
    return result;
}

void freeList(ListNode* head) {
    while (head) {
        ListNode* temp = head;
        head = head->next;
        delete temp;
    }
}

void printVector(const vector<int>& v, int maxLen = 8) {
    cout << "[";
    size_t limit = min(v.size(), (size_t)maxLen);
    for (size_t i = 0; i < limit; i++) {
        cout << v[i];
        if (i < limit - 1) cout << ", ";
    }
    if (v.size() > (size_t)maxLen) cout << "...";
    cout << "]";
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    vector<int> input;
    vector<int> expected;
    string desc;
};

int runTests(ListNode* (*fn)(ListNode*), const string& name, const vector<TestCase>& testCases) {
    cout << "\n" << name << ":" << endl;
    int passed = 0;

    for (size_t i = 0; i < testCases.size(); i++) {
        const auto& tc = testCases[i];
        ListNode* head = vectorToList(tc.input);
        ListNode* reversed = fn(head);
        vector<int> result = listToVector(reversed);
        freeList(reversed);

        bool ok = (result == tc.expected);
        if (ok) passed++;

        string status = ok ? "PASS" : "FAIL";
        cout << "  " << (i + 1) << ". [" << status << "] " << tc.desc << ": ";
        printVector(tc.input);
        cout << " -> ";
        printVector(result);
        cout << endl;
    }

    return passed;
}

int main() {
    vector<TestCase> testCases = {
        // Base cases
        {{}, {}, "Empty list"},
        {{1}, {1}, "Single element"},
        {{1, 2}, {2, 1}, "Two elements"},

        // Basic cases
        {{1, 2, 3}, {3, 2, 1}, "Three elements"},
        {{1, 2, 3, 4, 5}, {5, 4, 3, 2, 1}, "Five elements"},

        // Edge cases - duplicates
        {{1, 1, 1}, {1, 1, 1}, "All same values"},
        {{1, 2, 2, 1}, {1, 2, 2, 1}, "Palindrome"},
        {{1, 1, 2, 2}, {2, 2, 1, 1}, "Pairs"},

        // Edge cases - negative numbers
        {{-1, -2, -3}, {-3, -2, -1}, "Negative values"},
        {{-1, 0, 1}, {1, 0, -1}, "Mixed signs"},

        // Edge cases - large values
        {{1000000, 2000000}, {2000000, 1000000}, "Large values"},

        // Longer lists
        {{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, "0-9 sequence"},
    };

    cout << "======================================================================" << endl;
    cout << "REVERSE LINKED LIST - Test Results" << endl;
    cout << "======================================================================" << endl;

    int total = testCases.size();
    int iterPassed = runTests(reverseIterative, "Iterative", testCases);
    int recPassed = runTests(reverseRecursive, "Recursive", testCases);

    cout << "\n======================================================================" << endl;
    cout << "Summary: Iterative " << iterPassed << "/" << total
         << " | Recursive " << recPassed << "/" << total << endl;
    cout << "======================================================================" << endl;
    cout << "\nQuestions:" << endl;
    cout << "1. Draw pointer changes for [1,2,3] step by step (iterative)." << endl;
    cout << "2. What's the space complexity difference between approaches?" << endl;
    cout << "3. Why prefer iterative in production code?" << endl;

    return 0;
}
