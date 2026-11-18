// Problem: Take your Python User class with **init**, get_email, update_email and rewrite it in Go with struct + methods. Make email private (lowercase).

// Learn: Struct, Methods, *User pointer receiver vs User value receiver

// Test:

// Go
// u := User{Name: "Sanket"}
// u.UpdateEmail("new@email.com")
// fmt.Println(u.Email()) // should use getter, not direct access

package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Name  string
	email string
}

func (u *User) Email() string {
	return u.email
}

func (u *User) UpdateEmail(newEmail string) error {
	if !strings.Contains(newEmail, "@") {
		return errors.New("Invalid Email Address")
	}
	u.email = newEmail
	return nil
}

func main() {
	u := &User{Name: "Sanket"}
	err := u.UpdateEmail("sanket.kumbhare@gmail.com")
	if err != nil {
		fmt.Printf("Err: %s", err)
	}
	fmt.Println(u.Email())
}
