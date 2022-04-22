package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Kabayan"
	middleName = "bin"
	lastName = "Abah"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()

	fmt.Println(firstName, middleName, lastName)
}
