package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}
func main() {
	coba := getGoodBye

	fmt.Println(coba("Angga")) // jadi bikin variabel baru tapi dijadiin function
}
