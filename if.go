package main

import "fmt"

func main() {
	name := "Jhon"

	if name == "Jhon" || name == "budi"{
		fmt.Println("Hello", name);
	}else if name == "Ketut"{
		fmt.Println("Hello", name);
	}else{
		fmt.Println("Hi, Boleh Kenalan?");
	}

	if lenght :=len(name); lenght>5{
		fmt.Println("Panjang name ", lenght);
		// fmt.Printf("Panjang name %d", lenght);
	}else {
		fmt.Println("Panjang name kurang ");
        // fmt.Printf("Panjang name %d", lenght);
	}

}