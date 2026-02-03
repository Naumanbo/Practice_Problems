package main

import "fmt"

// Tests: Maps, iteration, two-value assignment, delete
//
// Implement the following map operations:
// - CharFrequency(s string) map[rune]int - count occurrences of each character
// - MostFrequent(m map[rune]int) (rune, int) - return the most frequent char and its count
// - Invert(m map[string]int) map[int][]string - invert a map (values become keys, keys become slice of values)
// - Merge(m1, m2 map[string]int) map[string]int - merge two maps (m2 overwrites m1 on conflict)

// TODO: Implement CharFrequency
func CharFrequency(s string) map[rune]int {
	return nil
}

// TODO: Implement MostFrequent (return 0, 0 for empty map)
func MostFrequent(m map[rune]int) (rune, int) {
	return 0, 0
}

// TODO: Implement Invert
// Example: {"a": 1, "b": 1, "c": 2} -> {1: ["a", "b"], 2: ["c"]}
func Invert(m map[string]int) map[int][]string {
	return nil
}

// TODO: Implement Merge
func Merge(m1, m2 map[string]int) map[string]int {
	return nil
}

func main() {
	// Test CharFrequency
	freq := CharFrequency("hello")
	fmt.Println("CharFrequency(\"hello\"):", freq)
	// Expected: map[e:1 h:1 l:2 o:1]

	// Test MostFrequent
	char, count := MostFrequent(freq)
	fmt.Printf("MostFrequent: '%c' appears %d times\n", char, count)
	// Expected: 'l' appears 2 times

	// Test Invert
	scores := map[string]int{"alice": 90, "bob": 85, "charlie": 90}
	inverted := Invert(scores)
	fmt.Println("Inverted:", inverted)
	// Expected: map[85:[bob] 90:[alice charlie]] (order may vary)

	// Test Merge
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	merged := Merge(m1, m2)
	fmt.Println("Merged:", merged)
	// Expected: map[a:1 b:3 c:4]

	// Run test cases
	allPassed := true

	// CharFrequency tests
	f := CharFrequency("aab")
	if f['a'] != 2 || f['b'] != 1 {
		fmt.Println("FAIL: CharFrequency(\"aab\")")
		allPassed = false
	}

	f = CharFrequency("")
	if len(f) != 0 {
		fmt.Println("FAIL: CharFrequency(\"\")")
		allPassed = false
	}

	// MostFrequent tests
	c, n := MostFrequent(map[rune]int{'x': 5, 'y': 3})
	if c != 'x' || n != 5 {
		fmt.Println("FAIL: MostFrequent")
		allPassed = false
	}

	c, n = MostFrequent(map[rune]int{})
	if c != 0 || n != 0 {
		fmt.Println("FAIL: MostFrequent empty map")
		allPassed = false
	}

	// Merge tests
	result := Merge(
		map[string]int{"x": 1},
		map[string]int{"x": 2, "y": 3},
	)
	if result["x"] != 2 || result["y"] != 3 {
		fmt.Println("FAIL: Merge")
		allPassed = false
	}

	// Test nil maps
	result = Merge(nil, map[string]int{"a": 1})
	if result["a"] != 1 {
		fmt.Println("FAIL: Merge with nil")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
