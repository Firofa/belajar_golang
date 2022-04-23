package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool { //function bisa dimasukan kedalam variable tanpa perlu memberikan nama function sehingga bisa disebut anonymous function
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("programmer", blacklist)

	registerUser("root", func(name string) bool { //Function juga bisa dimasukan langsung ke dalam parameter tanpa perlu memberikan nama function sehingga bisa disebut anonymous function
		return name == "root"
	})
	registerUser("developer", func(name string) bool {
		return name == "root"
	})
}

// Function yang tidak memiliki nama Function)
