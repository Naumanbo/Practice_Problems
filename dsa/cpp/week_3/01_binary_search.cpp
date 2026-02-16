// Compile: g++ -std=c++17 -o 01_binary_search 01_binary_search.cpp
#include <iostream>
#include <vector>
#include <string>
using namespace std;

int binarySearchIterative(vector<int>& nums, int target) {
    // Your implementation
    return -1;
}

int binarySearchRecursive(vector<int>& nums, int target) {
    // Your implementation
    return -1;
}

struct TestCase { vector<int> nums; int target; int expected; string desc; };

vector<int> makeRange(int s, int e) {
    vector<int> v;
    for (int i = s; i <= e; i++) v.push_back(i);
    return v;
}

int runTests(int (*fn)(vector<int>&, int), const string& name, vector<TestCase>& t) {
    cout << "\n" << name << ":" << endl;
    int p = 0;
    for (size_t i = 0; i < t.size(); i++) {
        vector<int> c = t[i].nums;
        int r = fn(c, t[i].target);
        bool ok = (r == t[i].expected);
        if (ok) p++;
        cout << "  " << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] "
             << t[i].desc << ": target=" << t[i].target
             << " -> " << r << " (expected " << t[i].expected << ")" << endl;
    }
    return p;
}

int main() {
    vector<TestCase> tests = {
        {{-1, 0, 3, 5, 9, 12}, 9, 4, "target in middle"},
        {{-1, 0, 3, 5, 9, 12}, 2, -1, "target not found"},
        {{1, 2, 3, 4, 5}, 1, 0, "target at start"},
        {{1, 2, 3, 4, 5}, 5, 4, "target at end"},
        {{1, 2, 3, 4, 5}, 3, 2, "target in center"},
        {{5}, 5, 0, "single element found"},
        {{5}, 3, -1, "single element not found"},
        {{1, 2}, 1, 0, "two elements first"},
        {{1, 2}, 2, 1, "two elements second"},
        {{1, 2}, 3, -1, "two elements not found"},
        {{-10, -5, 0, 5, 10}, -5, 1, "negative target"},
        {{-10, -5, 0, 5, 10}, 0, 2, "zero target"},
        {{-100, -50, -10, -1}, -100, 0, "all negative first"},
        {{-100, -50, -10, -1}, -1, 3, "all negative last"},
        {{-100, -50, -10, -1}, 5, -1, "all negative not found"},
        {{1, 2, 3, 4, 5}, 0, -1, "target below range"},
        {{1, 2, 3, 4, 5}, 6, -1, "target above range"},
        {makeRange(1, 100), 50, 49, "larger array middle"},
        {makeRange(1, 100), 1, 0, "larger array first"},
        {makeRange(1, 100), 100, 99, "larger array last"},
        {makeRange(1, 100), 101, -1, "larger array not found"},
        {{-9999, 0, 9999}, 9999, 2, "near constraint bounds"},
    };

    cout << "======================================================================" << endl;
    cout << "BINARY SEARCH - Test Results" << endl;
    cout << "======================================================================" << endl;

    int total = tests.size();
    int ip = runTests(binarySearchIterative, "Iterative", tests);
    int rp = runTests(binarySearchRecursive, "Recursive", tests);

    cout << "\n======================================================================" << endl;
    cout << "Summary: Iterative " << ip << "/" << total
         << " | Recursive " << rp << "/" << total << endl;
    cout << "======================================================================" << endl;
    cout << "\nQuestions:" << endl;
    cout << "1. Why must the array be sorted?" << endl;
    cout << "2. Integer overflow risk with (left+right)/2?" << endl;
    cout << "3. How does std::lower_bound relate?" << endl;
    return 0;
}
