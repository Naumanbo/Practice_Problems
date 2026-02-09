package main

import "fmt"

// Tests: Struct embedding, method promotion, value vs pointer receiver semantics
//
// Go doesn't have inheritance. It has EMBEDDING, which promotes methods
// from the inner struct to the outer struct. This problem forces you to
// understand how that works and where it breaks.
//
// 1. Define a Base struct with fields: ID int, CreatedAt string
//    - Method: Describe() string -> "ID: <id>, Created: <createdAt>"
//
// 2. Define a User struct that EMBEDS Base
//    - Fields: Name string, Email string
//    - Method: Describe() string -> "User: <name> (<email>), ID: <id>"
//    (This OVERRIDES Base.Describe - understand why)
//
// 3. Define an Admin struct that EMBEDS User
//    - Fields: Permissions []string
//    - Method: HasPermission(perm string) bool
//    - Method: Describe() string -> "Admin: <name>, Permissions: [<perms>]"
//
// 4. Implement Promote(user *User, perms []string) Admin
//    - Creates an Admin from an existing User with given permissions
//    - The Admin should share the same Base data (ID, CreatedAt)
//
// 5. Implement UpdateEmail(user *User, newEmail string)
//    - Pointer receiver question: if Admin embeds User (not *User),
//      does calling admin.UpdateEmail() modify the admin's copy?
//    - Implement and verify with the test cases below

type Base struct {
	ID        int
	CreatedAt string
}

func (b Base) Describe() string {
	return ""
}

type User struct {
	Base
	Name  string
	Email string
}

func (u User) Describe() string {
	return ""
}

// UpdateEmail modifies the user's email
// Think: should this be a value or pointer receiver?
func (u *User) UpdateEmail(newEmail string) {
}

type Admin struct {
	User
	Permissions []string
}

func (a Admin) Describe() string {
	return ""
}

func (a Admin) HasPermission(perm string) bool {
	return false
}

// Promote creates an Admin from a User
func Promote(user *User, perms []string) Admin {
	return Admin{}
}

func main() {
	// Test Base
	fmt.Println("=== Base ===")
	b := Base{ID: 1, CreatedAt: "2026-01-01"}
	fmt.Println(b.Describe()) // ID: 1, Created: 2026-01-01

	// Test User with embedding
	fmt.Println("\n=== User ===")
	u := User{
		Base:  Base{ID: 42, CreatedAt: "2026-02-01"},
		Name:  "Alice",
		Email: "alice@example.com",
	}
	fmt.Println(u.Describe())   // User: Alice (alice@example.com), ID: 42
	fmt.Println(u.ID)           // 42 - accessed directly through embedding
	fmt.Println(u.CreatedAt)    // 2026-02-01 - promoted field

	// Test overridden Describe - can still access base version
	fmt.Println(u.Base.Describe()) // ID: 42, Created: 2026-02-01

	// Test Admin
	fmt.Println("\n=== Admin ===")
	admin := Promote(&u, []string{"read", "write", "delete"})
	fmt.Println(admin.Describe()) // Admin: Alice, Permissions: [read write delete]
	fmt.Println("Has 'write':", admin.HasPermission("write"))   // true
	fmt.Println("Has 'deploy':", admin.HasPermission("deploy")) // false

	// Test the pointer receiver trap with embedding
	fmt.Println("\n=== Pointer Receiver + Embedding ===")
	admin.UpdateEmail("admin@example.com")
	fmt.Println("Admin email after update:", admin.Email)
	fmt.Println("Original user email:", u.Email)
	// Think: are these the same or different? Why?

	// Run test cases
	allPassed := true

	// Base describe
	if Base{ID: 5, CreatedAt: "today"}.Describe() != "ID: 5, Created: today" {
		fmt.Println("FAIL: Base.Describe")
		allPassed = false
	}

	// User describe includes ID from embedded Base
	u2 := User{Base: Base{ID: 10, CreatedAt: "now"}, Name: "Bob", Email: "bob@test.com"}
	if u2.Describe() != "User: Bob (bob@test.com), ID: 10" {
		fmt.Printf("FAIL: User.Describe, got: %q\n", u2.Describe())
		allPassed = false
	}

	// Promoted field access
	if u2.ID != 10 {
		fmt.Println("FAIL: Promoted field ID")
		allPassed = false
	}

	// Admin HasPermission
	a := Admin{User: u2, Permissions: []string{"read", "write"}}
	if !a.HasPermission("read") {
		fmt.Println("FAIL: HasPermission 'read'")
		allPassed = false
	}
	if a.HasPermission("delete") {
		fmt.Println("FAIL: HasPermission 'delete' should be false")
		allPassed = false
	}
	if a.HasPermission("") {
		fmt.Println("FAIL: HasPermission empty string")
		allPassed = false
	}

	// Empty permissions
	a2 := Admin{User: u2, Permissions: []string{}}
	if a2.HasPermission("anything") {
		fmt.Println("FAIL: Empty permissions should return false")
		allPassed = false
	}

	// Promote preserves Base data
	u3 := User{Base: Base{ID: 99, CreatedAt: "2026-01-15"}, Name: "Charlie", Email: "c@test.com"}
	promoted := Promote(&u3, []string{"admin"})
	if promoted.ID != 99 || promoted.Name != "Charlie" {
		fmt.Println("FAIL: Promote should preserve Base and User data")
		allPassed = false
	}

	// UpdateEmail with pointer receiver on embedded struct
	u4 := User{Base: Base{ID: 1, CreatedAt: "now"}, Name: "Dave", Email: "old@test.com"}
	u4.UpdateEmail("new@test.com")
	if u4.Email != "new@test.com" {
		fmt.Println("FAIL: UpdateEmail on User")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
