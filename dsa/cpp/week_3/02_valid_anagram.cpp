// Compile: g++ -std=c++17 -o 02_valid_anagram 02_valid_anagram.cpp
#include <iostream>
#include <string>
#include <vector>
using namespace std;

bool isAnagramSort(string s, string t) {
    // Your implementation
    return false;
}

bool isAnagramMap(string s, string t) {
    // Your implementation
    return false;
}

struct TestCase { string s; string t; bool expected; string desc; };

int main() {
    vector<TestCase> tests = {
        {"anagram", "nagaram", true, "classic anagram"},
        {"rat", "car", false, "not anagram"},
        {"listen", "silent", true, "listen/silent"},
        {"hello", "world", false, "different letters"},
        {"a", "a", true, "single char same"},
        {"a", "b", false, "single char different"},
        {"ab", "ba", true, "two chars swapped"},
        {"ab", "cd", false, "two chars different"},
        {"abc", "ab", false, "different lengths"},
        {"a", "ab", false, "subset string"},
        {"aaa", "aaa", true, "all same chars"},
        {"aab", "baa", true, "repeated chars anagram"},
        {"aacc", "ccac", false, "same chars wrong count"},
        {"aabb", "abab", true, "interleaved duplicates"},
        {"abcde", "abcdf", false, "one char different"},
        {"abcd", "abce", false, "last char differs"},
        {"aaab", "aaba", true, "rearranged with repeats"},
        {string(100, 'a'), string(100, 'a'), true, "long same string"},
        {string(99, 'a') + "b", "b" + string(99, 'a'), true, "long anagram"},
        {string(100, 'a'), string(99, 'a') + "b", false, "long near miss"},
        {"abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true, "full alphabet reversed"},
    };

    cout << "======================================================================" << endl;
    cout << "VALID ANAGRAM - Test Results" << endl;
    cout << "======================================================================" << endl;

    struct Approach { string name; bool (*fn)(string, string); };
    vector<Approach> approaches = {{"Sorting", isAnagramSort}, {"HashMap", isAnagramMap}};
    int total = tests.size();

    for (auto& a : approaches) {
        cout << "\n" << a.name << ":" << endl;
        int passed = 0;
        for (size_t i = 0; i < tests.size(); i++) {
            auto& tc = tests[i];
            bool result = a.fn(tc.s, tc.t);
            bool ok = (result == tc.expected);
            if (ok) passed++;
            string sd = tc.s.length() <= 15 ? tc.s : tc.s.substr(0, 12) + "...";
            string td = tc.t.length() <= 15 ? tc.t : tc.t.substr(0, 12) + "...";
            cout << "  " << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tc.desc
                 << ": '" << sd << "','" << td << "' -> "
                 << (result ? "true" : "false") << endl;
        }
        cout << "  Result: " << passed << "/" << total << endl;
    }

    cout << "\n======================================================================" << endl;
    cout << "Questions:" << endl;
    cout << "1. Which approach is better for Unicode?" << endl;
    cout << "2. Can you use int[26] array?" << endl;
    cout << "3. Space complexity of sorting?" << endl;
    return 0;
}
