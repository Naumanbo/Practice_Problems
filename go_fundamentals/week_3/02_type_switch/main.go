package main

import "fmt"

// Tests: Type assertions, type switches, empty interface (any)
//
// 1. Implement Describe(v any) string
//    - Returns a string describing the type and value
//    - Handle: int, float64, string, bool, []int, nil
//    - For unknown types, return "unknown type: <type>"
//
// 2. Implement SafeAdd(a, b any) (float64, bool)
//    - Attempts to add two values as numbers
//    - Supports int, float64, and string (parsed as number)
//    - Returns (result, true) on success
//    - Returns (0, false) if either value can't be converted to a number
//
// 3. Implement TypeCount(values []any) map[string]int
//    - Returns a count of each type in the slice
//    - Use type names: "int", "float64", "string", "bool", "nil", "other"
//
// 4. Implement ExtractStrings(values []any) []string
//    - Returns only the string values from the slice
//    - Maintains original order

import "strconv"

// TODO: Implement Describe
// Examples:
//   Describe(42)        -> "int: 42"
//   Describe(3.14)      -> "float64: 3.14"
//   Describe("hello")   -> "string: hello (len=5)"
//   Describe(true)      -> "bool: true"
//   Describe([]int{1,2})-> "[]int: [1 2] (len=2)"
//   Describe(nil)       -> "nil"
func Describe(v any) string {
	return ""
}

// TODO: Implement SafeAdd
// Examples:
//   SafeAdd(1, 2)           -> (3.0, true)
//   SafeAdd(1.5, 2)         -> (3.5, true)
//   SafeAdd("3", 4)         -> (7.0, true)
//   SafeAdd("3.5", "2.5")   -> (6.0, true)
//   SafeAdd("hello", 1)     -> (0.0, false)
//   SafeAdd(nil, 1)         -> (0.0, false)
func SafeAdd(a, b any) (float64, bool) {
	return 0, false
}

// TODO: Implement TypeCount
func TypeCount(values []any) map[string]int {
	return nil
}

// TODO: Implement ExtractStrings
func ExtractStrings(values []any) []string {
	return nil
}

func main() {
	// Test Describe
	fmt.Println("=== Describe ===")
	fmt.Println(Describe(42))
	fmt.Println(Describe(3.14159))
	fmt.Println(Describe("hello"))
	fmt.Println(Describe(true))
	fmt.Println(Describe([]int{1, 2, 3}))
	fmt.Println(Describe(nil))

	// Test SafeAdd
	fmt.Println("\n=== SafeAdd ===")
	if result, ok := SafeAdd(10, 20); ok {
		fmt.Printf("10 + 20 = %.1f\n", result)
	}
	if result, ok := SafeAdd(1.5, 2.5); ok {
		fmt.Printf("1.5 + 2.5 = %.1f\n", result)
	}
	if result, ok := SafeAdd("3", 4); ok {
		fmt.Printf("\"3\" + 4 = %.1f\n", result)
	}
	if _, ok := SafeAdd("hello", 1); !ok {
		fmt.Println("\"hello\" + 1 = failed (expected)")
	}

	// Test TypeCount
	fmt.Println("\n=== TypeCount ===")
	mixed := []any{1, 2, "a", "b", "c", 3.14, true, nil, false}
	counts := TypeCount(mixed)
	fmt.Println(counts)
	// Expected: map[bool:2 float64:1 int:2 nil:1 string:3]

	// Test ExtractStrings
	fmt.Println("\n=== ExtractStrings ===")
	strings := ExtractStrings([]any{1, "hello", 3.14, "world", true, "!"})
	fmt.Println(strings) // [hello world !]

	// Run test cases
	allPassed := true

	// Describe tests
	if Describe(0) != "int: 0" {
		fmt.Println("FAIL: Describe(0)")
		allPassed = false
	}
	if Describe("") != "string:  (len=0)" {
		fmt.Println("FAIL: Describe empty string")
		allPassed = false
	}
	if Describe(nil) != "nil" {
		fmt.Println("FAIL: Describe(nil)")
		allPassed = false
	}

	// SafeAdd tests
	if r, ok := SafeAdd(1, 2); !ok || r != 3.0 {
		fmt.Println("FAIL: SafeAdd(1, 2)")
		allPassed = false
	}
	if r, ok := SafeAdd("10", "20"); !ok || r != 30.0 {
		fmt.Println("FAIL: SafeAdd string numbers")
		allPassed = false
	}
	if _, ok := SafeAdd([]int{1}, 1); ok {
		fmt.Println("FAIL: SafeAdd should fail for slice")
		allPassed = false
	}

	// TypeCount tests
	tc := TypeCount([]any{1, 1, 1})
	if tc["int"] != 3 {
		fmt.Println("FAIL: TypeCount all ints")
		allPassed = false
	}
	tc2 := TypeCount([]any{})
	if len(tc2) != 0 {
		fmt.Println("FAIL: TypeCount empty")
		allPassed = false
	}

	// ExtractStrings tests
	es := ExtractStrings([]any{"a", 1, "b"})
	if len(es) != 2 || es[0] != "a" || es[1] != "b" {
		fmt.Println("FAIL: ExtractStrings")
		allPassed = false
	}
	if len(ExtractStrings([]any{1, 2, 3})) != 0 {
		fmt.Println("FAIL: ExtractStrings no strings")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: strconv is imported for you to use in SafeAdd
var _ = strconv.ParseFloat
