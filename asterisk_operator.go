package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Surabaya", "Jawa Tengah", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Malang"

	fmt.Println(address1) // berubah
	fmt.Println(address2) //berubah menjadi malang

	// dengan dan artinya address2 berubah pointernya jadi Address yang lain
	// membuat address 1 dan 2 berbeda value
	//address2 = &Address{"Sidoarjo", "Jawa Tengah", "Indonesia"}

	//fmt.Println(address1) // tetap nilainya malang
	//fmt.Println(address2) //berubah menjadi sidoarjo

	//ini akan membuat pointernya berubah ke Address yang surabaya baik address 1 atau 2
	*address2 = Address{"Pasuruan", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1) // berubah menjadi pasuruan
	fmt.Println(address2) //berubah menjadi pasuruan
}
