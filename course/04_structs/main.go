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
	// shubh := person{
	// 	firstName: "Shubham",
	// 	lastName:  "Satyawali",
	// 	contact: contactInfo{
	// 		email:   "shubh@shubh.com",
	// 		zipCode: 111111,
	// 	},
	// }

	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "alex@alex.com"
	alex.contact.zipCode = 222222

	// shubh.print()
	// alexPointer := &alex
	alex.updateName("Alexa")
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println(p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
