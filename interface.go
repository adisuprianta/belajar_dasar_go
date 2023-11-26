package main

import "fmt"

type HasName interface {
	GetName() string
}
type Person struct {
	Name string
}
type Animal struct {
	Name string
}

// ini artinya mengimplementasikan interface dari HasName GetName karena
// tipe retrunnya sama;
func (person Person) GetName() string {
	return person.Name
}

//jadi dalam golang tidak harus secara implisif mendeclarasikan bahwa function ini adalah
// implementasi dari sebuah interface cukup kalo retrunnya sama maka artinya dia sudah mengimplementasikan
// sebuah interface
func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(value HasName) {
	fmt.Printf("Hello %s \n", value.GetName())
}

func main() {
	person := Person{"Eko"}

	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
