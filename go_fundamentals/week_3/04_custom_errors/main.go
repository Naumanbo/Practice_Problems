// Key Takeaways:
// 1. Go has three layers of errors: sentinel errors (errors.New), custom error types
//    (structs implementing Error() string), and wrapped errors (fmt.Errorf with %w).
//    Use sentinels for simple fixed conditions, custom types when errors carry context,
//    and wrapping to add context while preserving the original error.
//
// 2. Sentinel errors are compared with == or errors.Is(). Custom error types are
//    checked with type assertions err.(MyError) or errors.As(). The error interface
//    only requires one method: Error() string.
//
// 3. Validation order matters — when multiple checks can fail, validate in priority
//    order so the most important error is returned first (e.g., empty name before
//    invalid age).
//
// 4. fmt.Errorf("context: %w", err) wraps an error, preserving it for errors.Is()
//    unwrapping. This is different from %v which formats it as a string and loses
//    the original error identity. Always use %w when callers need to inspect the cause.

package main

import (
	"errors"
	"fmt"
)

// Exercise: Custom Error Types & Error Handling
//
// Build a complete error handling system from scratch.
//
// What to build:
// - Two sentinel errors (package-level variables)
// - Two custom error types that satisfy the error interface
// - Three functions that return different kinds of errors
//
// Requirements:
// - A ValidationError type (stores field name and message)
// - A NotFoundError type (stores resource type and ID)
// - A ValidateUser function (validates name and age)
// - A Divide function (handles division by zero with error wrapping using %w)
// - A LookupUser function (hardcoded user database: "1"->Alice, "2"->Bob, "3"->Charlie)
//
// Read the tests in main() to understand exact function signatures,
// error message formats, and expected behavior.
//
// Tip: Start by creating minimal stubs so the file compiles,
// then implement one function at a time.

// === WRITE YOUR CODE BELOW ===
var ErrEmpty = errors.New("empty")
var ErrNegative = errors.New("negative")
var ErrDivideByZero = errors.New("divide by zero")

type ValidationError struct {
	Field   string
	Message string
}

type NotFoundError struct {
	Resource string
	ID       string
}

func (ve ValidationError) Error() string {

	return fmt.Sprintf("validation error on '%s': %s", ve.Field, ve.Message)

}

func (nfe NotFoundError) Error() string {

	return fmt.Sprintf("%s '%s' not found", nfe.Resource, nfe.ID)
}

func ValidateUser(name string, age int) error {

	if name == "" {
		return ErrEmpty
	}

	if age < 0 {
		return ErrNegative
	}

	if age >= 151 {
		return ValidationError{Field: "age", Message: "too high"}
	}
	return nil
}

func Divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero: %w", ErrDivideByZero) // wrapping error using %w
	} else {
		return float64(a) / float64(b), nil
	}

}

func LookupUser(id string) (string, error) {

	switch id {
	case "1":
		return "Alice", nil
	case "2":
		return "Bob", nil
	case "3":
		return "Charlie", nil
	}
	// if id not 1,2, or 3, return an error
	nfe := NotFoundError{Resource: "user", ID: id}
	return "", nfe

}

// === END YOUR CODE ===

func main() {
	// Test ValidationError
	fmt.Println("=== ValidationError ===")
	ve := ValidationError{Field: "email", Message: "invalid format"}
	fmt.Println(ve.Error())
	// validation error on 'email': invalid format

	// Test NotFoundError
	fmt.Println("\n=== NotFoundError ===")
	nfe := NotFoundError{Resource: "user", ID: "42"}
	fmt.Println(nfe.Error())
	// user '42' not found

	// Test ValidateUser
	fmt.Println("\n=== ValidateUser ===")
	fmt.Printf("ValidateUser(\"\", 25): %v\n", ValidateUser("", 25))
	fmt.Printf("ValidateUser(\"Alice\", -5): %v\n", ValidateUser("Alice", -5))
	fmt.Printf("ValidateUser(\"Alice\", 200): %v\n", ValidateUser("Alice", 200))
	fmt.Printf("ValidateUser(\"Alice\", 25): %v\n", ValidateUser("Alice", 25))

	// Test Divide
	fmt.Println("\n=== Divide ===")
	if result, err := Divide(10, 2); err == nil {
		fmt.Printf("10 / 2 = %.1f\n", result)
	}
	if _, err := Divide(10, 0); err != nil {
		fmt.Printf("10 / 0 error: %v\n", err)
	}

	// Test LookupUser
	fmt.Println("\n=== LookupUser ===")
	if name, err := LookupUser("1"); err == nil {
		fmt.Printf("User 1: %s\n", name)
	}
	if _, err := LookupUser("999"); err != nil {
		fmt.Printf("User 999: %v\n", err)
	}

	// Run test cases
	allPassed := true

	// === ValidateUser tests ===
	// sentinel: empty name
	if ValidateUser("", 25) != ErrEmpty {
		fmt.Println("FAIL: empty name should return ErrEmpty")
		allPassed = false
	}
	// sentinel: negative age
	if ValidateUser("Alice", -1) != ErrNegative {
		fmt.Println("FAIL: negative age should return ErrNegative")
		allPassed = false
	}
	// sentinel: very negative age
	if ValidateUser("Alice", -100) != ErrNegative {
		fmt.Println("FAIL: age -100 should return ErrNegative")
		allPassed = false
	}
	// errors.Is works with sentinel errors
	if !errors.Is(ValidateUser("", 25), ErrEmpty) {
		fmt.Println("FAIL: errors.Is should work with ErrEmpty")
		allPassed = false
	}
	if !errors.Is(ValidateUser("Bob", -5), ErrNegative) {
		fmt.Println("FAIL: errors.Is should work with ErrNegative")
		allPassed = false
	}
	// ValidationError for age > 150
	err := ValidateUser("Alice", 200)
	if _, ok := err.(ValidationError); !ok {
		fmt.Println("FAIL: age > 150 should return ValidationError")
		allPassed = false
	}
	// age exactly 151
	err = ValidateUser("Alice", 151)
	if _, ok := err.(ValidationError); !ok {
		fmt.Println("FAIL: age 151 should return ValidationError")
		allPassed = false
	}
	// boundary: age exactly 150 is valid
	if ValidateUser("Alice", 150) != nil {
		fmt.Println("FAIL: age 150 should be valid")
		allPassed = false
	}
	// boundary: age 0 is valid
	if ValidateUser("Baby", 0) != nil {
		fmt.Println("FAIL: age 0 should be valid")
		allPassed = false
	}
	// valid user
	if ValidateUser("Alice", 25) != nil {
		fmt.Println("FAIL: valid user should return nil")
		allPassed = false
	}
	// priority: empty name AND negative age — ErrEmpty should take priority
	if ValidateUser("", -5) != ErrEmpty {
		fmt.Println("FAIL: empty name should take priority over negative age")
		allPassed = false
	}

	// === Error message format tests ===
	ve2 := ValidationError{Field: "age", Message: "too high"}
	if ve2.Error() != "validation error on 'age': too high" {
		fmt.Printf("FAIL: ValidationError format, got %q\n", ve2.Error())
		allPassed = false
	}
	nfe2 := NotFoundError{Resource: "user", ID: "42"}
	if nfe2.Error() != "user '42' not found" {
		fmt.Printf("FAIL: NotFoundError format, got %q\n", nfe2.Error())
		allPassed = false
	}

	// === Divide tests ===
	_, err = Divide(5, 0)
	if err == nil {
		fmt.Println("FAIL: divide by zero should error")
		allPassed = false
	}
	// divide zero by zero
	_, err = Divide(0, 0)
	if err == nil {
		fmt.Println("FAIL: 0/0 should error")
		allPassed = false
	}
	// divide success
	result, err := Divide(10, 4)
	if err != nil || result != 2.5 {
		fmt.Println("FAIL: 10/4 should be 2.5")
		allPassed = false
	}
	// divide negative
	result, err = Divide(-10, 2)
	if err != nil || result != -5.0 {
		fmt.Println("FAIL: -10/2 should be -5")
		allPassed = false
	}
	// divide by negative
	result, err = Divide(10, -2)
	if err != nil || result != -5.0 {
		fmt.Println("FAIL: 10/-2 should be -5")
		allPassed = false
	}
	// divide zero by something
	result, err = Divide(0, 5)
	if err != nil || result != 0 {
		fmt.Println("FAIL: 0/5 should be 0")
		allPassed = false
	}
	// error wrapping: Divide error should contain %w
	_, divErr := Divide(10, 0)
	if divErr == nil || divErr.Error() == "" {
		fmt.Println("FAIL: Divide error message should not be empty")
		allPassed = false
	}

	// === LookupUser tests ===
	// all known users
	name, err := LookupUser("1")
	if err != nil || name != "Alice" {
		fmt.Println("FAIL: LookupUser(1) should return Alice")
		allPassed = false
	}
	name, err = LookupUser("2")
	if err != nil || name != "Bob" {
		fmt.Println("FAIL: LookupUser(2) should return Bob")
		allPassed = false
	}
	name, err = LookupUser("3")
	if err != nil || name != "Charlie" {
		fmt.Println("FAIL: LookupUser(3) should return Charlie")
		allPassed = false
	}
	// not found - type assertion
	_, err = LookupUser("999")
	if nfErr, ok := err.(NotFoundError); !ok {
		fmt.Println("FAIL: unknown user should return NotFoundError")
		allPassed = false
	} else if nfErr.Resource != "user" || nfErr.ID != "999" {
		fmt.Println("FAIL: NotFoundError fields incorrect")
		allPassed = false
	}
	// not found - empty ID
	_, err = LookupUser("")
	if _, ok := err.(NotFoundError); !ok {
		fmt.Println("FAIL: empty ID should return NotFoundError")
		allPassed = false
	}
	// not found - ID "0"
	_, err = LookupUser("0")
	if _, ok := err.(NotFoundError); !ok {
		fmt.Println("FAIL: ID 0 should return NotFoundError")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: errors package imported for your use
var _ = errors.New
