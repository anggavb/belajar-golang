package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHelloo(hasName HasName) {
	fmt.Println(hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "Angga"}
	sayHelloo(person)

	animal := Animal{"Kucing"}
	sayHelloo(animal)

	/*
		jadi maksudnya adalah kita mau panggil function sayHeloo
		tapi karena parameter dari sayHeloo ini adalah interface,
		maka data yang dimasukkan ke parameter harus sesuai dengan isi dari interface tsb
		disini contohnya interface ini isinya adalah sebuah function yang return nya string (GetName() string)
		data turunan dari interface harus berupa struct
		jadi ketika kita mau bikin function atau suatu keluaran dari GetName(),
		maka keluaran atau return data dari GetName() harus berupa struct
		contoh disini adalah bikin function anonymous yang parameternya dari struct
	*/
}
