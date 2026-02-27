// Tests: Graph traversal, BFS/DFS, 2D grid, connected components
//
// Number of Islands (LeetCode #200)
// Count islands in a 2D grid where '1' is land and '0' is water.
// Islands are formed by connecting adjacent land cells horizontally/vertically.

#include <iostream>
#include <vector>
using namespace std;

int numIslands(vector<vector<char>>& grid) {
    return 0;
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

struct TestCase {
    vector<vector<char>> grid;
    int expected;
    string desc;
};

int main() {
    vector<TestCase> tests = {
        {{{'1','1','1'},{'0','1','0'},{'0','1','0'}},                                                             1, "T-shape island"},
        {{{'0','0'},{'0','0'}},                                                                                    0, "all water"},
        {{{'1','1'},{'1','1'}},                                                                                    1, "all land"},
        {{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}},                3, "three islands"},
        {{{'1'}},                                                                                                  1, "single land cell"},
        {{{'0'}},                                                                                                  0, "single water cell"},
        {{{'1','0','1','0','1'}},                                                                                  3, "row of alternating"},
        {{{'1'},{'0'},{'1'},{'0'},{'1'}},                                                                          3, "column of alternating"},
        {{{'1','0'},{'0','1'}},                                                                                    2, "diagonal not connected"},
        {{{'1','1','1','1'}},                                                                                      1, "single row all land"},
        {{{'1','1','0'},{'0','1','0'},{'0','1','1'}},                                                              1, "L-shape connected"},
        {{{'1','0','1'},{'1','0','1'}},                                                                            2, "two columns of land"},
        {{{'1','0','1'},{'0','1','0'},{'1','0','1'}},                                                              5, "checkerboard"},
        {{{'1','1','1'},{'1','0','1'},{'1','1','1'}},                                                              1, "ring shape"},
        {{{'1','1','1'},{'0','0','0'}},                                                                            1, "top row island"},
        {{{'1','0','1'},{'0','0','0'},{'1','0','1'}},                                                              4, "four corners"},
    };

    cout << "======================================================================" << endl;
    cout << "NUMBER OF ISLANDS - Test Results" << endl;
    cout << "======================================================================" << endl;

    int passed = 0;
    for (size_t i = 0; i < tests.size(); i++) {
        vector<vector<char>> gridCopy = tests[i].grid;
        int result = numIslands(gridCopy);
        bool ok = (result == tests[i].expected);
        if (ok) passed++;
        cout << (i + 1) << ". [" << (ok ? "PASS" : "FAIL") << "] " << tests[i].desc << endl;
    }

    cout << "======================================================================" << endl;
    cout << "Summary: " << passed << "/" << tests.size() << " passed" << endl;
    cout << "======================================================================" << endl;
    return 0;
}
