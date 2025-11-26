package main

import "fmt"

func main() {
	//counter := 1
	//
	//for counter <= 10 {
	//	fmt.Println(counter)
	//	counter++
	//}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println(counter)
	}

	fmt.Println("Perulangan Selesai")

	names := []string{"Angga", "Dzaky", "Beny"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	for index, name := range names {
		fmt.Println("index", index, "name", name) // ternyata (,) juga bisa jadi concat, tar coba dicari tau dulu yang bener yang mana
	}

	for _, name := range names { // kalo ada yang gak kepake musti diganti jadi (_) supaya gak error
		fmt.Println(name)
	}
}
