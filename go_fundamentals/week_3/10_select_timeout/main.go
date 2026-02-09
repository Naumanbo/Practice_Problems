package main

import (
	"fmt"
	"time"
)

// Tests: Select statement, timeouts, non-blocking channel operations, tickers
//
// 1. Implement RaceTwo(a, b func() string) string
//    - Run both functions in separate goroutines
//    - Each goroutine sends its result to a channel
//    - Use select to return whichever finishes first
//
// 2. Implement WithTimeout(work func() int, timeout time.Duration) (int, bool)
//    - Run work in a goroutine
//    - If work finishes before timeout, return (result, true)
//    - If timeout fires first, return (0, false)
//    - Hint: use time.After(timeout)
//
// 3. Implement Heartbeat(interval time.Duration, done <-chan struct{}) <-chan string
//    - Returns a channel that emits "beat" at regular intervals
//    - Stops when done channel is closed
//    - Use time.NewTicker for the interval
//    - Remember to stop the ticker when done
//
// 4. Implement CollectWithDeadline(ch <-chan int, deadline time.Duration) []int
//    - Collect values from ch until either:
//      a) The channel is closed, OR
//      b) The deadline expires
//    - Return whatever was collected

// TODO: Implement RaceTwo
func RaceTwo(a, b func() string) string {
	return ""
}

// TODO: Implement WithTimeout
func WithTimeout(work func() int, timeout time.Duration) (int, bool) {
	return 0, false
}

// TODO: Implement Heartbeat
func Heartbeat(interval time.Duration, done <-chan struct{}) <-chan string {
	return nil
}

// TODO: Implement CollectWithDeadline
func CollectWithDeadline(ch <-chan int, deadline time.Duration) []int {
	return nil
}

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

	// RaceTwo - verify faster wins
	w := RaceTwo(
		func() string { time.Sleep(50 * time.Millisecond); return "A" },
		func() string { time.Sleep(5 * time.Millisecond); return "B" },
	)
	if w != "B" {
		fmt.Println("FAIL: RaceTwo faster should win")
		allPassed = false
	}

	// WithTimeout success
	r, o := WithTimeout(func() int { return 7 }, time.Second)
	if !o || r != 7 {
		fmt.Println("FAIL: WithTimeout instant work")
		allPassed = false
	}

	// WithTimeout failure
	_, o = WithTimeout(func() int {
		time.Sleep(500 * time.Millisecond)
		return 0
	}, 10*time.Millisecond)
	if o {
		fmt.Println("FAIL: WithTimeout should have timed out")
		allPassed = false
	}

	// CollectWithDeadline empty channel closes immediately
	emptyCh := make(chan int)
	close(emptyCh)
	if result := CollectWithDeadline(emptyCh, time.Second); len(result) != 0 {
		fmt.Println("FAIL: CollectWithDeadline closed channel")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
