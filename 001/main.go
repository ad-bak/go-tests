package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
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
		jim.updateName("Jimmy")
		jim.print()
	}

	func (pointerToPerson *person) updateName(newFirstName string) {
		(*pointerToPerson).firstName = newFirstName 
	}

	func (p person) print() {
		fmt.Printf("%+v", p)
	}