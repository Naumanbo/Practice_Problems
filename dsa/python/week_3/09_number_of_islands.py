"""
DSA Problem: Number of Islands

Tests: Graph traversal, BFS/DFS, 2D grid, connected components

Difficulty: Medium
Source: LeetCode #200

Problem:
Given an m x n 2D binary grid where '1' represents land and '0' represents water,
return the number of islands. An island is surrounded by water and formed by
connecting adjacent lands horizontally or vertically.

Constraints:
    - m == grid.length
    - n == grid[i].length
    - 1 <= m, n <= 300
    - grid[i][j] is '0' or '1'
"""

from typing import List


def num_islands(grid: List[List[str]]) -> int:
    pass


# =============================================================================
# Test Cases - Comprehensive
# =============================================================================
if __name__ == "__main__":
    test_cases = [
        {
            "grid": [["1","1","1"],["0","1","0"],["0","1","0"]],
            "expected": 1, "desc": "T-shape island"
        },
        {
            "grid": [["0","0"],["0","0"]],
            "expected": 0, "desc": "all water"
        },
        {
            "grid": [["1","1"],["1","1"]],
            "expected": 1, "desc": "all land"
        },
        {
            "grid": [["1","1","0","0","0"],["1","1","0","0","0"],["0","0","1","0","0"],["0","0","0","1","1"]],
            "expected": 3, "desc": "three islands"
        },
        {
            "grid": [["1"]],
            "expected": 1, "desc": "single land cell"
        },
        {
            "grid": [["0"]],
            "expected": 0, "desc": "single water cell"
        },
        {
            "grid": [["1","0","1","0","1"]],
            "expected": 3, "desc": "row of alternating"
        },
        {
            "grid": [["1"],["0"],["1"],["0"],["1"]],
            "expected": 3, "desc": "column of alternating"
        },
        {
            "grid": [["1","0"],["0","1"]],
            "expected": 2, "desc": "diagonal not connected"
        },
        {
            "grid": [["1","1","1","1"]],
            "expected": 1, "desc": "single row all land"
        },
        {
            "grid": [["1","1","0"],["0","1","0"],["0","1","1"]],
            "expected": 1, "desc": "L-shape connected"
        },
        {
            "grid": [["1","0","1"],["1","0","1"]],
            "expected": 2, "desc": "two columns of land"
        },
        {
            "grid": [["1","0","1"],["0","1","0"],["1","0","1"]],
            "expected": 5, "desc": "checkerboard"
        },
        {
            "grid": [["1","1","1"],["1","0","1"],["1","1","1"]],
            "expected": 1, "desc": "ring shape"
        },
        {
            "grid": [["1","1","1"],["0","0","0"]],
            "expected": 1, "desc": "top row island"
        },
        {
            "grid": [["1","0","1"],["0","0","0"],["1","0","1"]],
            "expected": 4, "desc": "four corners"
        },
    ]

    all_passed = True
    for tc in test_cases:
        import copy
        grid_copy = copy.deepcopy(tc["grid"])
        result = num_islands(grid_copy)
        if result != tc["expected"]:
            print(f"FAIL [{tc['desc']}]: num_islands(...) = {result}, expected {tc['expected']}")
            all_passed = False
        else:
            print(f"PASS [{tc['desc']}]")

    if all_passed:
        print("\nAll tests passed!")
