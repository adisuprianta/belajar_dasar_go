package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}
// ini cara membuat func menempel pada stcuct 
func (customer Customer) sayHello(name string){
	fmt.Printf("Hello %s, my name is %s\n", name, customer.Name)
}

func main() {

	// ini cara pertama membuat struct
	var Ragil Customer

	Ragil.Name = "Ragil"
	Ragil.Address = "Bandung"
	Ragil.Age = 25
	fmt.Println(Ragil)
	fmt.Println(Ragil.Name)

	// bisa kayak gini untuk struct
	eko := Customer{
		Name:    "eko ",
		Address: "ragunan",
		Age:     20}

	fmt.Println(eko)
	fmt.Println(eko.Name)

	// bisa juga seperti ini tapi harus urut
	budi := Customer{"Budi", "Surabaya", 10}
	fmt.Println(budi)
	fmt.Println(budi.Name)


	// cara memanggil method atau function yang menempel pada struct	
	eko.sayHello(budi.Name)
}
