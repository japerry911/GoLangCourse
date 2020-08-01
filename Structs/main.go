package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jack := person{"Jack", "Sky", contactInfo{"jack@skymail.com", 12345}}
	jack.lastName = "Skylord"
	jack.contact.zipCode = 54321
	jack.print()
	jack.updateName("Skylord")
	jack.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
