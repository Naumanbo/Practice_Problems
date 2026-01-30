/*
DSA Problem 3: Reverse Linked List

Tests: Linked list traversal, pointer manipulation, iterative vs recursive

Difficulty: Easy
Source: LeetCode #206

Problem:
Given the head of a singly linked list, reverse the list and return it.

Compile: gcc -o 03_reverse_linked_list 03_reverse_linked_list.c
Run: ./03_reverse_linked_list
*/

#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

typedef struct ListNode {
    int val;
    struct ListNode* next;
} ListNode;

// Iterative approach using three pointers
// Time: O(?)
// Space: O(?)
ListNode* reverseIterative(ListNode* head) {
    // Your implementation
    return NULL;
}

// Recursive approach
// Time: O(?)
// Space: O(?) - consider call stack
ListNode* reverseRecursive(ListNode* head) {
    // Your implementation
    return NULL;
}

// =============================================================================
// Helper functions
// =============================================================================

ListNode* createNode(int val) {
    ListNode* node = malloc(sizeof(ListNode));
    node->val = val;
    node->next = NULL;
    return node;
}

ListNode* arrayToList(int* arr, int size) {
    if (size == 0) return NULL;
    ListNode* head = createNode(arr[0]);
    ListNode* current = head;
    for (int i = 1; i < size; i++) {
        current->next = createNode(arr[i]);
        current = current->next;
    }
    return head;
}

int* listToArray(ListNode* head, int* outSize) {
    // Count nodes
    int count = 0;
    ListNode* temp = head;
    while (temp) {
        count++;
        temp = temp->next;
    }

    *outSize = count;
    if (count == 0) return NULL;

    int* arr = malloc(count * sizeof(int));
    temp = head;
    for (int i = 0; i < count; i++) {
        arr[i] = temp->val;
        temp = temp->next;
    }
    return arr;
}

void freeList(ListNode* head) {
    while (head) {
        ListNode* temp = head;
        head = head->next;
        free(temp);
    }
}

bool arraysEqual(int* a, int aSize, int* b, int bSize) {
    if (aSize != bSize) return false;
    for (int i = 0; i < aSize; i++) {
        if (a[i] != b[i]) return false;
    }
    return true;
}

void printArray(int* arr, int size, int maxLen) {
    printf("[");
    int limit = size < maxLen ? size : maxLen;
    for (int i = 0; i < limit; i++) {
        printf("%d", arr[i]);
        if (i < limit - 1) printf(", ");
    }
    if (size > maxLen) printf("...");
    printf("]");
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

typedef struct {
    int* input;
    int inputSize;
    int* expected;
    int expectedSize;
    const char* desc;
} TestCase;

int runTests(ListNode* (*fn)(ListNode*), const char* name, TestCase* testCases, int total) {
    printf("\n%s:\n", name);
    int passed = 0;

    for (int i = 0; i < total; i++) {
        TestCase tc = testCases[i];
        ListNode* head = arrayToList(tc.input, tc.inputSize);
        ListNode* reversed = fn(head);

        int resultSize;
        int* result = listToArray(reversed, &resultSize);

        bool ok = arraysEqual(result, resultSize, tc.expected, tc.expectedSize);
        if (ok) passed++;

        const char* status = ok ? "PASS" : "FAIL";
        printf("  %2d. [%s] %s: ", i + 1, status, tc.desc);
        printArray(tc.input, tc.inputSize, 8);
        printf(" -> ");
        printArray(result, resultSize, 8);
        printf("\n");

        freeList(reversed);
        free(result);
    }

    return passed;
}

int main() {
    // Test data
    int empty[] = {};
    int single[] = {1};
    int two[] = {1, 2};
    int three[] = {1, 2, 3};
    int five[] = {1, 2, 3, 4, 5};
    int same[] = {1, 1, 1};
    int palin[] = {1, 2, 2, 1};
    int pairs[] = {1, 1, 2, 2};
    int neg[] = {-1, -2, -3};
    int mixed[] = {-1, 0, 1};
    int large[] = {1000000, 2000000};
    int seq[] = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9};

    int exp_empty[] = {};
    int exp_single[] = {1};
    int exp_two[] = {2, 1};
    int exp_three[] = {3, 2, 1};
    int exp_five[] = {5, 4, 3, 2, 1};
    int exp_same[] = {1, 1, 1};
    int exp_palin[] = {1, 2, 2, 1};
    int exp_pairs[] = {2, 2, 1, 1};
    int exp_neg[] = {-3, -2, -1};
    int exp_mixed[] = {1, 0, -1};
    int exp_large[] = {2000000, 1000000};
    int exp_seq[] = {9, 8, 7, 6, 5, 4, 3, 2, 1, 0};

    TestCase testCases[] = {
        {empty, 0, exp_empty, 0, "Empty list"},
        {single, 1, exp_single, 1, "Single element"},
        {two, 2, exp_two, 2, "Two elements"},
        {three, 3, exp_three, 3, "Three elements"},
        {five, 5, exp_five, 5, "Five elements"},
        {same, 3, exp_same, 3, "All same values"},
        {palin, 4, exp_palin, 4, "Palindrome"},
        {pairs, 4, exp_pairs, 4, "Pairs"},
        {neg, 3, exp_neg, 3, "Negative values"},
        {mixed, 3, exp_mixed, 3, "Mixed signs"},
        {large, 2, exp_large, 2, "Large values"},
        {seq, 10, exp_seq, 10, "0-9 sequence"},
    };

    int total = sizeof(testCases) / sizeof(TestCase);

    printf("======================================================================\n");
    printf("REVERSE LINKED LIST - Test Results\n");
    printf("======================================================================\n");

    int iterPassed = runTests(reverseIterative, "Iterative", testCases, total);
    int recPassed = runTests(reverseRecursive, "Recursive", testCases, total);

    printf("\n======================================================================\n");
    printf("Summary: Iterative %d/%d | Recursive %d/%d\n", iterPassed, total, recPassed, total);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. Draw pointer changes for [1,2,3] step by step (iterative).\n");
    printf("2. What's the space complexity difference between approaches?\n");
    printf("3. Why prefer iterative in production code?\n");

    return 0;
}
