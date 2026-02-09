package main

import (
	"fmt"
	"strings"
)

// Tests: Closures as state machines, function composition, middleware pattern
//
// These are patterns you'll use constantly in Go backend development.
//
// 1. Implement RateLimiter(maxCalls int, window []int) func() bool
//    - Returns a function that tracks call counts
//    - The window parameter is ignored (simplified: just count total calls)
//    - Returns true if under limit, false if limit reached
//    - Also implement Reset: the returned function should accept a bool parameter
//      Actually, return TWO functions: func() bool (try) and func() (reset)
//    Signature: RateLimiter(maxCalls int) (try func() bool, reset func())
//
// 2. Implement Pipeline(fns ...func(string) string) func(string) string
//    - Chains multiple string transformations left to right
//    - Pipeline(f, g, h)(x) == h(g(f(x)))
//    - Empty pipeline returns input unchanged
//
// 3. Implement Logger(prefix string) func(string)
//    - Returns a function that prints "[prefix] [count]: message"
//    - Count increments each call, starting at 1
//    - Example: log := Logger("APP"); log("started") -> "[APP] [1]: started"
//
// 4. Implement Iterator(s []int) (next func() (int, bool), hasNext func() bool)
//    - Returns two closures sharing state
//    - next() returns the current element and advances; bool=false if exhausted
//    - hasNext() returns true if more elements remain
//    - This is the Iterator pattern from Java/Python, built with closures
//
// 5. Implement Debounce(fn func(string), delay int) func(string)
//    - Returns a debounced version of fn
//    - Simplified (no actual timing): only calls fn if the argument
//      is DIFFERENT from the last call's argument
//    - First call always executes

// TODO: Implement RateLimiter
func RateLimiter(maxCalls int) (try func() bool, reset func()) {
	return nil, nil
}

// TODO: Implement Pipeline
func Pipeline(fns ...func(string) string) func(string) string {
	return nil
}

// TODO: Implement Logger
func Logger(prefix string) func(string) {
	return nil
}

// TODO: Implement Iterator
func Iterator(s []int) (next func() (int, bool), hasNext func() bool) {
	return nil, nil
}

// TODO: Implement Debounce
func Debounce(fn func(string)) func(string) {
	return nil
}

func main() {
	// Test RateLimiter
	fmt.Println("=== RateLimiter ===")
	try, reset := RateLimiter(3)
	fmt.Println(try()) // true (1/3)
	fmt.Println(try()) // true (2/3)
	fmt.Println(try()) // true (3/3)
	fmt.Println(try()) // false (exceeded)
	fmt.Println(try()) // false
	reset()
	fmt.Println(try()) // true (reset, 1/3 again)

	// Test Pipeline
	fmt.Println("\n=== Pipeline ===")
	upper := func(s string) string { return strings.ToUpper(s) }
	addBang := func(s string) string { return s + "!" }
	trim := func(s string) string { return strings.TrimSpace(s) }

	shout := Pipeline(trim, upper, addBang)
	fmt.Println(shout("  hello  ")) // HELLO!

	identity := Pipeline() // empty pipeline
	fmt.Println(identity("unchanged")) // unchanged

	single := Pipeline(upper)
	fmt.Println(single("test")) // TEST

	// Test Logger
	fmt.Println("\n=== Logger ===")
	log := Logger("APP")
	log("server started")     // [APP] [1]: server started
	log("listening on :8080") // [APP] [2]: listening on :8080
	log("request received")   // [APP] [3]: request received

	// Independent loggers
	dbLog := Logger("DB")
	dbLog("connected") // [DB] [1]: connected

	// Test Iterator
	fmt.Println("\n=== Iterator ===")
	next, hasNext := Iterator([]int{10, 20, 30})
	for hasNext() {
		val, _ := next()
		fmt.Print(val, " ")
	}
	fmt.Println() // 10 20 30

	// Exhausted iterator
	_, ok := next()
	fmt.Println("After exhaustion, ok:", ok) // false

	// Empty iterator
	_, hasNextEmpty := Iterator([]int{})
	fmt.Println("Empty hasNext:", hasNextEmpty()) // false

	// Test Debounce
	fmt.Println("\n=== Debounce ===")
	var calls []string
	fn := func(s string) { calls = append(calls, s) }
	debounced := Debounce(fn)
	debounced("a") // calls fn("a")
	debounced("a") // skipped (same as last)
	debounced("a") // skipped
	debounced("b") // calls fn("b")
	debounced("b") // skipped
	debounced("a") // calls fn("a") (different from last)
	fmt.Println("Debounced calls:", calls) // [a b a]

	// Run test cases
	allPassed := true

	// RateLimiter limit of 1
	t1, r1 := RateLimiter(1)
	if !t1() {
		fmt.Println("FAIL: RateLimiter first call")
		allPassed = false
	}
	if t1() {
		fmt.Println("FAIL: RateLimiter should block after 1")
		allPassed = false
	}
	r1()
	if !t1() {
		fmt.Println("FAIL: RateLimiter after reset")
		allPassed = false
	}

	// Pipeline order matters
	addX := func(s string) string { return s + "X" }
	addY := func(s string) string { return s + "Y" }
	if Pipeline(addX, addY)("") != "XY" {
		fmt.Println("FAIL: Pipeline order, expected XY")
		allPassed = false
	}
	if Pipeline(addY, addX)("") != "YX" {
		fmt.Println("FAIL: Pipeline reverse order, expected YX")
		allPassed = false
	}

	// Iterator single element
	n, hn := Iterator([]int{42})
	if !hn() {
		fmt.Println("FAIL: Iterator single hasNext")
		allPassed = false
	}
	v, ok2 := n()
	if v != 42 || !ok2 {
		fmt.Println("FAIL: Iterator single next")
		allPassed = false
	}
	if hn() {
		fmt.Println("FAIL: Iterator single exhausted hasNext")
		allPassed = false
	}

	// Debounce first call always executes
	var dc []string
	db := Debounce(func(s string) { dc = append(dc, s) })
	db("first")
	if len(dc) != 1 || dc[0] != "first" {
		fmt.Println("FAIL: Debounce first call")
		allPassed = false
	}

	// Debounce with alternating values
	var dc2 []string
	db2 := Debounce(func(s string) { dc2 = append(dc2, s) })
	db2("a")
	db2("b")
	db2("a")
	db2("b")
	if len(dc2) != 4 {
		fmt.Println("FAIL: Debounce alternating should call all 4")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: strings package imported for your use
var _ = strings.ToUpper
