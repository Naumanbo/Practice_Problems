package main

import (
	"fmt"
	"strconv"
)

// Tests: Type assertions, type switches, empty interface (any)
//
// KEY TAKEAWAYS:
// - Type switch syntax: switch v := v.(type) { case int: ... }
//   The := reassignment gives you the concrete type inside each case — same value, just unwrapped.
//   Without it, v is still `any` and you can't call type-specific operations (len, float64(), etc).
// - You CAN omit the := if you only need to know the type (e.g., TypeCount) — just switch v.(type).
// - Go can't convert string to float64 directly — use strconv.ParseFloat(s, 64).
//   int <-> float64 conversions work with direct type conversion: float64(x), int(x).
// - strconv.ParseFloat returns (float64, error), not a bool. Check err != nil for failure.
// - %f gives 6 decimal places (3.140000). Use %g to drop trailing zeros (3.14).
// - %T (capital T) prints a value's Go type name. %t (lowercase) is for booleans only.
// - Variables declared inside an if/switch/case block are scoped to that block.
//   Declare outside or return from within to use values after the block.
// - Extract repeated logic into helper functions (toFloat) to avoid scoping issues
//   and reduce duplication across type switch cases.
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

// TODO: Implement Describe
// Examples:
//
//	Describe(42)        -> "int: 42"
//	Describe(3.14)      -> "float64: 3.14"
//	Describe("hello")   -> "string: hello (len=5)"
//	Describe(true)      -> "bool: true"
//	Describe([]int{1,2})-> "[]int: [1 2] (len=2)"
//	Describe(nil)       -> "nil"
func Describe(v any) string {
	switch v := v.(type) {
	case int:
		fmt.Printf("int: %d\n", v)
		return fmt.Sprintf("int: %d", v)
	case float64:
		fmt.Printf("float64: %g\n", v)
		return fmt.Sprintf("float64: %g", v)
	case string:
		fmt.Printf("string: %s (len=%d)\n", v, len(v))
		return fmt.Sprintf("string: %s (len=%d)", v, len(v))
	case bool:
		fmt.Printf("bool: %t\n", v)
		return fmt.Sprintf("bool: %t", v)
	case []int:
		fmt.Printf("[]int: %v (len=%d)\n", v, len(v))
		return fmt.Sprintf("[]int: %v (len=%d)", v, len(v))

	case nil:
		fmt.Println("nil")
		return "nil"

	default:
		fmt.Printf("unknown type: %T\n", v)
		return fmt.Sprintf("unknown type: %T", v)
	}

}

// TODO: Implement SafeAdd
// Examples:
//
//		SafeAdd(1, 2)           -> (3.0, true)
//		SafeAdd(1.5, 2)         -> (3.5, true)
//		SafeAdd("3", 4)         -> (7.0, true)
//		SafeAdd("3.5", "2.5")   -> (6.0, true)
//		SafeAdd("hello", 1)     -> (0.0, false)
//		SafeAdd(nil, 1)         -> (0.0, false)
//	   - Supports int, float64, and string (parsed as number)
func SafeAdd(a, b any) (float64, bool) {
	valA, okA := toFloat(a)
	valB, okB := toFloat(b)

	if !okA || !okB {
		return 0, false
	}
	return valA + valB, true
}

// Helper function for SafeAdd type conversion
func toFloat(v any) (float64, bool) {
	switch v := v.(type) {
	case string:
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0, false
		}
		return val, true
	case int:
		return float64(v), true

	case float64:
		return v, true

	default:
		return 0, false
	}
}

// TODO: Implement TypeCount
func TypeCount(values []any) map[string]int {
	var counts map[string]int = make(map[string]int)

	for _, elt := range values {
		switch elt.(type) {
		case string:
			counts["string"] += 1
		case int:
			counts["int"] += 1
		case float64:
			counts["float64"] += 1
		case bool:
			counts["bool"] += 1
		case nil:
			counts["nil"] += 1
		default:
			counts["other"] += 1
		}
	}

	return counts
}

// TODO: Implement ExtractStrings
//   - Returns only the string values from the slice
//   - Maintains original order
func ExtractStrings(values []any) []string {
	stringElts := []string{}
	for _, val := range values {
		switch val := val.(type) {
		case string:
			stringElts = append(stringElts, val)

		}
	}

	return stringElts
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

	// === Describe tests ===
	// int
	if Describe(0) != "int: 0" {
		fmt.Println("FAIL: Describe(0)")
		allPassed = false
	}
	if Describe(42) != "int: 42" {
		fmt.Println("FAIL: Describe(42)")
		allPassed = false
	}
	if Describe(-7) != "int: -7" {
		fmt.Println("FAIL: Describe(-7)")
		allPassed = false
	}
	// float64
	if Describe(3.14) != "float64: 3.14" {
		fmt.Println("FAIL: Describe(3.14)")
		allPassed = false
	}
	if Describe(0.0) != "float64: 0" {
		fmt.Println("FAIL: Describe(0.0)")
		allPassed = false
	}
	// string
	if Describe("") != "string:  (len=0)" {
		fmt.Println("FAIL: Describe empty string")
		allPassed = false
	}
	if Describe("hello") != "string: hello (len=5)" {
		fmt.Println("FAIL: Describe(hello)")
		allPassed = false
	}
	if Describe("a") != "string: a (len=1)" {
		fmt.Println("FAIL: Describe single char string")
		allPassed = false
	}
	// bool
	if Describe(true) != "bool: true" {
		fmt.Println("FAIL: Describe(true)")
		allPassed = false
	}
	if Describe(false) != "bool: false" {
		fmt.Println("FAIL: Describe(false)")
		allPassed = false
	}
	// []int
	if Describe([]int{1, 2, 3}) != "[]int: [1 2 3] (len=3)" {
		fmt.Println("FAIL: Describe([]int)")
		allPassed = false
	}
	if Describe([]int{}) != "[]int: [] (len=0)" {
		fmt.Println("FAIL: Describe empty []int")
		allPassed = false
	}
	if Describe([]int{42}) != "[]int: [42] (len=1)" {
		fmt.Println("FAIL: Describe single []int")
		allPassed = false
	}
	// nil
	if Describe(nil) != "nil" {
		fmt.Println("FAIL: Describe(nil)")
		allPassed = false
	}
	// unknown type
	if Describe([]string{"a"}) != fmt.Sprintf("unknown type: %T", []string{"a"}) {
		fmt.Println("FAIL: Describe unknown type")
		allPassed = false
	}

	// === SafeAdd tests ===
	// int + int
	if r, ok := SafeAdd(1, 2); !ok || r != 3.0 {
		fmt.Println("FAIL: SafeAdd(1, 2)")
		allPassed = false
	}
	// float + float
	if r, ok := SafeAdd(1.5, 2.5); !ok || r != 4.0 {
		fmt.Println("FAIL: SafeAdd(1.5, 2.5)")
		allPassed = false
	}
	// int + float
	if r, ok := SafeAdd(1, 2.5); !ok || r != 3.5 {
		fmt.Println("FAIL: SafeAdd(1, 2.5)")
		allPassed = false
	}
	// string + string
	if r, ok := SafeAdd("10", "20"); !ok || r != 30.0 {
		fmt.Println("FAIL: SafeAdd string numbers")
		allPassed = false
	}
	// string + int
	if r, ok := SafeAdd("3.5", 2); !ok || r != 5.5 {
		fmt.Println("FAIL: SafeAdd string float + int")
		allPassed = false
	}
	// invalid string
	if _, ok := SafeAdd("hello", 1); ok {
		fmt.Println("FAIL: SafeAdd invalid string should fail")
		allPassed = false
	}
	// nil
	if _, ok := SafeAdd(nil, 1); ok {
		fmt.Println("FAIL: SafeAdd nil should fail")
		allPassed = false
	}
	if _, ok := SafeAdd(1, nil); ok {
		fmt.Println("FAIL: SafeAdd second arg nil should fail")
		allPassed = false
	}
	// unsupported type
	if _, ok := SafeAdd([]int{1}, 1); ok {
		fmt.Println("FAIL: SafeAdd should fail for slice")
		allPassed = false
	}
	if _, ok := SafeAdd(true, 1); ok {
		fmt.Println("FAIL: SafeAdd should fail for bool")
		allPassed = false
	}
	// zero values
	if r, ok := SafeAdd(0, 0); !ok || r != 0.0 {
		fmt.Println("FAIL: SafeAdd(0, 0)")
		allPassed = false
	}
	// negative numbers
	if r, ok := SafeAdd(-5, 3); !ok || r != -2.0 {
		fmt.Println("FAIL: SafeAdd(-5, 3)")
		allPassed = false
	}

	// === TypeCount tests ===
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
	// all types present
	tc3 := TypeCount([]any{1, 3.14, "hi", true, nil, []int{1}})
	if tc3["int"] != 1 || tc3["float64"] != 1 || tc3["string"] != 1 || tc3["bool"] != 1 || tc3["nil"] != 1 || tc3["other"] != 1 {
		fmt.Println("FAIL: TypeCount all types")
		allPassed = false
	}
	// only nils
	tc4 := TypeCount([]any{nil, nil, nil})
	if tc4["nil"] != 3 || len(tc4) != 1 {
		fmt.Println("FAIL: TypeCount all nils")
		allPassed = false
	}
	// single element
	tc5 := TypeCount([]any{"solo"})
	if tc5["string"] != 1 || len(tc5) != 1 {
		fmt.Println("FAIL: TypeCount single string")
		allPassed = false
	}

	// === ExtractStrings tests ===
	es := ExtractStrings([]any{"a", 1, "b"})
	if len(es) != 2 || es[0] != "a" || es[1] != "b" {
		fmt.Println("FAIL: ExtractStrings")
		allPassed = false
	}
	if len(ExtractStrings([]any{1, 2, 3})) != 0 {
		fmt.Println("FAIL: ExtractStrings no strings")
		allPassed = false
	}
	// empty input
	if len(ExtractStrings([]any{})) != 0 {
		fmt.Println("FAIL: ExtractStrings empty")
		allPassed = false
	}
	// all strings
	esAll := ExtractStrings([]any{"x", "y", "z"})
	if len(esAll) != 3 || esAll[0] != "x" || esAll[1] != "y" || esAll[2] != "z" {
		fmt.Println("FAIL: ExtractStrings all strings")
		allPassed = false
	}
	// preserves order with mixed types
	esOrder := ExtractStrings([]any{nil, "first", 42, "second", true, "third"})
	if len(esOrder) != 3 || esOrder[0] != "first" || esOrder[1] != "second" || esOrder[2] != "third" {
		fmt.Println("FAIL: ExtractStrings order preservation")
		allPassed = false
	}
	// empty strings are still strings
	esEmpty := ExtractStrings([]any{"", "a", ""})
	if len(esEmpty) != 3 || esEmpty[0] != "" || esEmpty[1] != "a" || esEmpty[2] != "" {
		fmt.Println("FAIL: ExtractStrings with empty strings")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: strconv is imported for you to use in SafeAdd
var _ = strconv.ParseFloat
