package main

import "fmt"

func main() {
	const firstName string = "Angga"
	const lastName = "Vb"

	//firstName = "Joko"
	fmt.Println(firstName, lastName)

	const (
		first = "Angga"
		last  = "Vb"
	)

	fmt.Println(first, last)
}
