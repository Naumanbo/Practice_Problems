package main

import "fmt"

// Tests: Channel pipelines, range over channels, close(), producer-consumer pattern
//
// Build a data processing pipeline where each stage is a goroutine
// connected by channels.
//
// Pipeline: Generator -> Double -> Filter (keep evens) -> Collect results
//
// 1. Implement Generator(nums ...int) <-chan int
//    - Returns a receive-only channel
//    - Sends each number into the channel, then closes it
//
// 2. Implement Double(in <-chan int) <-chan int
//    - Reads from in, doubles each value, sends to output channel
//    - Closes output when input is exhausted
//
// 3. Implement FilterEven(in <-chan int) <-chan int
//    - Reads from in, only forwards even numbers
//    - Closes output when input is exhausted
//
// 4. Implement Sum(in <-chan int) int
//    - Reads all values from channel and returns their sum
//
// 5. Implement FanOut(in <-chan int, n int) []<-chan int
//    - Distribute values from in across n output channels (round-robin)
//    - Close all output channels when input is exhausted
//
// 6. Implement FanIn(channels ...<-chan int) <-chan int
//    - Merge multiple channels into one output channel
//    - Close output when ALL inputs are exhausted

// TODO: Implement Generator
func Generator(nums ...int) <-chan int {
	return nil
}

// TODO: Implement Double
func Double(in <-chan int) <-chan int {
	return nil
}

// TODO: Implement FilterEven
func FilterEven(in <-chan int) <-chan int {
	return nil
}

// TODO: Implement Sum
func Sum(in <-chan int) int {
	return 0
}

// TODO: Implement FanOut
func FanOut(in <-chan int, n int) []<-chan int {
	return nil
}

// TODO: Implement FanIn
func FanIn(channels ...<-chan int) <-chan int {
	return nil
}

func main() {
	// Test individual stages
	fmt.Println("=== Generator ===")
	gen := Generator(1, 2, 3, 4, 5)
	for v := range gen {
		fmt.Print(v, " ")
	}
	fmt.Println() // 1 2 3 4 5

	fmt.Println("\n=== Double ===")
	doubled := Double(Generator(1, 2, 3))
	for v := range doubled {
		fmt.Print(v, " ")
	}
	fmt.Println() // 2 4 6

	fmt.Println("\n=== FilterEven ===")
	evens := FilterEven(Generator(1, 2, 3, 4, 5, 6))
	for v := range evens {
		fmt.Print(v, " ")
	}
	fmt.Println() // 2 4 6

	fmt.Println("\n=== Sum ===")
	total := Sum(Generator(1, 2, 3, 4, 5))
	fmt.Println("Sum:", total) // 15

	// Test full pipeline: generate -> double -> filter evens -> sum
	fmt.Println("\n=== Full Pipeline ===")
	// 1,2,3,4,5 -> double -> 2,4,6,8,10 -> filter even -> 2,4,6,8,10 -> sum -> 30
	result := Sum(FilterEven(Double(Generator(1, 2, 3, 4, 5))))
	fmt.Println("Pipeline result:", result) // 30

	// Another pipeline: 1,3,5,7 -> double -> 2,6,10,14 -> filter even -> 2,6,10,14 -> sum -> 32
	result2 := Sum(FilterEven(Double(Generator(1, 3, 5, 7))))
	fmt.Println("Pipeline result 2:", result2) // 32

	// Test FanOut + FanIn
	fmt.Println("\n=== FanOut + FanIn ===")
	source := Generator(1, 2, 3, 4, 5, 6)
	outs := FanOut(source, 3)
	merged := FanIn(outs...)
	total = Sum(merged)
	fmt.Println("FanOut->FanIn sum:", total) // 21

	// Run test cases
	allPassed := true

	// === Generator tests ===
	count := 0
	for range Generator() {
		count++
	}
	if count != 0 {
		fmt.Println("FAIL: Generator empty")
		allPassed = false
	}
	// single value
	genCount := 0
	genVal := 0
	for v := range Generator(42) {
		genVal = v
		genCount++
	}
	if genCount != 1 || genVal != 42 {
		fmt.Println("FAIL: Generator single value")
		allPassed = false
	}
	// channel closes after all values sent
	gen2 := Generator(1, 2, 3)
	collected := []int{}
	for v := range gen2 {
		collected = append(collected, v)
	}
	if len(collected) != 3 || collected[0] != 1 || collected[2] != 3 {
		fmt.Println("FAIL: Generator values and order")
		allPassed = false
	}

	// === Double tests ===
	d := Double(Generator(5))
	val, ok := <-d
	if !ok || val != 10 {
		fmt.Println("FAIL: Double single value")
		allPassed = false
	}
	// double zero
	dz := Double(Generator(0))
	if v := <-dz; v != 0 {
		fmt.Println("FAIL: Double zero")
		allPassed = false
	}
	// double negative
	dn := Double(Generator(-3))
	if v := <-dn; v != -6 {
		fmt.Println("FAIL: Double negative")
		allPassed = false
	}
	// double empty
	de := Double(Generator())
	if _, open := <-de; open {
		fmt.Println("FAIL: Double empty should close channel")
		allPassed = false
	}

	// === FilterEven tests ===
	count = 0
	for range FilterEven(Generator(1, 3, 5, 7)) {
		count++
	}
	if count != 0 {
		fmt.Println("FAIL: FilterEven all odd")
		allPassed = false
	}
	// all even
	evenCount := 0
	for range FilterEven(Generator(2, 4, 6)) {
		evenCount++
	}
	if evenCount != 3 {
		fmt.Println("FAIL: FilterEven all even")
		allPassed = false
	}
	// zero is even
	feZero := FilterEven(Generator(0))
	if v, open := <-feZero; !open || v != 0 {
		fmt.Println("FAIL: FilterEven zero is even")
		allPassed = false
	}
	// negative evens
	feNeg := FilterEven(Generator(-2, -3, -4))
	negEvens := []int{}
	for v := range feNeg {
		negEvens = append(negEvens, v)
	}
	if len(negEvens) != 2 {
		fmt.Println("FAIL: FilterEven negative evens")
		allPassed = false
	}
	// empty
	feEmpty := FilterEven(Generator())
	if _, open := <-feEmpty; open {
		fmt.Println("FAIL: FilterEven empty should close")
		allPassed = false
	}

	// === Sum tests ===
	if Sum(Generator()) != 0 {
		fmt.Println("FAIL: Sum empty")
		allPassed = false
	}
	if Sum(Generator(5)) != 5 {
		fmt.Println("FAIL: Sum single value")
		allPassed = false
	}
	if Sum(Generator(-1, -2, -3)) != -6 {
		fmt.Println("FAIL: Sum negatives")
		allPassed = false
	}

	// === Pipeline correctness ===
	// 1,2,3 -> double -> 2,4,6 -> filter even -> 2,4,6 -> sum -> 12
	if Sum(FilterEven(Double(Generator(1, 2, 3)))) != 12 {
		fmt.Println("FAIL: Pipeline 1,2,3")
		allPassed = false
	}
	// empty pipeline
	if Sum(FilterEven(Double(Generator()))) != 0 {
		fmt.Println("FAIL: Pipeline empty")
		allPassed = false
	}
	// single value pipeline: 5 -> double -> 10 -> filter even -> 10 -> sum -> 10
	if Sum(FilterEven(Double(Generator(5)))) != 10 {
		fmt.Println("FAIL: Pipeline single value")
		allPassed = false
	}

	// === FanOut + FanIn tests ===
	// verify total sum is preserved after fan-out and fan-in
	fo2 := Generator(1, 2, 3, 4)
	outs2 := FanOut(fo2, 2)
	if len(outs2) != 2 {
		fmt.Println("FAIL: FanOut should return 2 channels")
		allPassed = false
	}
	merged2 := FanIn(outs2...)
	total2 := Sum(merged2)
	if total2 != 10 {
		fmt.Println("FAIL: FanOut->FanIn sum should be 10")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
