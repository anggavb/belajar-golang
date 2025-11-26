package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 { // kalo modulo 2 atau i habis dibagi 2 maka di skip
			continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
