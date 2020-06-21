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

	var thirdperson person
	thirdpersonPointer := &thirdperson
	thirdperson.firstName = "ezisalex"
	thirdperson.lastName = "ands"
	thirdperson.contactInfo.email = "a@gmail.com"
	thirdperson.contactInfo.zipCode = 1111
	thirdperson.print()
	thirdpersonPointer.updateName("stupidname")
	thirdperson.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
