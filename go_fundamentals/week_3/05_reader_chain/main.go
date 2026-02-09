package main

import (
	"io"
	"strings"
	"unicode"
)

// Tests: io.Reader interface, reader chaining/wrapping, byte transformation
//
// Create reader wrappers that transform data as it passes through.
// Each reader wraps another io.Reader and modifies the stream.
//
// 1. UpperReader - converts all lowercase letters to uppercase
//    - Only affects a-z, leaves everything else unchanged
//
// 2. FilterReader - removes all characters matching a predicate
//    - Takes an io.Reader and a func(byte) bool
//    - If predicate returns true, the byte is removed from output
//
// 3. LimitReader - reads at most N bytes total
//    - After N bytes, returns io.EOF
//    - Hint: track how many bytes have been read
//
// 4. ChainReader - chains multiple readers together
//    - When first reader returns EOF, continue with next
//    - Like io.MultiReader but implement it yourself
//
// Bonus understanding:
// - Read() fills the provided buffer with data
// - Returns (n, err) where n is bytes written to buffer
// - Return io.EOF when no more data

// TODO: Implement UpperReader
type UpperReader struct {
	r io.Reader
}

func (u *UpperReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

// TODO: Implement FilterReader
type FilterReader struct {
	r         io.Reader
	predicate func(byte) bool
}

func (f *FilterReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

// TODO: Implement LimitReader
type LimitReader struct {
	r io.Reader
	n int // bytes remaining
}

func (l *LimitReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

// TODO: Implement ChainReader
type ChainReader struct {
	readers []io.Reader
	current int
}

func (c *ChainReader) Read(p []byte) (int, error) {
	return 0, io.EOF
}

// Helper function to read all bytes from a reader
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

	// UpperReader tests
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

	// FilterReader tests
	isDigit := func(b byte) bool { return unicode.IsDigit(rune(b)) }
	if readAll(&FilterReader{strings.NewReader("a1b2c3"), isDigit}) != "abc" {
		println("FAIL: FilterReader digits")
		allPassed = false
	}
	if readAll(&FilterReader{strings.NewReader("abc"), isDigit}) != "abc" {
		println("FAIL: FilterReader nothing to filter")
		allPassed = false
	}

	// LimitReader tests
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

	// ChainReader tests
	cr := &ChainReader{readers: []io.Reader{
		strings.NewReader("a"),
		strings.NewReader("b"),
		strings.NewReader("c"),
	}}
	if readAll(cr) != "abc" {
		println("FAIL: ChainReader abc")
		allPassed = false
	}

	// Empty chain
	if readAll(&ChainReader{readers: []io.Reader{}}) != "" {
		println("FAIL: ChainReader empty")
		allPassed = false
	}

	if allPassed {
		println("\nAll tests passed!")
	}
}
