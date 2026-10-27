// Problem: Build a contact book that runs in terminal and keeps data in contacts.json.

// Commands:
// go run main.go add "Sanket" "9999999999"
// go run main.go list
// go run main.go search Sanket

// You will learn: struct, os.Args, encoding/json, struct tags, slices

// Go
// type Contact struct {
//     Name  string `json:"name"`
//     Phone string `json:"phone"`
// }
// Bonus: delete command.

// package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"os"
// )

// type Contact struct {
// 	Name  string `json:"name"`
// 	Phone string `json:"phone"`
// }

// func sliceTobytes(contacts []Contact) []byte {
// 	bytes, err := json.Marshal(contacts)
// 	if err != nil {
// 		fmt.Println("Err: ", err)
// 	}
// 	return bytes
// }

// func bytesToSlice(data []byte, contacts *[]Contact) {
// 	if len(data) == 0 {
// 		return
// 	}
// 	err := json.Unmarshal(data, contacts)
// 	if err != nil {
// 		fmt.Println("Err: 89", err)
// 	}
// }

// func main() {
// 	args := os.Args
// 	action := args[1]
// 	var userName, phone string

// 	if len(args) > 2 {
// 		userName = args[2]
// 	}

// 	if len(args) > 3 {
// 		phone = args[3]
// 	}

// 	var fileName string
// 	file, err := os.Stat("contact.json")
// 	if file != nil {
// 		fileName = file.Name()
// 	}
// 	if errors.Is(err, os.ErrNotExist) {
// 		f, err := os.Create("contact.json")
// 		if err != nil {
// 			fmt.Println("Err: ", err)
// 		}
// 		fileName = f.Name()
// 		fmt.Println("contact.json created")
// 	}

// 	content, err := os.ReadFile(fileName)
// 	if err != nil {
// 		fmt.Println("Err123: ", err)
// 	}

// 	var contacts []Contact
// 	bytesToSlice(content, &contacts)

// 	switch action {
// 	case "add":
// 		if userName == "" || phone == "" {
// 			fmt.Println("Name and Phone arguments are missing.")
// 			return
// 		}
// 		contact := Contact{
// 			userName,
// 			phone,
// 		}

// 		contacts = append(contacts, contact)
// 		jsonData := sliceTobytes(contacts)

// 		err := os.WriteFile(fileName, jsonData, 0666)
// 		if err != nil {
// 			fmt.Println("Err 67:", err)
// 		}

// 	case "list":
// 		if userName != "" || phone != "" {
// 			fmt.Println("invalid arguments")
// 			return
// 		}
// 		for _, contact := range contacts {
// 			fmt.Println(contact.Name, contact.Phone)
// 		}
// 	case "search":
// 		if userName == "" || phone != "" {
// 			fmt.Println("invalid arguments")
// 			return
// 		}
// 		for _, contact := range contacts {
// 			if userName == contact.Name {
// 				fmt.Println(contact.Name, contact.Phone)
// 			}
// 		}
// 	case "delete":
// 		if userName == "" || phone != "" {
// 			fmt.Println("invalid arguments")
// 			return
// 		}
// 		var filteredContacts []Contact
// 		for _, contact := range contacts {
// 			if userName != contact.Name {
// 				filteredContacts = append(filteredContacts, contact)
// 			}
// 		}

// 		jsonData := sliceTobytes(filteredContacts)

// 		err := os.WriteFile(fileName, jsonData, 0666)
// 		if err != nil {
// 			fmt.Println("Err: ", err)
// 		}
// 	}

// }

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
)

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func loadContacts(filename string) []Contact {
	data, err := os.ReadFile(filename)
	if err != nil || len(data) == 0 {
		return []Contact{}
	}
	var contacts []Contact
	json.Unmarshal(data, &contacts)
	return contacts
}

func saveContacts(filename string, contacts []Contact) {
	data, err := json.MarshalIndent(contacts, "", "  ") // Indented = readable JSON
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	os.WriteFile(filename, data, 0666)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: add <name> <phone> | list | search <name> | delete <name>")
		return
	}

	filename := "contacts.json"
	action := os.Args[1]
	contacts := loadContacts(filename)

	switch action {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Usage: add <name> <phone>")
			return
		}
		name := os.Args[2]
		phone := os.Args[3]
		contacts = append(contacts, Contact{Name: name, Phone: phone})
		saveContacts(filename, contacts)
		fmt.Println("Added:", name)

	case "list":
		if len(contacts) == 0 {
			fmt.Println("No contacts")
			return
		}
		for _, c := range contacts {
			fmt.Printf("%s: %s\n", c.Name, c.Phone)
		}

	case "search":
		if len(os.Args) < 3 {
			fmt.Println("Usage: search <name>")
			return
		}
		name := os.Args[2]
		found := false
		for _, c := range contacts {
			if c.Name == name {
				fmt.Printf("Found: %s - %s\n", c.Name, c.Phone)
				found = true
			}
		}
		if !found {
			fmt.Println("Not found:", name)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <name>")
			return
		}
		name := os.Args[2]
		originalLen := len(contacts)
		// Delete safe way: filter instead of deleting while looping
		contacts = slices.DeleteFunc(contacts, func(c Contact) bool {
			return c.Name == name
		})
		if len(contacts) == originalLen {
			fmt.Println("Not found:", name)
		} else {
			saveContacts(filename, contacts)
			fmt.Println("Deleted:", name)
		}

	default:
		fmt.Println("Unknown command:", action)
	}
}
