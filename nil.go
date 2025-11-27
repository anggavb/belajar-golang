package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil // nil hanya bisa dipakai di: interface, function (as value or parameter), map, slice, pointer, channel
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := NewMap("")
	if data == nil {
		fmt.Println("Data map masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}
