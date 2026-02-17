package main

import (
	"io"
	"strings"
	"unicode"
)

// Exercise: io.Reader Wrappers & Stream Transformation
//
// Build reader types that wrap other readers and transform data as it flows through.
// Each reader must satisfy the io.Reader interface: Read(p []byte) (int, error)
//
// What to build:
// - UpperReader: wraps a reader, converts lowercase letters to uppercase
// - FilterReader: wraps a reader, removes bytes matching a predicate function
// - LimitReader: wraps a reader, reads at most N bytes then returns io.EOF
// - ChainReader: reads from multiple readers sequentially (like io.MultiReader)
//
// Key concept: Read() fills the provided byte slice with data and returns
// (bytesRead, error). Return io.EOF when no more data is available.
//
// Read the tests in main() to understand exact type names, field names,
// and expected behavior. The readAll helper function is provided for testing.

// === WRITE YOUR CODE BELOW ===

// === END YOUR CODE ===

// readAll is a test helper — reads all bytes from a reader (DO NOT MODIFY)
func readAll(r io.Reader) string {
	var result []byte
	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		result = append(result, buf[:n]...)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
	return string(result)
}

func main() {
	// Test UpperReader
	println("=== UpperReader ===")
	upper := &UpperReader{strings.NewReader("Hello, World! 123")}
	println(readAll(upper))
	// Expected: HELLO, WORLD! 123

	// Test FilterReader - remove vowels
	println("\n=== FilterReader ===")
	isVowel := func(b byte) bool {
		return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' ||
			b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
	}
	filter := &FilterReader{strings.NewReader("Hello World"), isVowel}
	println(readAll(filter))
	// Expected: Hll Wrld

	// Test FilterReader - remove spaces
	println("\n=== FilterReader (no spaces) ===")
	isSpace := func(b byte) bool { return b == ' ' }
	filter2 := &FilterReader{strings.NewReader("a b c d e"), isSpace}
	println(readAll(filter2))
	// Expected: abcde

	// Test LimitReader
	println("\n=== LimitReader ===")
	limit := &LimitReader{strings.NewReader("Hello, World!"), 5}
	println(readAll(limit))
	// Expected: Hello

	// Test ChainReader
	println("\n=== ChainReader ===")
	chain := &ChainReader{
		readers: []io.Reader{
			strings.NewReader("Hello"),
			strings.NewReader(" "),
			strings.NewReader("World"),
		},
	}
	println(readAll(chain))
	// Expected: Hello World

	// Test chaining readers together
	println("\n=== Chained Transformations ===")
	// Uppercase then limit to 5 characters
	src := strings.NewReader("hello world")
	chained := &LimitReader{&UpperReader{src}, 5}
	println(readAll(chained))
	// Expected: HELLO

	// Run test cases
	allPassed := true

	// === UpperReader tests ===
	if readAll(&UpperReader{strings.NewReader("abc")}) != "ABC" {
		println("FAIL: UpperReader abc")
		allPassed = false
	}
	if readAll(&UpperReader{strings.NewReader("ABC123")}) != "ABC123" {
		println("FAIL: UpperReader already upper")
		allPassed = false
	}
	if readAll(&UpperReader{strings.NewReader("")}) != "" {
		println("FAIL: UpperReader empty")
		allPassed = false
	}
	// mixed case with punctuation
	if readAll(&UpperReader{strings.NewReader("Hello, World!")}) != "HELLO, WORLD!" {
		println("FAIL: UpperReader mixed case")
		allPassed = false
	}
	// only symbols and numbers (no change)
	if readAll(&UpperReader{strings.NewReader("123!@#")}) != "123!@#" {
		println("FAIL: UpperReader no letters")
		allPassed = false
	}
	// single character
	if readAll(&UpperReader{strings.NewReader("z")}) != "Z" {
		println("FAIL: UpperReader single char")
		allPassed = false
	}
	// spaces and tabs
	if readAll(&UpperReader{strings.NewReader("a b\tc")}) != "A B\tC" {
		println("FAIL: UpperReader whitespace preserved")
		allPassed = false
	}

	// === FilterReader tests ===
	isDigit := func(b byte) bool { return unicode.IsDigit(rune(b)) }
	if readAll(&FilterReader{strings.NewReader("a1b2c3"), isDigit}) != "abc" {
		println("FAIL: FilterReader digits")
		allPassed = false
	}
	if readAll(&FilterReader{strings.NewReader("abc"), isDigit}) != "abc" {
		println("FAIL: FilterReader nothing to filter")
		allPassed = false
	}
	// filter everything
	allMatch := func(b byte) bool { return true }
	if readAll(&FilterReader{strings.NewReader("hello"), allMatch}) != "" {
		println("FAIL: FilterReader remove all")
		allPassed = false
	}
	// filter nothing
	noMatch := func(b byte) bool { return false }
	if readAll(&FilterReader{strings.NewReader("hello"), noMatch}) != "hello" {
		println("FAIL: FilterReader remove none")
		allPassed = false
	}
	// empty input
	if readAll(&FilterReader{strings.NewReader(""), isDigit}) != "" {
		println("FAIL: FilterReader empty input")
		allPassed = false
	}
	// filter spaces
	isSpaceFn := func(b byte) bool { return b == ' ' }
	if readAll(&FilterReader{strings.NewReader("a b c"), isSpaceFn}) != "abc" {
		println("FAIL: FilterReader spaces")
		allPassed = false
	}

	// === LimitReader tests ===
	if readAll(&LimitReader{strings.NewReader("hello"), 3}) != "hel" {
		println("FAIL: LimitReader 3")
		allPassed = false
	}
	if readAll(&LimitReader{strings.NewReader("hi"), 10}) != "hi" {
		println("FAIL: LimitReader longer than content")
		allPassed = false
	}
	if readAll(&LimitReader{strings.NewReader("hello"), 0}) != "" {
		println("FAIL: LimitReader 0")
		allPassed = false
	}
	// limit exact length
	if readAll(&LimitReader{strings.NewReader("abc"), 3}) != "abc" {
		println("FAIL: LimitReader exact length")
		allPassed = false
	}
	// limit 1
	if readAll(&LimitReader{strings.NewReader("hello"), 1}) != "h" {
		println("FAIL: LimitReader 1")
		allPassed = false
	}
	// empty source
	if readAll(&LimitReader{strings.NewReader(""), 5}) != "" {
		println("FAIL: LimitReader empty source")
		allPassed = false
	}

	// === ChainReader tests ===
	cr := &ChainReader{readers: []io.Reader{
		strings.NewReader("a"),
		strings.NewReader("b"),
		strings.NewReader("c"),
	}}
	if readAll(cr) != "abc" {
		println("FAIL: ChainReader abc")
		allPassed = false
	}
	// empty chain
	if readAll(&ChainReader{readers: []io.Reader{}}) != "" {
		println("FAIL: ChainReader empty")
		allPassed = false
	}
	// single reader
	cr2 := &ChainReader{readers: []io.Reader{strings.NewReader("only")}}
	if readAll(cr2) != "only" {
		println("FAIL: ChainReader single reader")
		allPassed = false
	}
	// chain with empty readers mixed in
	cr3 := &ChainReader{readers: []io.Reader{
		strings.NewReader(""),
		strings.NewReader("hello"),
		strings.NewReader(""),
		strings.NewReader("world"),
		strings.NewReader(""),
	}}
	if readAll(cr3) != "helloworld" {
		println("FAIL: ChainReader with empty readers")
		allPassed = false
	}

	// === Chained transformations ===
	// UpperReader + FilterReader
	srcChain := strings.NewReader("Hello World 123")
	chained2 := &FilterReader{&UpperReader{srcChain}, isDigit}
	if readAll(chained2) != "HELLO WORLD " {
		println("FAIL: Upper then Filter digits")
		allPassed = false
	}
	// LimitReader on ChainReader
	cr4 := &ChainReader{readers: []io.Reader{
		strings.NewReader("abc"),
		strings.NewReader("def"),
	}}
	limited := &LimitReader{cr4, 4}
	if readAll(limited) != "abcd" {
		println("FAIL: LimitReader on ChainReader")
		allPassed = false
	}

	if allPassed {
		println("\nAll tests passed!")
	}
}
