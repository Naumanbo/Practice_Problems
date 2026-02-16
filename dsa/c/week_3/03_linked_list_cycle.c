// Compile: gcc -o 03_linked_list_cycle 03_linked_list_cycle.c
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>

typedef struct ListNode {
    int val;
    struct ListNode* next;
} ListNode;

bool hasCycle(ListNode* head) {
    // Your implementation (Floyd's tortoise and hare)
    return false;
}

ListNode* createNode(int val) {
    ListNode* node = (ListNode*)malloc(sizeof(ListNode));
    node->val = val;
    node->next = NULL;
    return node;
}

ListNode* makeListWithCycle(int* values, int size, int cyclePos) {
    if (size == 0) return NULL;
    ListNode** nodes = (ListNode**)malloc(size * sizeof(ListNode*));
    for (int i = 0; i < size; i++) nodes[i] = createNode(values[i]);
    for (int i = 0; i < size - 1; i++) nodes[i]->next = nodes[i + 1];
    if (cyclePos >= 0 && cyclePos < size) nodes[size - 1]->next = nodes[cyclePos];
    ListNode* head = nodes[0];
    free(nodes);
    return head;
}

struct TestCase {
    int values[10];
    int size;
    int cyclePos;
    bool expected;
    char* desc;
};

int main() {
    struct TestCase tests[] = {
        {{3, 2, 0, -4}, 4, 1, true, "cycle at pos 1"},
        {{1, 2}, 2, 0, true, "cycle at head"},
        {{1, 2, 3, 4}, 4, 0, true, "cycle back to head"},
        {{1, 2, 3, 4}, 4, 2, true, "cycle in middle"},
        {{1}, 1, -1, false, "single node no cycle"},
        {{1, 2}, 2, -1, false, "two nodes no cycle"},
        {{1, 2, 3, 4, 5}, 5, -1, false, "five nodes no cycle"},
        {{0}, 0, -1, false, "empty list"},
        {{1}, 1, 0, true, "self-loop"},
        {{1, 2, 3}, 3, -1, false, "three nodes no cycle"},
        {{1, 2, 3, 4, 5}, 5, 4, true, "tail self-loop"},
        {{1, 2, 3, 4, 5}, 5, 3, true, "cycle at second to last"},
        {{1, 2, 3, 4, 5}, 5, 1, true, "cycle at pos 1"},
        {{-1, -2, -3}, 3, -1, false, "negative values no cycle"},
        {{-1, -2, -3}, 3, 0, true, "negative values with cycle"},
        {{1, 2}, 2, 1, true, "two nodes tail self-loop"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);
    int passed = 0;

    printf("======================================================================\n");
    printf("LINKED LIST CYCLE - Test Results\n");
    printf("======================================================================\n\n");

    for (int i = 0; i < numTests; i++) {
        ListNode* head = makeListWithCycle(tests[i].values, tests[i].size, tests[i].cyclePos);
        bool result = hasCycle(head);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        printf("%2d. [%s] %s: pos=%d -> %s\n",
               i + 1, ok ? "PASS" : "FAIL", tests[i].desc,
               tests[i].cyclePos, result ? "true" : "false");
    }

    printf("\n======================================================================\n");
    printf("Summary: %d/%d passed\n", passed, numTests);
    printf("======================================================================\n");
    printf("\nQuestions:\n");
    printf("1. Why does Floyd's guarantee fast meets slow?\n");
    printf("2. Can you implement hash set in C? How?\n");
    printf("3. Find WHERE cycle begins (LeetCode #142)\n");
    return 0;
}
