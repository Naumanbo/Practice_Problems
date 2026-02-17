package main

import (
	"fmt"
	"time"
)

// Exercise: Select, Timeouts & Non-Blocking Channel Ops
//
// Build concurrency patterns using Go's select statement.
// Design all types and functions from scratch.
//
// What to build:
// - RaceTwo: run two functions concurrently, return whichever finishes first
// - WithTimeout: run work in a goroutine, return result or timeout
// - Heartbeat: emit "beat" at regular intervals until a done signal
// - CollectWithDeadline: collect values from a channel until it closes or deadline expires
//
// Key tools:
// - select {} for multiplexing channels
// - time.After(duration) returns a channel that fires once after duration
// - time.NewTicker(duration) for repeated signals (remember to Stop() it)
//
// Read the tests in main() to understand exact function signatures
// and expected behavior.

// === WRITE YOUR CODE BELOW ===

// === END YOUR CODE ===

func main() {
	// Test RaceTwo
	fmt.Println("=== RaceTwo ===")
	winner := RaceTwo(
		func() string {
			time.Sleep(100 * time.Millisecond)
			return "slow"
		},
		func() string {
			time.Sleep(10 * time.Millisecond)
			return "fast"
		},
	)
	fmt.Println("Winner:", winner) // fast

	// Test WithTimeout - work finishes in time
	fmt.Println("\n=== WithTimeout (success) ===")
	result, ok := WithTimeout(func() int {
		time.Sleep(10 * time.Millisecond)
		return 42
	}, 100*time.Millisecond)
	fmt.Printf("Result: %d, OK: %v\n", result, ok) // 42, true

	// Test WithTimeout - work too slow
	fmt.Println("\n=== WithTimeout (timeout) ===")
	result, ok = WithTimeout(func() int {
		time.Sleep(200 * time.Millisecond)
		return 99
	}, 50*time.Millisecond)
	fmt.Printf("Result: %d, OK: %v\n", result, ok) // 0, false

	// Test Heartbeat
	fmt.Println("\n=== Heartbeat ===")
	done := make(chan struct{})
	beats := Heartbeat(50*time.Millisecond, done)

	beatCount := 0
	timeout := time.After(180 * time.Millisecond)
	func() {
		for {
			select {
			case msg, ok := <-beats:
				if !ok {
					return
				}
				beatCount++
				fmt.Printf("  %s #%d\n", msg, beatCount)
			case <-timeout:
				close(done)
				// Drain remaining
				time.Sleep(10 * time.Millisecond)
				return
			}
		}
	}()
	fmt.Printf("Got %d beats (expected ~3)\n", beatCount)

	// Test CollectWithDeadline - channel closes first
	fmt.Println("\n=== CollectWithDeadline (channel closes) ===")
	ch := make(chan int)
	go func() {
		for _, v := range []int{10, 20, 30} {
			ch <- v
			time.Sleep(10 * time.Millisecond)
		}
		close(ch)
	}()
	collected := CollectWithDeadline(ch, 1*time.Second)
	fmt.Println("Collected:", collected) // [10 20 30]

	// Test CollectWithDeadline - deadline fires first
	fmt.Println("\n=== CollectWithDeadline (deadline) ===")
	slowCh := make(chan int)
	go func() {
		for i := 0; ; i++ {
			slowCh <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	partial := CollectWithDeadline(slowCh, 130*time.Millisecond)
	fmt.Printf("Collected %d values before deadline\n", len(partial))

	// Run test cases
	allPassed := true

	// === RaceTwo tests ===
	w := RaceTwo(
		func() string { time.Sleep(50 * time.Millisecond); return "A" },
		func() string { time.Sleep(5 * time.Millisecond); return "B" },
	)
	if w != "B" {
		fmt.Println("FAIL: RaceTwo faster should win")
		allPassed = false
	}
	// first function faster
	w2 := RaceTwo(
		func() string { time.Sleep(5 * time.Millisecond); return "first" },
		func() string { time.Sleep(100 * time.Millisecond); return "second" },
	)
	if w2 != "first" {
		fmt.Println("FAIL: RaceTwo first faster")
		allPassed = false
	}
	// both instant — either is acceptable
	w3 := RaceTwo(
		func() string { return "X" },
		func() string { return "Y" },
	)
	if w3 != "X" && w3 != "Y" {
		fmt.Println("FAIL: RaceTwo both instant should return one")
		allPassed = false
	}

	// === WithTimeout tests ===
	// instant work
	r, o := WithTimeout(func() int { return 7 }, time.Second)
	if !o || r != 7 {
		fmt.Println("FAIL: WithTimeout instant work")
		allPassed = false
	}
	// timeout
	_, o = WithTimeout(func() int {
		time.Sleep(500 * time.Millisecond)
		return 0
	}, 10*time.Millisecond)
	if o {
		fmt.Println("FAIL: WithTimeout should have timed out")
		allPassed = false
	}
	// zero return value (verify 0 is returned as result, not confused with timeout)
	r2, o2 := WithTimeout(func() int { return 0 }, time.Second)
	if !o2 || r2 != 0 {
		fmt.Println("FAIL: WithTimeout returning 0 should succeed")
		allPassed = false
	}
	// negative return value
	r3, o3 := WithTimeout(func() int { return -42 }, time.Second)
	if !o3 || r3 != -42 {
		fmt.Println("FAIL: WithTimeout negative return")
		allPassed = false
	}

	// === CollectWithDeadline tests ===
	// closed channel immediately
	emptyCh := make(chan int)
	close(emptyCh)
	if cResult := CollectWithDeadline(emptyCh, time.Second); len(cResult) != 0 {
		fmt.Println("FAIL: CollectWithDeadline closed channel")
		allPassed = false
	}
	// fast channel closes before deadline
	fastCh := make(chan int)
	go func() {
		fastCh <- 1
		fastCh <- 2
		fastCh <- 3
		close(fastCh)
	}()
	fastResult := CollectWithDeadline(fastCh, time.Second)
	if len(fastResult) != 3 || fastResult[0] != 1 || fastResult[2] != 3 {
		fmt.Println("FAIL: CollectWithDeadline fast channel")
		allPassed = false
	}
	// deadline fires before channel closes
	slowCh2 := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			slowCh2 <- i
			time.Sleep(50 * time.Millisecond)
		}
	}()
	deadlineResult := CollectWithDeadline(slowCh2, 120*time.Millisecond)
	if len(deadlineResult) < 1 || len(deadlineResult) > 5 {
		fmt.Printf("FAIL: CollectWithDeadline should collect ~2-3, got %d\n", len(deadlineResult))
		allPassed = false
	}
	// single value then close
	singleCh := make(chan int, 1)
	singleCh <- 42
	close(singleCh)
	singleResult := CollectWithDeadline(singleCh, time.Second)
	if len(singleResult) != 1 || singleResult[0] != 42 {
		fmt.Println("FAIL: CollectWithDeadline single value")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
