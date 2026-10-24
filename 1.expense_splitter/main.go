// Problem: Your friends went for dinner. Build a CLI tool that takes total bill, number of people, tip % and tells each person how much to pay. Handle bad inputs.

// You will learn: package main, fmt.Scan, variables, if/else, go run, go mod init

// Input: 1000 4 10 (bill, people, tip %)
// Output: Each pays: 275.00

// Bonus: If tip is >20%, print "Wow, generous!"

package main

import "fmt"

func main() {
	var bill, people, tip int

	fmt.Println("Type total bill, number of people, and tip percentage:")
	_, err := fmt.Scan(&bill, &people, &tip)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if tip > 20 {
		fmt.Println("Wow, genrous!")
	}

	totalBill := float64(bill) + (float64(bill) * float64(tip) / 100)

	perPerson := totalBill / float64(people)

	fmt.Printf("Each person pays: %.2f\n", perPerson)

}
