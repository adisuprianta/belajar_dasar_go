package main;

import "fmt";

func getFullName(firstName string, lastName string) (string, string){
	return firstName, lastName;
}



func main() {
	// firstName,lastName:= getFullName("Leon","Khannedy");
	// fmt.Println(firstName+" "+lastName);

	//ini kalo ingin menghiraukan salat satu retrunnya
	firstName,_:= getFullName("Leon","Khannedy");
	fmt.Println(firstName);



	

	
}