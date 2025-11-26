package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var hasil = a + b

	fmt.Println(hasil)
	fmt.Println(20 % 3)

	var i = 100
	i += 20
	fmt.Println(i)
	i %= 3
	fmt.Println(i)

	var j = 1
	j++ // j = j+j (1+1)
	fmt.Println(j)
	j--
	fmt.Println(j)
}
