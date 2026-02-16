package main

import (
	"errors"
	"fmt"
)

// Tests: Error interface, custom error types, error wrapping, sentinel errors
//
// The error interface is: type error interface { Error() string }
//
// 1. Create a ValidationError type with fields:
//    - Field string (which field failed)
//    - Message string (what went wrong)
//    Error() should return: "validation error on 'Field': Message"
//
// 2. Create a NotFoundError type with fields:
//    - Resource string (e.g., "user", "file")
//    - ID string (the identifier that wasn't found)
//    Error() should return: "Resource 'ID' not found"
//
// 3. Create sentinel errors (package-level var):
//    - ErrEmpty = errors.New("empty input")
//    - ErrNegative = errors.New("negative number not allowed")
//
// 4. Implement ValidateUser(name string, age int) error
//    - Return ErrEmpty if name is empty
//    - Return ErrNegative if age < 0
//    - Return ValidationError if age > 150
//    - Return nil if valid
//
// 5. Implement Divide(a, b float64) (float64, error)
//    - Return error if b is 0
//    - Wrap the error with context: fmt.Errorf("divide %v by %v: %w", a, b, err)
//
// 6. Implement LookupUser(id string) (string, error)
//    - Known IDs: "1" -> "Alice", "2" -> "Bob", "3" -> "Charlie"
//    - Return NotFoundError for unknown IDs

// TODO: Define sentinel errors
var ErrEmpty = errors.New("empty input")
var ErrNegative = errors.New("negative number not allowed")

// TODO: Implement Error() for ValidationError
type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string { return "" } // TODO: implement

// TODO: Implement Error() for NotFoundError
type NotFoundError struct {
	Resource string
	ID       string
}

func (n NotFoundError) Error() string { return "" } // TODO: implement

// TODO: Implement ValidateUser
func ValidateUser(name string, age int) error {
	return nil
}

// TODO: Implement Divide
func Divide(a, b float64) (float64, error) {
	return 0, nil
}

// TODO: Implement LookupUser
func LookupUser(id string) (string, error) {
	return "", nil
}

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
