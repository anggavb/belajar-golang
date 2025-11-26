package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Angga"
	names[1] = "Visual"
	names[2] = "Basic"

	// names = [Angga Visual Basic]
	// kalo di PHP = ["Angga", "Visual", "Basic"]

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// error karena variabel names hanya menampung 3 data saja
	//names[3] = "Joko"

	var values = [3]int{1, 2, 3} // kalo mau diturunin harus dikasih koma (,) di akhir value

	fmt.Println(values)
	fmt.Println(values[2])
	fmt.Println(len(values))

	values[2] = 90
	fmt.Println(values)

	var allValues = [...]int{ // [...] wajib di deklarasikan langsung isi datanya tidak bisa setelahnya (line 6 gabisa pake ini)
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(allValues)
	fmt.Println(len(allValues))
	fmt.Println(allValues[2])

	// pada Go tidak bisa menghapus atau mengurangi isi dari array
	// kalo udah di tentukan ada 3 maka harus 3 isinya
	// jadi kalo memang mau hapus tinggal di kosongkan "" jika string atau diubah jadi 0 jika int
}
