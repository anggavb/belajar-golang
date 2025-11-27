package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Maried() { // yang penting tanda (*) untuk menandakan pointer supaya struct bisa terupdate
	man.Name = "Mr. " + man.Name
}

func main() {
	man := Man{Name: "Angga"}
	fmt.Println(man)

	man.Maried()
	fmt.Println(man)

}
