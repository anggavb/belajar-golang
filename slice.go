package main

import "fmt"

func main() {
	names := [...]string{
		"John",
		"Paul",
		"George",
		"Ringo",
		"Jackson",
		"Jigsaw",
		"Koch",
	}

	slice := names[4:6]
	fmt.Println(slice)

	slice2 := names[2:]
	fmt.Println(slice2)

	slice3 := names[:2]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1) // Sabtu Baru, Minggu Baru
	fmt.Println(days)      // Senin, Selasa, Rabu, Kamis, Jumat, Sabtu, Minggu

	daySlice2 := append(daySlice1, "Libur Baru") // kalo data dari arraynya udah melewati batasnya, maka data slice baru akan dijadikan array yang baru
	fmt.Println(daySlice1)
	fmt.Println(daySlice2) // Sabtu Baru, Minggu Baru, Libur Baru
	fmt.Println(days)      // data disini tidak menampilkan update data dari daySlice2 walaupun ada yang diubah

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Angga"
	newSlice[1] = "Angga"
	//newSlice[2] = "Angga" // kalo lebih dari batas (pointer), gabisa pake cara ini. harus di append()

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Angga")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2[0] = "Joko"
	fmt.Println(newSlice2)
	fmt.Println(newSlice) // ketika ada perubahan dan kapatasnya masih cukup, maka data yang di update juga akan terupdate ke slice lama

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), len(fromSlice))

	copy(toSlice, fromSlice) // untuk mengcopy data dari array ke slice baru
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	// isi antara array dan slice sebenarnya sama saja, hanya penggunaanya saja yang berbeda
	// kalo array datanya static jadi dari awal perlu di tentuin dulu kapasitasnya berapa
	// kalo slice itu bebas aja, mirip kayak di PHP atau JS dan datanya bisa ditambah lagi sesuai dengan kapasitas yang ditentuka
}
