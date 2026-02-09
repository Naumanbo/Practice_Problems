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
// var ErrEmpty = ...
// var ErrNegative = ...

// TODO: Implement Error() for ValidationError
type ValidationError struct {
	Field   string
	Message string
}

// TODO: Implement Error() for NotFoundError
type NotFoundError struct {
	Resource string
	ID       string
}

// TODO: Implement Error() for NotFoundError

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

	// Sentinel error tests
	if ValidateUser("", 25) != ErrEmpty {
		fmt.Println("FAIL: empty name should return ErrEmpty")
		allPassed = false
	}
	if ValidateUser("Alice", -1) != ErrNegative {
		fmt.Println("FAIL: negative age should return ErrNegative")
		allPassed = false
	}

	// Can use errors.Is with sentinel errors
	if !errors.Is(ValidateUser("", 25), ErrEmpty) {
		fmt.Println("FAIL: errors.Is should work with ErrEmpty")
		allPassed = false
	}

	// ValidationError type assertion
	err := ValidateUser("Alice", 200)
	if _, ok := err.(ValidationError); !ok {
		fmt.Println("FAIL: age > 150 should return ValidationError")
		allPassed = false
	}

	// Valid user
	if ValidateUser("Alice", 25) != nil {
		fmt.Println("FAIL: valid user should return nil")
		allPassed = false
	}

	// Divide by zero
	_, err = Divide(5, 0)
	if err == nil {
		fmt.Println("FAIL: divide by zero should error")
		allPassed = false
	}

	// Divide success
	result, err := Divide(10, 4)
	if err != nil || result != 2.5 {
		fmt.Println("FAIL: 10/4 should be 2.5")
		allPassed = false
	}

	// LookupUser success
	name, err := LookupUser("2")
	if err != nil || name != "Bob" {
		fmt.Println("FAIL: LookupUser(2) should return Bob")
		allPassed = false
	}

	// LookupUser not found - type assertion
	_, err = LookupUser("999")
	if nfErr, ok := err.(NotFoundError); !ok {
		fmt.Println("FAIL: unknown user should return NotFoundError")
		allPassed = false
	} else if nfErr.Resource != "user" || nfErr.ID != "999" {
		fmt.Println("FAIL: NotFoundError fields incorrect")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}

// Note: errors package imported for your use
var _ = errors.New
