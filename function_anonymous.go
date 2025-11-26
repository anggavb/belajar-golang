package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool { // mirip di JS ==> let blacklist = fn()
		return name == "anjrit"
	}

	registerUser("anjrit", blacklist)
	registerUser("Angga", func(name string) bool {
		return name == "anjrit"
	})
}
