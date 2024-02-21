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
	// alex := person{"Shubham", "Satyawali"}
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
