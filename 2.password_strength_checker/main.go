// Problem: Build a tool that checks if a password is strong. Strong = >8 chars, has uppercase, lowercase, number, special char. Don't use regex library, use loops.

// You will learn: for loop (Go has only for), strings, functions, multiple returns

// Bonus: Return exactly why it's weak.

// ----------------------------------
// package main

// import "fmt"

// func checkPassword(pwd string) (bool, []string) {
// 	var isStrong bool
// 	var whyWeak []string

// 	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	lowercase := "abcdefghijklmnopqrstuvwxyz"
// 	number := "0123456789"
// 	specialChar := "!@#$%^&*()"

// 	var hasUppercase, hasLowercase, hasNumber, hasSpecialCharacter bool

// 	if len(pwd) < 8 {
// 		whyWeak = append(whyWeak, "Password length must be greater than or equals to 8")
// 	}

// 	for _, letter := range pwd {
// 		for _, u := range uppercase {
// 			if string(letter) == string(u) {
// 				hasUppercase = true
// 				break
// 			}
// 		}
// 		for _, l := range lowercase {
// 			if string(letter) == string(l) {
// 				hasLowercase = true
// 				break
// 			}
// 		}
// 		for _, n := range number {
// 			if string(letter) == string(n) {
// 				hasNumber = true
// 				break
// 			}
// 		}
// 		for _, s := range specialChar {
// 			if string(letter) == string(s) {
// 				hasSpecialCharacter = true
// 				break
// 			}
// 		}
// 	}

// 	if !hasUppercase {
// 		whyWeak = append(whyWeak, "Password must need one Uppercase")
// 	}

// 	if !hasLowercase {
// 		whyWeak = append(whyWeak, "Password must need one Lowercase")
// 	}

// 	if !hasNumber {
// 		whyWeak = append(whyWeak, "Password must need one Number")
// 	}
// 	if !hasSpecialCharacter {
// 		whyWeak = append(whyWeak, "Password must need one Special Character")
// 	}

// 	if len(whyWeak) == 0 {
// 		isStrong = true
// 	}

//		return isStrong, whyWeak
//	}

// -------------------------------

package main

import "fmt"

func main() {
	var pwd string

	fmt.Println("Enter your password")

	_, err := fmt.Scan(&pwd)

	if err != nil {
		fmt.Println("Err: ", err)
	}

	isStrong, whyWeak := checkPassword(pwd)

	fmt.Println("isStrong: ", isStrong)
	fmt.Println("whyWeak: ", whyWeak)
}

func checkPassword(pwd string) (bool, []string) {
	var whyWeak []string
	var hasUpper, hasLower, hasNumber, hasSpecial bool

	if len(pwd) < 8 {
		whyWeak = append(whyWeak, "Password must be at least 8 characters")
	}

	for _, ch := range pwd {
		switch {
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= '0' && ch <= '9':
			hasNumber = true
		case ch == '!' || ch == '@' || ch == '#' || ch == '$' || ch == '%' || ch == '^' || ch == '&' || ch == '*' || ch == '(' || ch == ')':
			hasSpecial = true
		}
	}

	if !hasUpper {
		whyWeak = append(whyWeak, "Needs at least one Uppercase")
	}
	if !hasLower {
		whyWeak = append(whyWeak, "Needs at least one Lowercase")
	}
	if !hasNumber {
		whyWeak = append(whyWeak, "Needs at least one Number")
	}
	if !hasSpecial {
		whyWeak = append(whyWeak, "Needs at least one Special Character (!@#$%^&*())")
	}

	isStrong := len(whyWeak) == 0
	return isStrong, whyWeak
}
