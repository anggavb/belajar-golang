package main

import "fmt"

// interface kosong / any
func Coba() any {
	//return 1
	//return true
	return "String nih"
}

func main() {
	kosong := Coba()
	fmt.Println(kosong)
}
