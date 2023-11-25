package main;

import "fmt";
func hello(){
	fmt.Println("Hello world ");
}

func sayHello(name string, age int){
	fmt.Println("Hello ", name);
	fmt.Println("Age: ", age);
}

func main() {
	sayHello("Adi", 23);
	hello();
}