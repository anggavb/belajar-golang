package main

import "fmt"

func main() {
	name := "asdfas"

	switch name {
	case "Angga":
		fmt.Println("Halo Angga")
	case "Joko":
		fmt.Println("Halo Joko")
	default:
		fmt.Println("Ini siapa???")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama udah bener")
	}

	// no condition
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan bener")
	default:
		fmt.Println("Nama udah bener")
	}
}
