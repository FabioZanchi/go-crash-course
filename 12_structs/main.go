package main

import (
	"fmt"
	"strconv"
)

//Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	//Alternative
	firstName, lastName, city, gender string
	age                               int
}

//Methods
//Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + "." //similar to this.blah
}

//hasBirthday method (pointer receiver) --> pointer receiver neeeds the *
func (p *Person) hasBirthday() {
	p.age++
}

// getmarried (pointer receiver)
func (p *Person) getmarried(spouseLastName string) {
	if p.gender == "M" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "F", age: 25}

	//Alternative
	person2 := Person{"Samantha", "Smith", "Boston", "F", 25}

	fmt.Println(person1)
	fmt.Println(person2)

	//Get single field
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

	person1.hasBirthday()

	person1.getmarried("Williams")

	fmt.Println(person1.greet())

}
