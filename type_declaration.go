package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAngga NoKTP = "1391273981273"

	var ktp string = "1234567890"
	var changeKTP NoKTP = NoKTP(ktp)

	fmt.Println(ktpAngga)
	fmt.Println(changeKTP)
}
