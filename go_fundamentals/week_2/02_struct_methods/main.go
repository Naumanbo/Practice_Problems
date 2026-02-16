package main

import "fmt"

// Tests: Struct embedding, method promotion, value vs pointer receiver semantics
//
// KEY TAKEAWAYS:
// - Go has no inheritance. EMBEDDING promotes fields/methods from inner to outer struct.
// - Embedding `User` (value) = Admin owns a COPY. Mutating admin's User does NOT affect the original.
// - Embedding `*User` (pointer) = Admin SHARES the original. Mutations propagate to both.
// - The receiver type (*User vs User) controls whether a method can mutate its own struct,
//   but does NOT control whether two variables point to the same data — that's the embedding type.
// - Overriding: User.Describe() shadows Base.Describe(), but you can still call u.Base.Describe().
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
	return fmt.Sprintf(("ID: %d, Created: %s"), b.ID, b.CreatedAt)
}

type User struct {
	Base  // embedding, done by adding another struct without assigning variable to it
	Name  string
	Email string
}

func (u User) Describe() string {
	return fmt.Sprintf("User: %s (%s), ID: %d", u.Name, u.Email, u.ID)
}

// 5. Implement UpdateEmail(user *User, newEmail string)
//   - Pointer receiver question: if Admin embeds User (not *User),
//     does calling admin.UpdateEmail() modify the admin's copy?
//   - Implement and verify with the test cases below
//
// UpdateEmail modifies the user's email
// Think: should this be a value or pointer receiver?
func (u *User) UpdateEmail(newEmail string) {
	u.Email = newEmail
}

// 3. Define an Admin struct that EMBEDS User
//   - Fields: Permissions []string
//   - Method: HasPermission(perm string) bool
//   - Method: Describe() string -> "Admin: <name>, Permissions: [<perms>]"
type Admin struct {
	User        // embeds a COPY of user (value embedding)
	Permissions []string
}

func (a Admin) Describe() string {

	return fmt.Sprintf("Admin: %s, Permissions: [%s]", a.Name, a.Permissions)
}

func (a Admin) HasPermission(perm string) bool {

	for _, v := range a.Permissions {
		if v == perm {
			return true
		}
	}

	return false
}

// 4. Implement Promote(user *User, perms []string) Admin
//    - Creates an Admin from an existing User with given permissions
//    - The Admin should share the same Base data (ID, CreatedAt)
//

// Promote creates an Admin from a User
func Promote(user *User, perms []string) Admin {
	return Admin{*user, perms}
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
	fmt.Println(u.Describe()) // User: Alice (alice@example.com), ID: 42
	fmt.Println(u.ID)         // 42 - accessed directly through embedding
	fmt.Println(u.CreatedAt)  // 2026-02-01 - promoted field

	// Test overridden Describe - can still access base version
	fmt.Println(u.Base.Describe()) // ID: 42, Created: 2026-02-01

	// Test Admin
	fmt.Println("\n=== Admin ===")
	admin := Promote(&u, []string{"read", "write", "delete"})
	fmt.Println(admin.Describe())                               // Admin: Alice, Permissions: [read write delete]
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
	b2 := Base{ID: 5, CreatedAt: "today"}
	if b2.Describe() != "ID: 5, Created: today" {
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

	// === Additional edge case tests ===

	// Value embedding: Admin gets a COPY, not shared reference
	u5 := User{Base: Base{ID: 50, CreatedAt: "2026-01-01"}, Name: "Eve", Email: "eve@test.com"}
	admin2 := Promote(&u5, []string{"read"})
	admin2.UpdateEmail("admin@test.com")
	if u5.Email != "eve@test.com" {
		fmt.Println("FAIL: Value embedding - original user should not be modified")
		allPassed = false
	}
	if admin2.Email != "admin@test.com" {
		fmt.Println("FAIL: Value embedding - admin copy should be modified")
		allPassed = false
	}

	// Base.Describe accessible through Admin (two levels of embedding)
	admin3 := Admin{User: User{Base: Base{ID: 7, CreatedAt: "2026-03-01"}, Name: "Frank", Email: "f@test.com"}, Permissions: []string{}}
	if admin3.Base.Describe() != "ID: 7, Created: 2026-03-01" {
		fmt.Println("FAIL: Admin should access Base.Describe through embedding chain")
		allPassed = false
	}

	// Promoted field access through Admin
	if admin3.ID != 7 || admin3.CreatedAt != "2026-03-01" {
		fmt.Println("FAIL: Admin promoted fields from Base")
		allPassed = false
	}
	if admin3.Name != "Frank" || admin3.Email != "f@test.com" {
		fmt.Println("FAIL: Admin promoted fields from User")
		allPassed = false
	}

	// HasPermission with single permission
	admin4 := Admin{User: u2, Permissions: []string{"only"}}
	if !admin4.HasPermission("only") {
		fmt.Println("FAIL: HasPermission single permission present")
		allPassed = false
	}
	if admin4.HasPermission("other") {
		fmt.Println("FAIL: HasPermission single permission absent")
		allPassed = false
	}

	// HasPermission case sensitivity
	admin5 := Admin{User: u2, Permissions: []string{"Read", "Write"}}
	if admin5.HasPermission("read") {
		fmt.Println("FAIL: HasPermission should be case-sensitive")
		allPassed = false
	}

	// Promote with nil-like empty permissions
	promoted2 := Promote(&u2, []string{})
	if len(promoted2.Permissions) != 0 {
		fmt.Println("FAIL: Promote with empty permissions")
		allPassed = false
	}

	// User.Describe overrides Base.Describe
	u6 := User{Base: Base{ID: 1, CreatedAt: "now"}, Name: "Test", Email: "t@t.com"}
	if u6.Describe() == u6.Base.Describe() {
		fmt.Println("FAIL: User.Describe should differ from Base.Describe")
		allPassed = false
	}

	if allPassed {
		fmt.Println("\nAll tests passed!")
	}
}
