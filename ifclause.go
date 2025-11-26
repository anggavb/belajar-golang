package main

import "fmt"

func main() {
	name := "Angga"

	if name == "Angga" {
		fmt.Println("Halo " + name) // concat mirip kayak di JS pake (+)
	} else if name == "Jarwo" {
		fmt.Println("Jiahahaha " + name)
	} else {
		fmt.Println("Siapa yaa???")
	}

	// short statement

	// length := len(name)
	//if length > 5 {}

	if length := len(name); length > 5 { // disini bikin variabel baru sekalian kasih if clause
		fmt.Println("Nama kepanjangan")
	} else {
		fmt.Println("Benerr")
	}
}
