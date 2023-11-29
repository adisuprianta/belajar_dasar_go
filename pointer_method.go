package main

import "fmt"

type Man struct {
	Name string
}

// kalo tanpa operator * maka Name di dalam struct Man tidak akan berubah
// karena func ini akan pass by value bukan pass by reference
// nah dengan menambahkan tanda * maka passnya adalah pass by reference
func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func main() {
	ekoo := Man{"Eko"}
	ekoo.Married()
	fmt.Println(ekoo)
}
