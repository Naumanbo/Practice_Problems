package main

import "fmt"

// Tests: Graph traversal, BFS/DFS, 2D grid, connected components
//
// Number of Islands (LeetCode #200)
// Count islands in a 2D grid where '1' is land and '0' is water.
// Islands are formed by connecting adjacent land cells horizontally/vertically.

func numIslands(grid [][]byte) int {
	return 0
}

// =============================================================================
// Test Cases - Comprehensive
// =============================================================================

type testCase struct {
	grid     [][]byte
	expected int
	desc     string
}

func copyGrid(g [][]byte) [][]byte {
	cp := make([][]byte, len(g))
	for i := range g {
		cp[i] = make([]byte, len(g[i]))
		copy(cp[i], g[i])
	}
	return cp
}

func main() {
	tests := []testCase{
		{[][]byte{{'1','1','1'},{'0','1','0'},{'0','1','0'}}, 1, "T-shape island"},
		{[][]byte{{'0','0'},{'0','0'}}, 0, "all water"},
		{[][]byte{{'1','1'},{'1','1'}}, 1, "all land"},
		{[][]byte{{'1','1','0','0','0'},{'1','1','0','0','0'},{'0','0','1','0','0'},{'0','0','0','1','1'}}, 3, "three islands"},
		{[][]byte{{'1'}}, 1, "single land cell"},
		{[][]byte{{'0'}}, 0, "single water cell"},
		{[][]byte{{'1','0','1','0','1'}}, 3, "row of alternating"},
		{[][]byte{{'1'},{'0'},{'1'},{'0'},{'1'}}, 3, "column of alternating"},
		{[][]byte{{'1','0'},{'0','1'}}, 2, "diagonal not connected"},
		{[][]byte{{'1','1','1','1'}}, 1, "single row all land"},
		{[][]byte{{'1','1','0'},{'0','1','0'},{'0','1','1'}}, 1, "L-shape connected"},
		{[][]byte{{'1','0','1'},{'1','0','1'}}, 2, "two columns of land"},
		{[][]byte{{'1','0','1'},{'0','1','0'},{'1','0','1'}}, 5, "checkerboard"},
		{[][]byte{{'1','1','1'},{'1','0','1'},{'1','1','1'}}, 1, "ring shape"},
		{[][]byte{{'1','1','1'},{'0','0','0'}}, 1, "top row island"},
		{[][]byte{{'1','0','1'},{'0','0','0'},{'1','0','1'}}, 4, "four corners"},
	}

	fmt.Println("======================================================================")
	fmt.Println("NUMBER OF ISLANDS - Test Results")
	fmt.Println("======================================================================")

	passed := 0
	for i, tc := range tests {
		result := numIslands(copyGrid(tc.grid))
		ok := result == tc.expected
		if ok {
			passed++
		}
		status := "FAIL"
		if ok {
			status = "PASS"
		}
		fmt.Printf("  %2d. [%s] %s\n", i+1, status, tc.desc)
	}

	fmt.Println("======================================================================")
	fmt.Printf("Summary: %d/%d passed\n", passed, len(tests))
	fmt.Println("======================================================================")
}
