package main

import "fmt"

func main() {
	//ini untuk initial variable 
	var name string

	name  = "ketut"
	fmt.Println(name)
	name = "ketut adi"
	fmt.Println(name)

	//tipe data bisa optional kalo langsung kasi nilai
	var number = 123
	fmt.Println(number)



	//kata kunci bisa optional kalo langsung kasi nilai tapi harus dengan titik dua :, 
	//ini hanya di pakek saat declarasi awal variable

	//contoh
	// name := "contoh" ini error karena membuat variable name ulang

	contohNae := "ketut"
	fmt.Println(contohNae)


	//ini cara mendeclarasikan variable secara banyak
	// dan variable itu harus digunakan jika tidak maka error
	var(
		firstName = "first"
		lastName  = "last"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)




} 