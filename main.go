package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contact struct {
	Name    string
	Phone   string
	Address string
}

func main() {
	CreateContact()
}

func CreateContact() *Contact {
	reader := bufio.NewReader(os.Stdin)

	// Ask for contact name
	fmt.Print("Enter contact name: ")
	name, _ := reader.ReadString('\n')

	// Ask for contact phone
	fmt.Print("Enter contact phone: ")
	phone, _ := reader.ReadString('\n')

	// Ask for contact address
	fmt.Print("Enter contact address: ")
	address, _ := reader.ReadString('\n')

	fmt.Println("==== Contact summary ====")
	fmt.Println("Name:", name)
	fmt.Println("Phone:", phone)
	fmt.Println("Address:", address)

	return NewContact(name, phone, address)
}

func NewContact(name, phone, address string) *Contact {
	return &Contact{
		Name:    name,
		Phone:   phone,
		Address: address,
	}
}
