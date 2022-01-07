package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo // just writing contactInfo equivalent to contactInfo contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func main() {
	// alex := person{"Alex", "Anderson"}
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	var alex person

	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact.email = "apaarshrm39@gmail.com"
	fmt.Println(alex) // {Alex Anderson}
	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "apaarshrm39@gmail.com",
			zipCode: 12345,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("pappu")

	jim.updateName("jimmy") // this is a shortcut in Go {I	F RECIEVER IS DEFINED WITH POINTER}
	jim.print()
	x := 0xc000014110
	fmt.Println(&x)
}
