package main;

import "fmt";

// ini cara named retrun values
// dengan cara ini kita bisa tidak mengisi value dari variable 
// saat tidak mengisi value maka retrunnya akan berisi default value
// dari tipe datanya itu sendiri
func getComplete()(firstName, lastName string,age int){
	firstName = "Kevin";
    lastName = "Candra";
    age = 23;
    return firstName,lastName,age;
}

func main() {
	firstName,lastName,age := getComplete();
	fmt.Printf("firstName: %s lastName: %s age: %d \n", firstName, lastName, age);
	// fmt.Println("tes");
}