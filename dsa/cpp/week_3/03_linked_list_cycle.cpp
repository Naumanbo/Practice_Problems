// Key Takeaways:
// 1. Two approaches: hash set (O(n) space) stores visited node pointers,
//    Floyd's cycle detection (O(1) space) uses fast/slow pointers.
//    Always prefer Floyd's when space is constrained.
//
// 2. In C++, use unordered_set<ListNode*> for O(1) average lookup.
//    unordered_set uses a hash table; set uses a red-black tree (O(log n)).
//    For pointer storage with no ordering needed, unordered_set is correct.
//
// 3. Floyd's: slow moves 1 step, fast moves 2. If they meet, cycle exists.
//    Always check fast != nullptr && fast->next != nullptr before advancing
//    to avoid dereferencing a null pointer (segfault).
//
// 4. This same Floyd's algorithm works on arrays (LeetCode #287) by treating
//    array values as next pointers — same pattern, different data structure.
//
// Complexity: Hash set — O(n) time, O(n) space. Floyd's — O(n) time, O(1) space.

// Compile: g++ -std=c++17 -o 03_linked_list_cycle 03_linked_list_cycle.cpp
#include <iostream>
#include <vector>
#include <string>
#include <set>
using namespace std;

struct ListNode {
    int val;
    ListNode* next;
    ListNode(int x) : val(x), next(nullptr) {}
};

bool hasCycleSet(ListNode* head) {
    // Your implementation
    set<ListNode*> seen;
    
    for (ListNode* ptr = head; ptr != nullptr; ptr = ptr->next) {
        if (seen.find(ptr) == seen.end()) {
            seen.insert(ptr);
        }
        else {
            return true;
        }
    }
    
    return false;
}

bool hasCycleFloyd(ListNode* head) {
    // Your implementation
    ListNode* slow = head;
    ListNode* fast = head;

    while (fast != nullptr && fast->next != nullptr) {
        slow = slow->next;
        fast = fast->next->next;

        if (slow == fast) {
            return true;
        }
    }

    return false;
}

ListNode* makeListWithCycle(const vector<int>& values, int cyclePos) {
    if (values.empty()) return nullptr;
    vector<ListNode*> nodes;
    for (int v : values) nodes.push_back(new ListNode(v));
    for (size_t i = 0; i < nodes.size() - 1; i++) nodes[i]->next = nodes[i + 1];
    if (cyclePos >= 0 && cyclePos < (int)nodes.size()) nodes.back()->next = nodes[cyclePos];
    return nodes[0];
}

vector<int> makeRange(int n) {
    vector<int> v(n);
    for (int i = 0; i < n; i++) v[i] = i;
    return v;
}

struct TestCase { vector<int> values; int cyclePos; bool expected; string desc; };

int main() {
    vector<TestCase> tests = {
        {{3, 2, 0, -4}, 1, true, "cycle at pos 1"},
        {{1, 2}, 0, true, "cycle at head"},
        {{1, 2, 3, 4}, 0, true, "cycle back to head"},
        {{1, 2, 3, 4}, 2, true, "cycle in middle"},
        {{1}, -1, false, "single node no cycle"},
        {{1, 2}, -1, false, "two nodes no cycle"},
        {{1, 2, 3, 4, 5}, -1, false, "five nodes no cycle"},
        {{}, -1, false, "empty list"},
        {{1}, 0, true, "self-loop"},
        {{1, 2, 3}, -1, false, "three nodes no cycle"},
        {makeRange(100), -1, false, "large list no cycle"},
        {makeRange(100), 50, true, "large list cycle at 50"},
        {makeRange(100), 0, true, "large list cycle at head"},
        {makeRange(100), 99, true, "large list self-loop tail"},
        {{1, 2, 3, 4, 5}, 4, true, "tail self-loop"},
        {{1, 2, 3, 4, 5}, 3, true, "cycle at second to last"},
        {{1, 2, 3, 4, 5}, 1, true, "cycle at pos 1 five nodes"},
        {{-1, -2, -3}, -1, false, "negative values no cycle"},
        {{-1, -2, -3}, 0, true, "negative values with cycle"},
        {{1, 2}, 1, true, "two nodes tail self-loop"},
    };

    struct Approach { string name; bool (*fn)(ListNode*); };
    vector<Approach> approaches = {{"HashSet", hasCycleSet}, {"Floyd's", hasCycleFloyd}};

    cout << "======================================================================" << endl;
    cout << "LINKED LIST CYCLE - Test Results" << endl;
    cout << "======================================================================" << endl;

    int total = tests.size();
    for (auto& a : approaches) {
        cout << "\n" << a.name << ":" << endl;
        int passed = 0;
        for (size_t i = 0; i < tests.size(); i++) {
            auto& tc = tests[i];
            ListNode* head = makeListWithCycle(tc.values, tc.cyclePos);
            bool result = a.fn(head);
            bool ok = (result == tc.expected);
            if (ok) passed++;
            cout << "  " << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tc.desc
                 << ": pos=" << tc.cyclePos << " -> " << (result ? "true" : "false") << endl;
        }
        cout << "  Result: " << passed << "/" << total << endl;
    }

    cout << "\n======================================================================" << endl;
    cout << "Questions:" << endl;
    cout << "1. Why does Floyd's guarantee fast meets slow?" << endl;
    cout << "2. Hash set vs Floyd's trade-off?" << endl;
    cout << "3. Find WHERE cycle begins (LeetCode #142)?" << endl;
    return 0;
}
