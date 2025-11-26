package main

import "fmt"

// lebih menguntungkan dibanding pake slice karena  ketika function digunakan perlu dikirim dalam bentuk slice juga
func sumData(numbers ...int) int { // pake ... itu jadinya seperti tanpa batas, asal type datanya pasti sama
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
func main() {
	total := sumData(1, 2, 3, 4, 5)
	fmt.Println(total)

	fmt.Println(sumData(10, 10, 10, 10, 10))

	sliceNumbers := []int{12, 23, 34, 45, 56}
	fmt.Println(sumData(sliceNumbers...)) // dibuat begini supaya diturunin jadi variabel argument
	// mirip kayak di JS kalo ada data json misalkan mau diturunkan jadi valuenya aja yang di kirim pake ... juga
	// atau kalo di PHP itu kayak array_values
}
