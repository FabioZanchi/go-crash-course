package main

import "fmt"

func main() {

	//Deifine Map
	emails := make(map[string]string)

	//Assign Key Values
	emails["Bob"] = "bob@gmail.com"
	emails["Eddy"] = "eddy@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	//Declar map and add Key Values
	kvemails := map[string]string{"NewBob": "newbob@gmail.com", "NewEddy": "eddy@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	fmt.Println(kvemails)

	//Delete from Map
	delete(emails, "Bob")
	fmt.Println(emails)

}
