package main

import "fmt"

func endApp() {
	fmt.Println("End Application")
	message := recover()
	fmt.Println(message)
}

func runApplication(error bool) {
	defer endApp()

	if error {
		panic("Error!") // buat berhentiin app dan keluarin valuenya, di JS PHP mirip try/catch bagian catchnya tapi versi jauh lebih simpel
		// setelah panic tidak ada code yang di eksekusi kebawahnya
	}

	fmt.Println("Welcome to the main application!")
}

func main() {
	runApplication(true)

	// panic = kayak parameter error yang ada di catch PHP/JS
	// recover = ambil value yang ada di dalam panic ketika terjadi panic
}
