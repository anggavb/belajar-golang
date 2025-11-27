package main

import (
	"belajar-golang/database"
	"belajar-golang/helper"
	_ "belajar-golang/internal" // blank identifier (_), ini berfungsi supaya package yang mau dipake cuma initnya aja tanpa ada method yang dipakai
	"fmt"
)

func main() {
	result := helper.SayHello("Angga")
	fmt.Println(result)

	fmt.Println(database.GetConnection())
}
