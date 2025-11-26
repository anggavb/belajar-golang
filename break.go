package main

import "fmt"

func main() {
	// break
	for i := 0; i < 10; i++ {
		if i == 5 { // ini akan setop sampe i ke 5
			break
		}
		fmt.Println("Perulangan ke", i)
	}
}
