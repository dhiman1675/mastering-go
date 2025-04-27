package user

import (
	"errors"
	"fmt"
	"time"
)

// In Golang to export a struct, method or anything we need to start it with a uppercase character
type User struct {
	firstName string
	lastName  string
	email     string
	age       int
	createdAt time.Time
}

// struct embedding
type Admin struct {
	password string
	User
}

// Constructor function.
// This is just a convention, not a built-in function.
func New(firstName string, lastName string, email string, age int) (*User, error) {
	if firstName == "" || lastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	// Returning a pointer to the User struct.
	return &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		age:       age,
		createdAt: time.Now(),
	}, nil

	// Note: You can also omit the field names if you follow the same order as the struct definition.
	// Example:
	// appUser := User{
	// 	"R",
	// 	"D",
	// 	"r@d.com",
	// 	18,
	// 	time.Now(),
	// }
}

func NewAdmin(password string) Admin {
	return Admin{
		password: password,
		User: User{
			firstName: "Admin",
		},
	}
}

// By using 'func (u User)', we are attaching a method to a struct.
// A function attached to a struct in Go is called a method.
// The '(u User)' part is called the receiver.
func (u User) OutputUserData() {
	fmt.Println(u)

	// Note: When working with a struct pointer, we can access its fields directly (like u.firstName).
	// Unlike basic types (e.g., *int), we don't need to manually dereference it using *u.
	// Go provides automatic dereferencing for struct pointers to simplify access.
	fmt.Println("=>", u.firstName, u.createdAt)
}

// We need to use a pointer receiver here because we want to update the original struct values.
// Otherwise, Go would only update a copy of the struct, not the original.
func (u *User) ChangeUserName() {
	u.firstName = "C"
	u.lastName = "J"
}
