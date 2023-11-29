package main

import "fmt"

type Address struct {
	City, Province, Country string
}

// dengan kasi tanda bintang di tipe data parameter ini membuat
// parameter yang bisa dikirim berupa pointer
func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Surabaya", "Jawa Timur", ""}
	// ini mengirim parameter pointer
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}
