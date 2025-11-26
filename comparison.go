package main

import "fmt"

func main() {
	var name1 = "Angga"
	var name2 = "Angga"

	var result = name1 == name2

	fmt.Println(result)
	fmt.Println(name1 != name2)

	var a = "a"
	var b = "b"

	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}
