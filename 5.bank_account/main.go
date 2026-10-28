// Problem: Create a bank account that doesn't allow negative balance.

// Create:

// type Account struct { Name string; balance int } (balance is private - lowercase)
// Methods: Deposit(amount int), Withdraw(amount int) error, Balance() int
// You will learn: Methods, pointer receiver (a *Account) vs value receiver, encapsulation, custom errors with errors.New

// Test case:

// Go
// acc := Account{Name: "Sanket"}
// acc.Deposit(1000)
// err := acc.Withdraw(1500) // should return error

package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Name    string
	balance int
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Withdrawal is more than Deposits")
	}
	a.balance -= amount
	return nil
}

func (a Account) Balance() int {
	return a.balance
}

func main() {
	acc := Account{Name: "Sanket"}

	fmt.Println("Initial Balance", acc.Balance())
	acc.Deposit(1000)
	fmt.Println("Balance after deposit 1000:", acc.Balance())

	err := acc.Withdraw(1500)
	if err != nil {
		fmt.Println("Withdraw Err:", err)
	}
	fmt.Println("Balance after withdraw 1500:", acc.Balance())

	err = acc.Withdraw(500)
	if err != nil {
		fmt.Println("Withdraw Err:", err)
	}
	fmt.Println("Balance after withdraw 500:", acc.Balance())
}
