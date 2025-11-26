package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func spamFilter(name string) string {
	if name == "Anjritt" {
		return "..."
	} else {
		return name
	}
}

// type Declarators
type Filter func(string) string

func sayHelloWithFilter2(name string, filter Filter) { // jadinya diganti ke sebuah type variabel untuk nampung typenya supaya gak kepanjangan codenya
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func main() {
	sayHelloWithFilter("Angga", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjritt", filter)

	sayHelloWithFilter2("Anggaaaa", filter)
}
