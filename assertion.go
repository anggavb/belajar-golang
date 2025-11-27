package main

import "fmt"

func random() any {
	return 12312
}

func main() {
	result := random()
	//randomString := result.(string)
	//fmt.Println(randomString)

	//resultInt := result.(int) // ini bakal error atau panic
	//fmt.Println(resultInt)

	// contoh penanganan panic jika tidak tau type nya apa, nebak berdasarkan type
	switch val := result.(type) {
	case string:
		fmt.Println("String", val)
	case int:
		fmt.Println("Int", val)
	default:
		fmt.Println("Unknow Type")
	}
}
