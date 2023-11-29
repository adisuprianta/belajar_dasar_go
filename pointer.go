package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Tengah", "Indonesia"}
	address2 := address1 // copy by value

	address2.City = "Sidoarjo"

	fmt.Println(address1) //tidak berubah
	fmt.Println(address2) //berubah menjadi sidoarjo

	// menggunakan operator & untuk copy by reference atau memorynya bukan by value
	//address3 := &address1
	//address3.City = "Malang"
	//
	//fmt.Println(address1) // berubah
	//fmt.Println(address3) //berubah menjadi malang

	// cara lain kalo pakek var
	var address Address = Address{"Surabaya", "Jawa Tengah", "Indonesia"}
	var address3 *Address = &address
	address3.City = "Malang"

	fmt.Println(address)  // berubah
	fmt.Println(address3) //berubah menjadi malang

}
