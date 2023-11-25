package main

import "fmt"	

func main() {
	var names [3]string
	names[0] = "John"
	names[1] = "Mary"
	names[2] = "david"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])


	// buat array langsung

	var values = [3] int{1, 2, 3}
	fmt.Println(values)
}