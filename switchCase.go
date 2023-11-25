package main

import "fmt"

func main() {
	name := "Jhonner"

	switch name {
	case "Jhon":
		fmt.Println("Hello", name)
	case "Ketut":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	switch length := len(name); length>5{
		case true:
            fmt.Println("Panjang name ", length)
        case false:
            fmt.Println("Panjang name kurang ")
	}
	length := len(name);

	switch {
	case length>7:
		fmt.Println("Panjang name ", length)
	case length<=7:
		fmt.Println("Panjang name kurang dari atau sama dengan 7")
	default:
		fmt.Println("Hi")
	}

}