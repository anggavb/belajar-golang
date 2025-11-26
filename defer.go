package main

import "fmt"

func logging() {
	fmt.Println("Logging...")
}

func runApp() {
	defer logging() // fungsi defer adalah untuk eksekusi code di akhir baris code, tetapi tetap akan di eksekusi walaupun ada error
	fmt.Println("Running app...")
}

func main() {
	runApp()
}
