// Tests: Linked lists, two pointers, iteration
//
// Merge Two Sorted Lists (LeetCode 21)
// Merge two sorted linked lists into one sorted list.

#include <stdio.h>
#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

struct ListNode* mergeTwoLists(struct ListNode* list1, struct ListNode* list2) {
    // TODO: Implement solution
    // Hint: Use a dummy head node to simplify edge cases
    return NULL;
}

// Helper: create new node
struct ListNode* createNode(int val) {
    struct ListNode* node = (struct ListNode*)malloc(sizeof(struct ListNode));
    node->val = val;
    node->next = NULL;
    return node;
}

// Helper: convert array to linked list
struct ListNode* arrayToList(int* arr, int size) {
    if (size == 0) return NULL;
    struct ListNode* head = createNode(arr[0]);
    struct ListNode* current = head;
    for (int i = 1; i < size; i++) {
        current->next = createNode(arr[i]);
        current = current->next;
    }
    return head;
}

// Helper: check if list matches expected array
int listMatchesArray(struct ListNode* head, int* expected, int size) {
    int i = 0;
    while (head && i < size) {
        if (head->val != expected[i]) return 0;
        head = head->next;
        i++;
    }
    return (head == NULL && i == size);
}

// Helper: print list
void printList(struct ListNode* head) {
    printf("[");
    while (head) {
        printf("%d", head->val);
        if (head->next) printf(", ");
        head = head->next;
    }
    printf("]");
}

// Helper: free list
void freeList(struct ListNode* head) {
    while (head) {
        struct ListNode* temp = head;
        head = head->next;
        free(temp);
    }
}

int main() {
    struct TestCase {
        int list1[6];
        int size1;
        int list2[6];
        int size2;
        int expected[12];
        int expectedSize;
        char* desc;
    };

    struct TestCase tests[] = {
        {{1, 2, 4}, 3, {1, 3, 4}, 3, {1, 1, 2, 3, 4, 4}, 6, "basic merge"},
        {{}, 0, {}, 0, {}, 0, "both empty"},
        {{}, 0, {0}, 1, {0}, 1, "first empty"},
        {{1}, 1, {}, 0, {1}, 1, "second empty"},
        {{1, 3, 5}, 3, {2, 4, 6}, 3, {1, 2, 3, 4, 5, 6}, 6, "interleaved"},
        {{1, 2, 3}, 3, {4, 5, 6}, 3, {1, 2, 3, 4, 5, 6}, 6, "no overlap"},
        {{5}, 1, {1, 2, 3}, 3, {1, 2, 3, 5}, 4, "single vs multiple"},
    };

    int numTests = sizeof(tests) / sizeof(tests[0]);
    int allPassed = 1;

    for (int i = 0; i < numTests; i++) {
        struct ListNode* l1 = arrayToList(tests[i].list1, tests[i].size1);
        struct ListNode* l2 = arrayToList(tests[i].list2, tests[i].size2);
        struct ListNode* result = mergeTwoLists(l1, l2);

        if (!listMatchesArray(result, tests[i].expected, tests[i].expectedSize)) {
            printf("FAIL [%s]: got ", tests[i].desc);
            printList(result);
            printf("\n");
            allPassed = 0;
        } else {
            printf("PASS [%s]\n", tests[i].desc);
        }

        freeList(result);
    }

    if (allPassed) {
        printf("\nAll tests passed!\n");
    }

    return 0;
}
