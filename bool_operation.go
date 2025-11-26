package main

import "fmt"

func main() {
	var nilai = 80
	var absensi = 90

	var lulusNilai bool = nilai > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilai && lulusAbsensi // semua kondisi harus true
	fmt.Println(lulus)

	fmt.Println(lulusAbsensi || lulusNilai) // hanya butuh salah satu true saja
}
