package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// jimPointer := &jim // turn value into address with &value
	jim.updateName("Jimmy")
	jim.print()
}

// turn address into value with *address
func (pointerToPerson *person) updateName(newFirstName string) { // The star where a type should be is a type description - it means we're working with a pointer to a person
	(*pointerToPerson).firstName = newFirstName // the star here is an operator - it means we want to manipulate the value the pointer is referencing.
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
