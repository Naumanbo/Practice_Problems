package main

import (
	"fmt"
	"math"
	"sync"
)

// Tests: Goroutines, channels, sending/receiving, basic concurrency
//
// 1. Implement Squarer(nums []int) []int
//    - Launch a goroutine that sends each number squared into a channel
//    - Main function receives from the channel and collects results
//    - Return the squared values (order doesn't need to match input)
//
// 2. Implement ParallelSum(slices [][]int) int
//    - Split work: each slice gets its own goroutine to compute its sum
//    - Each goroutine sends its partial sum to a shared channel
//    - Collect and return the total sum
//
// 3. Implement IsPrimeParallel(nums []int) map[int]bool
//    - Check each number for primality in its own goroutine
//    - Use a channel to send back results as (number, isPrime) pairs
//    - Return a map of number -> isPrime
//    - Hint: define a small result struct to send through the channel
//
// 4. Implement Merge(ch1, ch2 <-chan int) <-chan int
//    - Merge two input channels into one output channel
//    - Use a goroutine + sync.WaitGroup to read from both and forward
//    - Close the output channel when both inputs are exhausted

// TODO: Implement Squarer
func Squarer(nums []int) []int {
	return nil
}

// TODO: Implement ParallelSum
func ParallelSum(slices [][]int) int {
	return 0
}

// Hint: use this struct for IsPrimeParallel results
// type primeResult struct {
// 	num     int
// 	isPrime bool
// }

// TODO: Implement IsPrimeParallel
func IsPrimeParallel(nums []int) map[int]bool {
	return nil
}

// isPrime checks if a number is prime (helper - already done for you)
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// TODO: Implement Merge
func Merge(ch1, ch2 <-chan int) <-chan int {
	return nil
}

func main() {
	// Test Squarer
	fmt.Println("=== Squarer ===")
	squared := Squarer([]int{1, 2, 3, 4, 5})
	fmt.Println("Squared:", squared) // [1 4 9 16 25] (order may vary)

	// Test ParallelSum
	fmt.Println("\n=== ParallelSum ===")
	slices := [][]int{
		{1, 2, 3},       // 6
		{4, 5, 6},       // 15
		{7, 8, 9, 10},   // 34
	}
	total := ParallelSum(slices)
	fmt.Println("Total sum:", total) // 55

	// Test IsPrimeParallel
	fmt.Println("\n=== IsPrimeParallel ===")
	primes := IsPrimeParallel([]int{2, 3, 4, 5, 10, 13, 15, 17})
	for num, prime := range primes {
		fmt.Printf("%d: %v\n", num, prime)
	}

	// Test Merge
	fmt.Println("\n=== Merge ===")
	ch1 := make(chan int)
	ch2 := make(chan int)
	merged := Merge(ch1, ch2)

	go func() {
		for _, v := range []int{1, 3, 5} {
			ch1 <- v
		}
		close(ch1)
	}()
	go func() {
		for _, v := range []int{2, 4, 6} {
			ch2 <- v
		}
		close(ch2)
	}()

	fmt.Print("Merged: ")
	for v := range merged {
		fmt.Print(v, " ")
	}
	fmt.Println() // 1 2 3 4 5 6 (interleaved, order may vary)

	// Run test cases
	allPassed := true

	// === Squarer tests ===
	sq := Squarer([]int{0, -2, 3})
	sqMap := make(map[int]bool)
	for _, v := range sq {
		sqMap[v] = true
	}
	if !sqMap[0] || !sqMap[4] || !sqMap[9] || len(sq) != 3 {
		fmt.Println("FAIL: Squarer with 0 and negative")
		allPassed = false
	}
	// empty
	if len(Squarer([]int{})) != 0 {
		fmt.Println("FAIL: Squarer empty")
		allPassed = false
	}
	// single element
	sq1 := Squarer([]int{7})
	if len(sq1) != 1 || sq1[0] != 49 {
		fmt.Println("FAIL: Squarer single element")
		allPassed = false
	}
	// all zeros
	sqZero := Squarer([]int{0, 0, 0})
	for _, v := range sqZero {
		if v != 0 {
			fmt.Println("FAIL: Squarer all zeros")
			allPassed = false
			break
		}
	}
	// large input (verify no deadlock)
	largeSq := Squarer([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	if len(largeSq) != 10 {
		fmt.Println("FAIL: Squarer large input count")
		allPassed = false
	}

	// === ParallelSum tests ===
	if ParallelSum([][]int{{1, 2, 3}}) != 6 {
		fmt.Println("FAIL: ParallelSum single slice")
		allPassed = false
	}
	if ParallelSum([][]int{}) != 0 {
		fmt.Println("FAIL: ParallelSum empty")
		allPassed = false
	}
	// slices with zeros
	if ParallelSum([][]int{{0, 0}, {0, 0}}) != 0 {
		fmt.Println("FAIL: ParallelSum all zeros")
		allPassed = false
	}
	// negative numbers
	if ParallelSum([][]int{{-1, -2}, {3, 4}}) != 4 {
		fmt.Println("FAIL: ParallelSum negatives")
		allPassed = false
	}
	// many slices
	if ParallelSum([][]int{{1}, {2}, {3}, {4}, {5}}) != 15 {
		fmt.Println("FAIL: ParallelSum many single-element slices")
		allPassed = false
	}
	// empty inner slices
	if ParallelSum([][]int{{}, {1, 2}, {}}) != 3 {
		fmt.Println("FAIL: ParallelSum with empty inner slices")
		allPassed = false
	}

	// === IsPrimeParallel tests ===
	pr := IsPrimeParallel([]int{1, 2, 4, 7})
	if pr[1] != false || pr[2] != true || pr[4] != false || pr[7] != true {
		fmt.Println("FAIL: IsPrimeParallel basic")
		allPassed = false
	}
	// edge cases
	pr2 := IsPrimeParallel([]int{0, 1, 2, 3})
	if pr2[0] != false || pr2[1] != false || pr2[2] != true || pr2[3] != true {
		fmt.Println("FAIL: IsPrimeParallel edge cases 0,1,2,3")
		allPassed = false
	}
	// larger primes
	pr3 := IsPrimeParallel([]int{97, 100, 101})
	if pr3[97] != true || pr3[100] != false || pr3[101] != true {
		fmt.Println("FAIL: IsPrimeParallel larger numbers")
		allPassed = false
	}
	// single element
	pr4 := IsPrimeParallel([]int{13})
	if pr4[13] != true {
		fmt.Println("FAIL: IsPrimeParallel single prime")
		allPassed = false
	}
	// empty input
	pr5 := IsPrimeParallel([]int{})
	if len(pr5) != 0 {
		fmt.Println("FAIL: IsPrimeParallel empty")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: sync package is imported for your use in Merge
var _ = sync.WaitGroup{}
