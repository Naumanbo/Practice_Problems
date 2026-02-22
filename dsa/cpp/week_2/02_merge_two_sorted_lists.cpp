// Key Takeaways:
// 1. Same dummy head pattern as Python: fixed anchor + walking tail pointer.
//    Return dummy->next, not tail. The logic is identical across languages.
//
// 2. In C++, the dummy must be stack-allocated as an object first, then take
//    its address: `ListNode dummy_obj(0); ListNode* dummy = &dummy_obj;`
//    In Python, `dummy = ListNode()` gives you a pointer directly — no address needed.
//
// 3. C++ uses `->` for pointer member access (list1->val, tail->next).
//    Python uses `.` for everything (list1.val, tail.next).
//
// 4. C++ null checks are explicit: `list1 != nullptr`. Python uses truthiness:
//    `list1 and list2`. The remaining list is `tail.next = list1 or list2` in
//    Python, but requires an explicit if/else if in C++.
//
// Complexity: Time O(n + m), Space O(1) — nodes are reused, not copied.

// Tests: Linked lists, two pointers, iteration
//
// Merge Two Sorted Lists (LeetCode 21)
// Merge two sorted linked lists into one sorted list.

#include <iostream>
#include <vector>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
    // TODO: Implement solution
    // Hint: Use a dummy head node to simplify edge cases
    ListNode dummy_obj = ListNode(0);
    ListNode* dummy = &dummy_obj;
    ListNode* tail = dummy;

    // cout << "b4 while\n";
    while (list1 != nullptr && list2 != nullptr) {
        if (list1->val >= list2->val){
            tail->next = list2;
            list2 = list2->next;
        } else if (list2->val >= list1->val) {
            tail->next = list1;
            list1 = list1->next;
        }
        // cout << "while\n";

        tail = tail->next;
    }

    // cout  << "after while\n";

    if (list1 != nullptr) {
        tail->next = list1;
    }
    else if (list2 != nullptr) {
        tail->next = list2;
    }



    return dummy->next;
}

// Helper: convert vector to linked list
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

// Helper: convert linked list to vector
vector<int> listToVector(ListNode* head) {
    vector<int> result;
    while (head) {
        result.push_back(head->val);
        head = head->next;
    }
    return result;
}

// Helper: print vector
void printVector(const vector<int>& v) {
    cout << "[";
    for (size_t i = 0; i < v.size(); i++) {
        cout << v[i];
        if (i < v.size() - 1) cout << ", ";
    }
    cout << "]";
}

int main() {
    struct TestCase {
        vector<int> list1;
        vector<int> list2;
        vector<int> expected;
        string desc;
    };

    vector<TestCase> tests = {
        {{1, 2, 4}, {1, 3, 4}, {1, 1, 2, 3, 4, 4}, "basic merge"},
        {{}, {}, {}, "both empty"},
        {{}, {0}, {0}, "first empty"},
        {{1}, {}, {1}, "second empty"},
        {{1, 3, 5}, {2, 4, 6}, {1, 2, 3, 4, 5, 6}, "interleaved"},
        {{1, 2, 3}, {4, 5, 6}, {1, 2, 3, 4, 5, 6}, "no overlap"},
        {{5}, {1, 2, 3}, {1, 2, 3, 5}, "single vs multiple"},
        {{1, 1, 1}, {1, 1}, {1, 1, 1, 1, 1}, "all duplicates"},
        {{-3, -1, 0}, {-2, 5}, {-3, -2, -1, 0, 5}, "negative values"},
        {{1}, {2}, {1, 2}, "single element each"},
        {{2}, {1}, {1, 2}, "single elements reversed"},
        {{5, 5, 5}, {5, 5, 5}, {5, 5, 5, 5, 5, 5}, "all same values"},
        {{1, 2, 3}, {1, 2, 3}, {1, 1, 2, 2, 3, 3}, "identical lists"},
        {{1000000}, {-1000000}, {-1000000, 1000000}, "large values"},
        {{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, "no overlap longer"},
    };

    bool allPassed = true;
    for (auto& tc : tests) {
        ListNode* l1 = vectorToList(tc.list1);
        ListNode* l2 = vectorToList(tc.list2);
        vector<int> result = listToVector(mergeTwoLists(l1, l2));

        if (result != tc.expected) {
            cout << "FAIL [" << tc.desc << "]: got ";
            printVector(result);
            cout << ", expected ";
            printVector(tc.expected);
            cout << endl;
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
